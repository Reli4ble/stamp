package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"unicode/utf8"

	"github.com/Reli4ble/stamp/parser"
	tpl "github.com/Reli4ble/stamp/template"
)

// RunAutoScan recursively scans the directory specified in opts.InDir,
// renders every file as a template using the merged configuration,
// and overwrites the file in place if the rendered content differs from the original.
// Files that are not valid UTF-8 or that produce render errors are skipped.
// Only files that are actually modified are printed (with their absolute paths).
// The --force-success option controls the exit status.
func RunAutoScan(opts Options) {
	// Merge data from .env and YAML (if provided)
	envVars, _ := parser.LoadEnv(opts.EnvPath)
	yamlVars, _ := parser.LoadYAML(opts.YamlPath)
	data := parser.MergeMaps(envVars, yamlVars)

	encounteredError := false

	err := filepath.WalkDir(opts.InDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error accessing", path, ":", err)
			encounteredError = true
			return nil
		}
		if d.IsDir() {
			return nil
		}

		originalBytes, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading file", path, ":", err)
			encounteredError = true
			return nil
		}

		// Skip file if content is not valid UTF-8
		if !utf8.Valid(originalBytes) {
			fmt.Printf("Skipping non-UTF8 file: %s\n", path)
			return nil
		}
		original := string(originalBytes)

		// Render template using the merged data
		rendered, err := tpl.RenderTemplate(original, data, opts.Strict)
		if err != nil {
			fmt.Printf("Render error in %s: %v\n", path, err)
			// Bei Render-Fehlern wird die Datei übersprungen,
			// aber der Fehler wird nicht als kritischer Fehler gewertet.
			return nil
		}

		// Überschreibe die Datei nur, wenn sich der Inhalt ändert
		if rendered != original {
			absPath, _ := filepath.Abs(path)
			if opts.DryRun {
				fmt.Printf("Dry-run: %s would be overwritten.\n", absPath)
			} else {
				err = os.WriteFile(path, []byte(rendered), 0644)
				if err != nil {
					fmt.Printf("Error writing file %s: %v\n", path, err)
					encounteredError = true
				} else {
					fmt.Printf("✔ Processed: %s\n", absPath)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the directory:", err)
		encounteredError = true
	}

	if encounteredError && !opts.ForceSuccess {
		os.Exit(1)
	}
}

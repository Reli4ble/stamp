package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Reli4ble/stamp/parser"
	tpl "github.com/Reli4ble/stamp/template"
)

// RunAutoScan recursively scans the directory specified in opts.InDir,
// renders every file as a template, and overwrites the file in place
// if the rendered content differs from the original. The absolute path
// of the processed file is printed. The output directory is not used in auto-scan mode.
func RunAutoScan(opts Options) {
	// Merge data from .env and YAML
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

		// Read file content
		originalBytes, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Error reading file", path, ":", err)
			encounteredError = true
			return nil
		}
		original := string(originalBytes)

		// Render template
		rendered, err := tpl.RenderTemplate(original, data, opts.Strict)
		if err != nil {
			fmt.Println("Render error in", path, ":", err)
			encounteredError = true
			return nil
		}

		// If rendered content differs, overwrite the file
		if rendered != original {
			absPath, _ := filepath.Abs(path)
			if opts.DryRun {
				fmt.Printf("Dry-run: %s would be overwritten.\n", absPath)
			} else {
				err = os.WriteFile(path, []byte(rendered), 0644)
				if err != nil {
					fmt.Println("Error writing file", path, ":", err)
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

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/markus/stamp/parser"
	tpl "github.com/markus/stamp/template"
)

func RunRender(opts Options) {
	envVars, _ := parser.LoadEnv(opts.EnvPath)
	yamlVars, _ := parser.LoadYAML(opts.YamlPath)
	data := parser.MergeMaps(envVars, yamlVars)

	// Einzelne Template-Verarbeitung
	if opts.InFile != "" {
		tmplContent, err := os.ReadFile(opts.InFile)
		if err != nil {
			fmt.Println("Fehler beim Lesen der Template-Datei:", err)
			os.Exit(1)
		}
		rendered, err := tpl.RenderTemplate(string(tmplContent), data, opts.Strict)
		if err != nil {
			fmt.Println("Render-Fehler:", err)
			os.Exit(1)
		}
		if opts.DryRun {
			absIn, _ := filepath.Abs(opts.InFile)
			absOut, _ := filepath.Abs(opts.OutFile)
			fmt.Printf("Dry-run: %s would be written to %s\n", absIn, absOut)
			fmt.Println(rendered)
		} else {
			err = os.WriteFile(opts.OutFile, []byte(rendered), 0644)
			if err != nil {
				fmt.Println("Fehler beim Schreiben in", opts.OutFile, ":", err)
				os.Exit(1)
			}
			absIn, _ := filepath.Abs(opts.InFile)
			absOut, _ := filepath.Abs(opts.OutFile)
			fmt.Printf("✔ Template processed: %s -> %s\n", absIn, absOut)
		}
		return
	}

	// Batch-Verarbeitung
	if opts.InDir != "" {
		files, _ := os.ReadDir(opts.InDir)
		os.MkdirAll(opts.OutDir, 0755)
		for _, f := range files {
			if !strings.HasSuffix(f.Name(), ".st") {
				continue
			}
			inPath := filepath.Join(opts.InDir, f.Name())
			outName := strings.TrimSuffix(f.Name(), ".st")
			outPath := filepath.Join(opts.OutDir, outName)
			content, err := os.ReadFile(inPath)
			if err != nil {
				fmt.Println("Fehler beim Lesen:", inPath)
				continue
			}
			rendered, err := tpl.RenderTemplate(string(content), data, opts.Strict)
			if err != nil {
				fmt.Println("Fehler in", f.Name()+":", err)
				continue
			}
			if opts.DryRun {
				absIn, _ := filepath.Abs(inPath)
				absOut, _ := filepath.Abs(outPath)
				fmt.Printf("Dry-run: %s would be written to %s\n", absIn, absOut)
				fmt.Printf("=== %s ===\n%s\n", f.Name(), rendered)
			} else {
				err = os.WriteFile(outPath, []byte(rendered), 0644)
				if err != nil {
					fmt.Println("Fehler beim Schreiben in", outPath, ":", err)
					continue
				}
				absIn, _ := filepath.Abs(inPath)
				absOut, _ := filepath.Abs(outPath)
				fmt.Printf("✔ Processed: %s -> %s\n", absIn, absOut)
			}
		}
	}
}

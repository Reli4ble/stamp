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
			fmt.Println(rendered)
		} else {
			os.WriteFile(opts.OutFile, []byte(rendered), 0644)
		}
		fmt.Println("✔ Template erfolgreich verarbeitet.")
		return
	}

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
				fmt.Printf("=== %s ===\n%s\n", f.Name(), rendered)
			} else {
				os.WriteFile(outPath, []byte(rendered), 0644)
			}
			fmt.Println("✔", f.Name(), "verarbeitet.")
		}
	}
}

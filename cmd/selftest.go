package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Reli4ble/stamp/parser"
	tpl "github.com/Reli4ble/stamp/template"
)

func RunSelfTest(opts Options) {
	fmt.Println("== Stamp Self-Test ==")

	// Check .env file
	if opts.EnvPath == "" {
		fmt.Println("[!] No .env file specified")
	} else {
		_, err := parser.LoadEnv(opts.EnvPath)
		printResult("ENV loaded", err)
	}

	// Check YAML file
	if opts.YamlPath == "" {
		fmt.Println("[!] No YAML file specified")
	} else {
		_, err := parser.LoadYAML(opts.YamlPath)
		printResult("YAML loaded", err)
	}

	// Check template directory
	if opts.InDir == "" {
		fmt.Println("[!] No template directory specified")
	} else {
		templates, ok := listTemplates(opts.InDir)
		printCheck("Templates found", ok)
		if !ok {
			fmt.Println("[!] No templates to render.")
		}

		var data map[string]interface{}
		if opts.EnvPath != "" || opts.YamlPath != "" {
			env, _ := parser.LoadEnv(opts.EnvPath)
			yml, _ := parser.LoadYAML(opts.YamlPath)
			data = parser.MergeMaps(env, yml)
		} else {
			data = make(map[string]interface{})
		}

		allOk := true
		for _, file := range templates {
			content, _ := os.ReadFile(file)
			_, err := tpl.RenderTemplate(string(content), data, false)
			if err != nil {
				fmt.Println("[✗]", file, ":", err)
				allOk = false
			}
		}
		if len(templates) > 0 && allOk {
			fmt.Println("[✓] All templates rendered successfully.")
		}
	}

	// Check output directory
	if opts.OutDir == "" {
		fmt.Println("[!] No output directory specified")
	} else {
		testFile := filepath.Join(opts.OutDir, ".stamp_test")
		err := os.WriteFile(testFile, []byte("test"), 0644)
		printResult("Output directory writable", err)
		_ = os.Remove(testFile)
	}

	fmt.Println("== Finished ==")
}

func listTemplates(dir string) ([]string, bool) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, false
	}
	var result []string
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".st") {
			result = append(result, filepath.Join(dir, e.Name()))
		}
	}
	return result, len(result) > 0
}

func printResult(name string, err error) {
	if err != nil {
		fmt.Println("[✗]", name, ":", err)
	} else {
		fmt.Println("[✓]", name)
	}
}

func printCheck(name string, ok bool) {
	if ok {
		fmt.Println("[✓]", name)
	} else {
		fmt.Println("[✗]", name)
	}
}

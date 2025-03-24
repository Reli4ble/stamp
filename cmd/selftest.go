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

	// Prüfe .env-Datei
	if opts.EnvPath == "" {
		fmt.Println("[!] Keine .env-Datei angegeben")
	} else {
		_, err := parser.LoadEnv(opts.EnvPath)
		printResult("ENV geladen", err)
	}

	// Prüfe YAML-Datei
	if opts.YamlPath == "" {
		fmt.Println("[!] Keine YAML-Datei angegeben")
	} else {
		_, err := parser.LoadYAML(opts.YamlPath)
		printResult("YAML geladen", err)
	}

	// Prüfe Template-Ordner
	if opts.InDir == "" {
		fmt.Println("[!] Kein Template-Ordner angegeben")
	} else {
		templates, ok := listTemplates(opts.InDir)
		printCheck("Templates gefunden", ok)
		if !ok {
			fmt.Println("[!] Es wurden keine Templates zum Rendern gefunden.")
		}

		// Daten zusammenführen, falls ENV oder YAML angegeben wurden
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
			fmt.Println("[✓] Alle Templates rendern erfolgreich.")
		}
	}

	// Prüfe Ausgabeordner
	if opts.OutDir == "" {
		fmt.Println("[!] Kein Ausgabeordner angegeben")
	} else {
		testFile := filepath.Join(opts.OutDir, ".stamp_test")
		err := os.WriteFile(testFile, []byte("test"), 0644)
		printResult("Ausgabeverzeichnis beschreibbar", err)
		_ = os.Remove(testFile)
	}

	fmt.Println("== Fertig ==")
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

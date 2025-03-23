package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/markus/stamp/parser"
	tpl "github.com/markus/stamp/template"
)

func RunSelfTest(opts Options) {
	fmt.Println("== Stamp Self-Test ==")

	_, err1 := parser.LoadEnv(opts.EnvPath)
	printResult("ENV geladen", err1)

	_, err2 := parser.LoadYAML(opts.YamlPath)
	printResult("YAML geladen", err2)

	templates, ok := listTemplates(opts.InDir)
	printCheck("Templates gefunden", ok)

	data := make(map[string]interface{})
	if err1 == nil {
		env, _ := parser.LoadEnv(opts.EnvPath)
		yml, _ := parser.LoadYAML(opts.YamlPath)
		data = parser.MergeMaps(env, yml)
	}

	allOk := true
	for _, file := range templates {
		content, _ := os.ReadFile(file)
		_, err := tpl.RenderTemplate(string(content), data, false)
		if err != nil {
			fmt.Println("[✗]", file)
			allOk = false
		}
	}
	if allOk {
		fmt.Println("[✓] Alle Templates rendern erfolgreich.")
	}

	testFile := filepath.Join(opts.OutDir, ".stamp_test")
	err := os.WriteFile(testFile, []byte("test"), 0644)
	printResult("Ausgabeverzeichnis beschreibbar", err)
	_ = os.Remove(testFile)

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
		fmt.Println("[✗]", name)
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

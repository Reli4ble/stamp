package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/markus/stamp/cmd"
)

func main() {
	selfTest := flag.Bool("self-test", false, "Führt einen Selbsttest durch")
	render := flag.Bool("render", false, "Rendert Templates")
	strict := flag.Bool("strict", false, "Bricht bei fehlenden Platzhaltern ab")
	dryRun := flag.Bool("dry-run", false, "Zeigt Ergebnis nur im Terminal")
	env := flag.String("env", "", "Pfad zur .env-Datei")
	yaml := flag.String("yaml", "", "Pfad zur .yaml-Datei")
	in := flag.String("in", "", "Template-Datei")
	out := flag.String("out", "", "Ausgabedatei")
	inDir := flag.String("in-dir", "", "Ordner mit .st-Dateien")
	outDir := flag.String("out-dir", "", "Zielordner")
	flag.Parse()

	opts := cmd.Options{
		EnvPath:  *env,
		YamlPath: *yaml,
		InFile:   *in,
		OutFile:  *out,
		InDir:    *inDir,
		OutDir:   *outDir,
		Strict:   *strict,
		DryRun:   *dryRun,
	}

	switch {
	case *selfTest:
		cmd.RunSelfTest(opts)
	case *render:
		cmd.RunRender(opts)
	default:
		fmt.Println("Bitte --render oder --self-test verwenden.")
		os.Exit(1)
	}
}

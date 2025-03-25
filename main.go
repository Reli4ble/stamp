package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Reli4ble/stamp/cmd"
)

func main() {
	selfTest := flag.Bool("self-test", false, "Run self-test")
	render := flag.Bool("render", false, "Render a single template")
	autoScan := flag.Bool("auto-scan", false, "Recursively render all files in the directory and overwrite in place")
	strict := flag.Bool("strict", false, "Enable strict mode (error on missing placeholders)")
	dryRun := flag.Bool("dry-run", false, "Display output in terminal without writing to file")
	forceSuccess := flag.Bool("force-success", false, "Force exit 0 even if errors occur")
	env := flag.String("env", "", "Path to .env file")
	yaml := flag.String("yaml", "", "Path to YAML file")
	in := flag.String("in", "", "Template file to render")
	out := flag.String("out", "", "Output file for rendered template")
	inDir := flag.String("in-dir", "", "Directory containing templates (for batch processing or auto-scan)")
	outDir := flag.String("out-dir", "", "Output directory for batch processing (ignored in auto-scan mode)")
	flag.Parse()

	// Prüfen: auto-scan und render dürfen nicht gemeinsam verwendet werden.
	if *autoScan && *render {
		fmt.Println("Error: You cannot use -render with -auto-scan. Please remove the -render flag when using auto-scan mode.")
		os.Exit(1)
	}

	opts := cmd.Options{
		EnvPath:      *env,
		YamlPath:     *yaml,
		InFile:       *in,
		OutFile:      *out,
		InDir:        *inDir,
		OutDir:       *outDir,
		Strict:       *strict,
		DryRun:       *dryRun,
		AutoScan:     *autoScan,
		ForceSuccess: *forceSuccess,
	}

	switch {
	case *selfTest:
		cmd.RunSelfTest(opts)
	case *render:
		cmd.RunRender(opts)
	case *autoScan:
		cmd.RunAutoScan(opts)
	default:
		fmt.Println("Please specify a mode: --render, --self-test, or --auto-scan")
		os.Exit(1)
	}
}

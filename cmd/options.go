package cmd

// Options holds all CLI options for stamp.
type Options struct {
	EnvPath      string
	YamlPath     string
	InFile       string
	OutFile      string
	InDir        string
	OutDir       string
	Strict       bool
	DryRun       bool
	AutoScan     bool
	ForceSuccess bool
}

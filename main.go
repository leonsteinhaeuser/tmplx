package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/leonsteinhaeuser/tmplx/template"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "tmpls [flags]",
		Short: "tmpls is a CLI tool for generating files from templates.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// read template file
			tmplTemplate, err := os.ReadFile(*flagValueTemplatePath)
			if err != nil {
				return err
			}

			dtSrc, err := template.DataSource(*flagValueFormat)
			if err != nil {
				return err
			}

			tmplData, err := dtSrc(*flagValueSource)
			if err != nil {
				return err
			}

			tmplBts, err := template.Render(tmplTemplate, tmplData)
			if err != nil {
				return err
			}

			if *flagValueDryRun {
				// print to stdout
				fmt.Println(string(tmplBts))
				return nil
			}

			// write output file
			err = os.WriteFile(*flagValueOutputPath, tmplTemplate, 0644)
			if err != nil {
				return err
			}

			return nil
		},
	}

	revision   string
	buildTime  time.Time
	dirtyBuild bool
	arch       string
	goos       string

	flagValueTemplatePath *string
	flagValueOutputPath   *string
	flagValueFormat       *string
	flagValueSource       *string
	flagValueDryRun       *bool
)

// we use the init function to setup the version information and flags for the root command
func init() {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	for _, set := range buildInfo.Settings {
		switch set.Key {
		case "vcs.revision":
			revision = set.Value
		case "vcs.time":
			buildTime, _ = time.Parse(time.RFC3339, set.Value)
		case "GOARCH":
			arch = set.Value
		case "GOOS":
			goos = set.Value
		case "vcs.modified":
			dirtyBuild = set.Value == "true"
		}
	}
	// add version info
	rootCmd.Version = fmt.Sprintf("%s - %v [ %s/%s ] [ %v ]", revision, buildTime, arch, goos, dirtyBuild)

	// setup flags
	flagValueTemplatePath = rootCmd.Flags().StringP("template", "t", "template.tmpl", "The path to the template file.")
	flagValueOutputPath = rootCmd.Flags().StringP("output", "o", "output.txt", "The path to the output file.")
	flagValueFormat = rootCmd.Flags().StringP("format", "f", "env", "The format of the template data. Valid values are: env, json, yaml.")
	flagValueSource = rootCmd.Flags().StringP("source", "s", "TMPLX_", "The path to the source file containing template data. Valid values are: <file>.<json|yaml>. If format is env, this flag caries the prefix for the environment variables.")
	flagValueDryRun = rootCmd.Flags().BoolP("dry-run", "d", false, "If set, the output will not be written to the output file.")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

package cmd

import (
  "fmt"
  "github.com/s0h311/go-monorepo/internal/config"
  "github.com/spf13/cobra"
)

var packageManagerConfigs = []string{
  "package.json",
  "composer.json",
}

var packages []string

func GetRunCmd() *cobra.Command {
  cmd := cobra.Command{
    Use:     "run",
    Aliases: []string{"r"},
    Short:   "Run a script in all packages or specific ones",
    Run: func(cmd *cobra.Command, args []string) {
      for _, packageName := range packages {
        runScript(packageName, args[0])
      }
    },
  }
  packages = config.GetConfig().Packages

  specifiedPackages := *cmd.Flags().StringArrayP("packages", "p", []string{}, "Packages to run the script")

  if len(specifiedPackages) > 0 {
    packages = specifiedPackages
  }

  return &cmd
}

func runScript(packageName string, script string) {
  fmt.Println(packageName)
  fmt.Println(script)
}

package main

import (
  "github.com/s0h311/go-monorepo/cmd"
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "mono",
  Short: "The simple and efficient monorepo tool",
  Run:   func(cmd *cobra.Command, args []string) {},
}

func main() {
  rootCmd.AddCommand(cmd.GetRunCmd())

  if err := rootCmd.Execute(); err != nil {
    panic(err)
  }
}

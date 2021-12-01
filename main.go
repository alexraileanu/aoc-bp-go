package main

import (
    "errors"
    "github.com/alexraileanu/aoc-init/cmd"
    "github.com/spf13/cobra"
)

func main() {
    rootCmd := &cobra.Command{
        Use: "aoc-init [path]",
        Args: func(cmd *cobra.Command, args []string) error {
            if len(args) < 1 {
                return errors.New("path is required")
            }

            return nil
        },
        Run: func(c *cobra.Command, args []string) {
            cmd.Create(args[0])
        },
    }

    rootCmd.Execute()
}

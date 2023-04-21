package cmd

import (
	"bigtools_cli/cmd/hadoop"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "bigtools [command]",
}

func init() {
	RootCmd.AddCommand(hadoop.Commands)
}

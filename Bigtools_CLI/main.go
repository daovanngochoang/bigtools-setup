package main

import "bigtools_cli/cmd"

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		return
	}

}

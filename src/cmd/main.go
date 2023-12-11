package cmd

import "github.com/spf13/cobra"

func Execute() error {
	root.AddCommand(aliases)
	root.AddCommand(cleanup)

	root.AddGroup(&cobra.Group{
		ID:    "config",
		Title: "Configurations:",
	})

	return root.Execute()
}

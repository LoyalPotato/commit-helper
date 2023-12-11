package cmd

import (
	"github.com/spf13/cobra"
	"github.com/LoyalPotato/commit-helper/src/git"
)

const cleanupShort = "Cleaner to revert aliases set"

var cleanup = &cobra.Command{
	Use: "cleanup",
	GroupID: "config",
	Short: cleanupShort,
	Run: func(cmd *cobra.Command, args []string) {
		git.UnsetAlias(alias)
	},
}
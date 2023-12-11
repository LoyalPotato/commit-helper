package cmd

import (
	"github.com/spf13/cobra"
	"github.com/LoyalPotato/commit-helper/src/git"
)

const aliasShort = "Command used to register alias in git"

const aliasLong = `
Command used to register alias in git

This is to make it easier to use by using it through git.

Alias registered is:
- ch

After registering the alias, use it like:
git ch
`

const alias = "ch"

var aliases = &cobra.Command{
	Use:     "aliases",
	GroupID: "config",
	Run:     runAliases,
	Short:   aliasShort,
	Long:    aliasLong,
}

func runAliases(cmd *cobra.Command, args []string) {
	git.AddAlias(alias, alias)
}

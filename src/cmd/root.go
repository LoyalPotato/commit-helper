package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/LoyalPotato/commit-helper/src/git"
	"github.com/LoyalPotato/commit-helper/src/messages"
)

const rootDesc = "CLI tool to help adhere to EposNow's commit style."

const rootInfo = `
CLI tool to help adhere to our commit style.

It asks you for the type of commit:
- BUGFIX
- FEATURE
- REFACTOR
- WIP
- OTHER

Then for the ticket associated with the commit.

And finally, asks for a description and commits whatever you've added.
`

var root = &cobra.Command{
	Use:   "ch",
	Short: rootDesc,
	DisableFlagsInUseLine: true,
	Long:  rootInfo,
	Run:   runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
	getCommitType()
	getTicket()
	getCommitMsg()
	git.Commit()
}

func getCommitType() {
	commitType, err := selectCommitTypePrompt()
	if err != nil {
		fmt.Println(fmt.Errorf(messages.Error, err))
		os.Exit(0)
	}

	if commitType == "OTHER" {
		other, err := otherPrompt()
		if err != nil {
			fmt.Println(fmt.Errorf(messages.Error, err))
			os.Exit(1)
		}
		commitType = other
	}

	git.BuildCommitMsg(fmt.Sprintf("[%s]", commitType))
}

func getTicket() {
	ticketPrefix := os.Getenv("TICKET_PREFIX")
	prompt := promptui.Prompt{
		Label: "Ticket",
		AllowEdit: true,
		HideEntered: true,
		Default: ticketPrefix,
	}
	ticket, err := prompt.Run()
	if err != nil {
		fmt.Println(fmt.Errorf(messages.Error, err))
		os.Exit(0)
	}

	if ticket != "" {
		git.BuildCommitMsg(fmt.Sprintf("[%s]", ticket))
	}
}

func getCommitMsg() {
	prompt := promptui.Prompt{
		Label: "Commit Message",
		HideEntered: true,
		Validate: func(s string) error {
			if s == "" {
				return errors.New("value cannot be empty")
			}

			return nil
		},
	}

	commitMsg, err := prompt.Run()
	if err != nil {
		fmt.Println(fmt.Errorf(messages.Error, err))
		os.Exit(0)
	}

	if commitMsg != "" {
		git.BuildCommitMsg(fmt.Sprintf(" %s", commitMsg))
	}
}

func selectCommitTypePrompt() (string, error){
	prompt := promptui.Select{
		Label:     "Commit Type",
		Items:     parseCommitTypes(messages.CommitTypes),
		IsVimMode: true,
		HideSelected: true,
		Templates: &promptui.SelectTemplates{
			Active: "{{. | bold}}",
			Help:   `Use arrow keys or vim movement {{ "(jk)" | faint }}`,
		},
	}
	_, commitType, err := prompt.Run()

	return commitType, err
}

func parseCommitTypes(commitTypes []string) []string {
	commits := []string{}
	commits = append(commits, commitTypes...)
	commits = append(commits, "OTHER")

	return commits
}

func otherPrompt() (string, error) {
	prompt := promptui.Prompt{
		Label: "Other",
		HideEntered: true,
		Validate: func(s string) error {
			if s == "" {
				return errors.New("value cannot be empty")
			}

			return nil
		},
	}

	value, err := prompt.Run()

	return value, err
}
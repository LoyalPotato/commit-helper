package git

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/LoyalPotato/commit-helper/src/cli"
	"github.com/LoyalPotato/commit-helper/src/messages"
)

func AddAlias(key string, value string) {
	fmt.Printf(messages.Git_Alias_Adding, promptui.Styler(promptui.FGFaint)(value))
	UpdateGitConfig("--global", fmt.Sprintf("alias.%s", key), fmt.Sprintf("!%s", value))
	fmt.Println(messages.Git_Alias_Added)
}

func UpdateGitConfig(args ...string) {
	configArgs := []string{"config"}
	configArgs = append(configArgs, args...)
	err := cli.RunCmd("git", configArgs...)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		os.Exit(1)
	}
}

func UnsetAlias(key string) {
	fmt.Printf(messages.Git_Alias_Removing, promptui.Styler(promptui.FGFaint)(key))
	UpdateGitConfig("--global", "--unset", fmt.Sprintf("alias.%s", key))
	fmt.Println(messages.Git_Alias_Removed)
}
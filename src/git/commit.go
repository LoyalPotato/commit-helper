package git

import (
	"fmt"
	"os"
	"strings"

	"github.com/LoyalPotato/commit-helper/src/cli"
	"github.com/LoyalPotato/commit-helper/src/messages"
)

var CommitMsg string

func BuildCommitMsg(msg string) {
	CommitMsg += msg
}

func Commit() {
	args := []string{"commit"}
	splitMsg := strings.Split(CommitMsg, "-m")
	for _, msg := range splitMsg {
		args = append(args, fmt.Sprintf("-m %s", strings.TrimSpace(msg)))
	}
	err := cli.RunCmd("git", args...)
	if err != nil {
		fmt.Println(fmt.Errorf(messages.Error, err))
		os.Exit(1)
	}
}

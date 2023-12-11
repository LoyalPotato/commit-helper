package cli

import (
	"os"
	"os/exec"
)

func RunCmd(cmd string, args ...string) error {
	subProccess := exec.Command(cmd, args...);
	subProccess.Stderr = os.Stderr
	subProccess.Stdout = os.Stdout
	return subProccess.Run()
}
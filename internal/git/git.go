package git

import (
	"os/exec"
)

func GetDiff() (string, error) {
	cmd := exec.Command("git", "diff")

	out, err := cmd.CombinedOutput()

	return string(out), err
}

func GetStatus() (string, error) {
	cmd := exec.Command("git", "status")

	out, err := cmd.CombinedOutput()

	return string(out), err
}

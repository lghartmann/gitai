package git

import (
	"fmt"
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

func CommitChanges(message string) error {
	cmd := exec.Command("git", "commit", "-am", message)

	_, err := cmd.CombinedOutput()

	return err
}

func AddChanges() error {
	cmd := exec.Command("git", "add", ".")

	_, err := cmd.CombinedOutput()

	return err
}

func PushChanges() error {
	fmt.Println("test push")
	cmd := exec.Command("git", "push", "origin", "$(git branch --show-current)")

	_, err := cmd.CombinedOutput()

	return err
}

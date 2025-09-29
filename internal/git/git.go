package git

import (
	"fmt"
	"os/exec"
	"strings"
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
	branchCmd := exec.Command("git", "branch", "--show-current")
	branchOutput, err := branchCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %w", err)
	}

	currentBranch := strings.TrimSpace(string(branchOutput))
	pushCmd := exec.Command("git", "push", "origin", currentBranch)
	output, err := pushCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("push failed: %w\nOutput: %s", err, string(output))
	}

	return err
}

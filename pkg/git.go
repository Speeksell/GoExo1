package pkg

import (
	"fmt"
	"os/exec"
)

func GitFetchAndPull(repoPath string) error {
	cmd := exec.Command("git", "fetch", "--all")
	cmd.Dir = repoPath
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("erreur lors de l'exécution de git fetch: %w", err)
	}

	cmd = exec.Command("git", "pull")
	cmd.Dir = repoPath
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("erreur lors de l'exécution de git pull: %w", err)
	}

	return nil
}

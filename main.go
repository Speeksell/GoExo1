package main

import (
	"RENDU/pkg"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	username := os.Getenv("GITHUB_PSEUDO")
	repos, err := pkg.GetRepositories(username)
	if err != nil {
		panic(err)
	}

	err = pkg.WriteToCSV(repos, "repositories.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Les informations des dépôts ont été écrites dans repositories.csv.")

	// Cloner les dépôts
	for _, repo := range repos[:2] {
		err := cloneRepo(repo.CloneURL)
		if err != nil {
			fmt.Printf("Erreur lors du clonage du dépôt %s: %v\n", repo.Name, err)
		}
		parts := strings.Split(repo.CloneURL, "/")
		repoName := strings.TrimSuffix(parts[len(parts)-1], ".git")
		repoPath := "repositories/" + repoName
		err = pkg.GitFetchAndPull(repoPath)
		if err != nil {
			fmt.Printf("Erreur lors de l'exécution de git fetch et pull pour le dépôt %s: %v\n", repo.Name, err)
		}
		err = pkg.ArchiveRepositories("repositories", "repositories.zip")
		if err != nil {
			panic(err)
		}
		fmt.Println("Les dépôts clonés ont été archivés dans repositories.zip.")
	}
}

func cloneRepo(url string) error {
	// Créer le dossier "repositories" si il n'existe pas déjà
	err := exec.Command("mkdir", "-p", "repositories").Run()
	if err != nil {
		return err
	}

	// Extraire le nom du dépôt de l'URL de clonage
	parts := strings.Split(url, "/")
	repoName := parts[len(parts)-1]

	// Exécuter la commande git clone
	cmd := exec.Command("git", "clone", url, "repositories/"+repoName)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

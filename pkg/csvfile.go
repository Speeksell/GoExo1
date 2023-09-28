package pkg

import (
	"encoding/csv"
	"os"
)

func WriteToCSV(repos []Repository, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Écrire l'en-tête du CSV
	err = writer.Write([]string{"Nom", "URL de Clonage", "Dernière Mise à Jour"})
	if err != nil {
		return err
	}

	// Écrire les données du CSV
	for _, repo := range repos {
		err = writer.Write([]string{repo.Name, repo.CloneURL, repo.UpdatedAt.String()})
		if err != nil {
			return err
		}
	}
	return nil
}

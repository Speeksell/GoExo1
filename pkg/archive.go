package pkg

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ArchiveRepositories(reposDir, archivePath string) error {
	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	zipWriter := zip.NewWriter(archiveFile)
	defer zipWriter.Close()

	err = filepath.Walk(reposDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		infoHeader, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		infoHeader.Name = path
		writer, err := zipWriter.CreateHeader(infoHeader)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

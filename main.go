package main

import (
	"flag"
	"fork-down/custom_errors"
	"fork-down/models"
	data_repo "fork-down/repository"
	"fork-down/restore"
	"fork-down/utils"
	"log"
)

func main() {
	filePath := flag.String("file", "", "path to file")
	manifestPath := flag.String("manifest", "", "path to manifest")
	oldManifestPath := flag.String("oldmanifest", "", "path to old manifest")

	flag.Parse()

	err := utils.ValidateInput(filePath, manifestPath, oldManifestPath)
	if err != nil {
		log.Fatal(err)
	}

	manifest, err := utils.ReadManifest(*manifestPath)

	if err != nil {
		log.Fatal(custom_errors.ErrorReadingManifest)
		return
	}

	manifestForOld, err := utils.ReadManifest(*oldManifestPath)

	fileChunks, err := utils.ToChunks(*filePath, manifestForOld)
	if err != nil {
		return
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	config := models.Config{
		ConfigRepository: models.ConfigRepository{
			Host: "127.0.0.1",
			Port: "1234",
		},
		ConfigRestore: models.ConfigRestore{
			SaveFilePath: "result.bin",
		},
	}

	repository := data_repo.NewS3Repository(&config.ConfigRepository)

	restore := restore.InitRestore(&config.ConfigRestore, repository)

	restore.RestoreFile(fileChunks, manifest)
}

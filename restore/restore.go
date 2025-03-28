package restore

import (
	"fork-down/custom_errors"
	"fork-down/models"
	"fork-down/repository"
	"log"
	"os"
)

type Restore struct {
	config     *models.ConfigRestore
	repository repository.Repository
}

func InitRestore(config *models.ConfigRestore, repo repository.Repository) *Restore {
	// валидация по-хорошему

	return &Restore{
		config:     config,
		repository: repo,
	}
}

func (r *Restore) RestoreFile(fileChunksData map[string][]byte, manifest []models.Chunk) {
	outputFile, err := os.Create(r.config.SaveFilePath)

	if err != nil {
		log.Fatal(custom_errors.ErrorOpenFile)
	}

	var chunksToDownload []models.Chunk

	for _, chunk := range manifest {
		_, isExist := fileChunksData[chunk.Hash]

		if !isExist {
			chunksToDownload = append(chunksToDownload, chunk)
		} else {
			log.Printf("chunk %v founded", chunk.Hash)
		}
	}

	for _, chunk := range chunksToDownload {
		downloadChunk, err := r.repository.DownloadChunk(chunk.Hash)
		log.Printf("chunk %v downloaded", chunk.Hash)

		if err != nil {
			log.Fatal(custom_errors.ErrorDownloadChunk)
		}

		fileChunksData[chunk.Hash] = downloadChunk
	}

	for _, chunk := range manifest {
		chunkData, exists := fileChunksData[chunk.Hash]

		if !exists {
			log.Fatal(custom_errors.FatalError)
		}

		_, err := outputFile.Write(chunkData)
		if err != nil {
			log.Fatal(custom_errors.ErrorWriteFile)
		}
	}
}

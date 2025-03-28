package utils

import (
	"encoding/json"
	"fork-down/models"
	"os"
)

func ReadManifest(filename string) ([]models.Chunk, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var chunks []models.Chunk

	err = json.Unmarshal(data, &chunks)

	if err != nil {
		return nil, err
	}

	return chunks, nil
}

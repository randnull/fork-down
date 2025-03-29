package utils

import (
	"encoding/json"
	"errors"
	"fork-down/models"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ReadManifest(filename string) ([]models.Chunk, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(filename)
	switch ext {
	case ".json":
		var chunks []models.Chunk
		err = json.Unmarshal(data, &chunks)
		if err != nil {
			return nil, err
		}
		return chunks, nil
	case ".rdx":
		return parseRDX(data)
	default:
		return nil, errors.New("unsupported manifest format")
	}
}

func parseRDX(data []byte) ([]models.Chunk, error) {
	s := strings.TrimSpace(string(data))
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		s = s[1 : len(s)-1]
	}
	var chunks []models.Chunk
	entries := strings.Split(s, "«")
	for _, entry := range entries {
		entry = strings.TrimSpace(entry)
		if entry == "" {
			continue
		}
		if !strings.HasSuffix(entry, "»") {
			return nil, errors.New("invalid rdx format: missing close del")
		}
		entry = entry[:len(entry)-len("»")]
		parts := strings.SplitN(entry, ":", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid rdx format: missing colon del")
		}
		sizeStr := strings.TrimSpace(parts[0])
		hashStr := strings.TrimSpace(parts[1])
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			return nil, err
		}
		chunks = append(chunks, models.Chunk{Size: size, Hash: hashStr})
	}
	return chunks, nil
}


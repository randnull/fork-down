package utils

import (
	"bufio"
	"errors"
	"fork-down/custom_errors"
	"fork-down/models"
	"io"
	"os"
)

func ToChunks(filePath string, manifest []models.Chunk) (map[string][]byte, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, custom_errors.ErrorWithReadFile
	}

	defer file.Close()

	chunks := make(map[string][]byte)

	reader := bufio.NewReader(file)

	for _, chunk := range manifest {
		buffer := make([]byte, chunk.Size)

		chunkData, err := io.ReadFull(reader, buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			if errors.Is(err, io.ErrUnexpectedEOF) {
				if chunkData > 0 {
					data := make([]byte, chunkData)

					copy(data, buffer)

					buffer = data
				}
				break
			}
			return nil, custom_errors.ErrorWithReadFile
		}

		data := make([]byte, chunkData)

		copy(data, buffer)

		hash := Sha256Hash(data)

		chunks[hash] = data
	}

	return chunks, nil
}

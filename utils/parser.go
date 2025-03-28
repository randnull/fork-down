package utils

import (
	"bufio"
	"errors"
	"fork-down/custom_errors"
	"io"
	"os"
)

func ToChunks(filePath string) (map[string][]byte, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, custom_errors.ErrorWithReadFile
	}

	defer file.Close()

	chunkSize := 1024 * 64 * 2

	chunks := make(map[string][]byte)

	buffer := make([]byte, chunkSize)

	reader := bufio.NewReader(file)

	for {
		n, err := io.ReadFull(reader, buffer[:cap(buffer)])
		buffer = buffer[:n]

		if err != nil {
			if err == io.EOF {
				break
			}
			if errors.Is(err, io.ErrUnexpectedEOF) {
				if n > 0 {
					data := make([]byte, n)

					copy(data, buffer)

					hash := Sha256Hash(data)

					chunks[hash] = data
				}
				break
			}
			return nil, custom_errors.ErrorWithReadFile
		}

		data := make([]byte, n)

		copy(data, buffer)

		hash := Sha256Hash(data)

		chunks[hash] = data
	}

	return chunks, nil
}

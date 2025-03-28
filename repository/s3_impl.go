package repository

import "fork-down/models"

type S3 struct {
	config *models.ConfigRepository
}

func NewS3Repository(config *models.ConfigRepository) *S3 {
	// тут пинг, ждем команду по S3

	return &S3{
		config: config,
	}
}

func (s *S3) DownloadChunk(hash string) ([]byte, error) {
	return []byte("mockData"), nil
}

package repository

type Repository interface {
	DownloadChunk(hash string) ([]byte, error)
}

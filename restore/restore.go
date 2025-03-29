package restore

import (
 "fork-down/custom_errors"
 "fork-down/models"
 "fork-down/repository"
 "log"
 "os"
 "sync"
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

 var wg sync.WaitGroup

 chunksToDownload := make(chan models.Chunk, len(manifest))

 for _, chunk := range manifest {
  _, isExist := fileChunksData[chunk.Hash]

  if !isExist {
   chunksToDownload <- chunk
  } else {
   log.Printf("chunk %v founded", chunk.Hash)
  }
 }

 close(chunksToDownload)

 var mu sync.Mutex

 downloadChunks := make(map[string][]byte)

 worker := func(workerNum int) {
  defer wg.Done()

  for chunk := range chunksToDownload {
   downloadChunk, err := r.repository.DownloadChunk(chunk.Hash)
   if err != nil {
    log.Fatal(custom_errors.ErrorDownloadChunk)
   }
   log.Printf("chunk %v downloaded (%v)", chunk.Hash, workerNum)
   mu.Lock()
   downloadChunks[chunk.Hash] = downloadChunk
   mu.Unlock()
  }
 }

 for i := 0; i < 3; i++ {
  wg.Add(1)
  go worker(i)
 }

 wg.Wait()

 for _, chunk := range manifest {
  var data []byte

  chunkData, exists := fileChunksData[chunk.Hash]
  if !exists {
   chunkDataDownloaded, existsDownload := downloadChunks[chunk.Hash]
   if !existsDownload {
    log.Fatal(custom_errors.FatalError)
   }
   data = chunkDataDownloaded
  } else {
   data = chunkData
  }

  _, err := outputFile.Write(data)
  if err != nil {
   log.Fatal(custom_errors.ErrorWriteFile)
  }
 }
}

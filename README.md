# fork-down
fork-down given a manifest and an S3 storage URI, downloads a blob chunk by chunk

Использование:
```
go mod tidy

go build -o fork-down main.go

./fork-down -file "test15.bin" -manifest "manifest.json"
```


Пример тестирования:

```
➜  fork-down1 git:(main) ✗ ./test.sh                                                        
testing manifest similar to .bin
2025/03/28 17:04:48 chunk d2e033a85b7585ff4748efe61cff688a0934a838ada3de431c43fa33e96b058d founded
2025/03/28 17:04:48 chunk ae2e87d3ed0c1691de4bbd511672fbd2bcfa4f4d80ced8f58413336c9b1929dc founded
2025/03/28 17:04:48 chunk ed64d5aa63d9c8da562e8fa8111d5a3ec65e383c2bba30b0ed97032580f7816d founded
2025/03/28 17:04:48 chunk 6d4d4e4c0a1198dd8bd7a1d33ad07bd02e903ee82f7d91c64f9cb4c575346dc9 founded
testing manifest not similar to .bin
2025/03/28 17:04:48 chunk ae2e87d3ed0c1691de4bbd511672fbd2bcfa4f4d80ced8f58413336c9b1929dc founded
2025/03/28 17:04:48 chunk ed64d5aa63d9c8da562e8fa8111d5a3ec65e383c2bba30b0ed97032580f7816d founded
2025/03/28 17:04:48 chunk 6d4d4e4c0a1198dd8bd7a1d33ad07bd02e903ee82f7d91c64f9cb4c575346dc9 founded
2025/03/28 17:04:48 chunk d2e033a85b7585ff4748efe61cff688a0934a838ada3de431c43fa33e96b058d downloaded
```

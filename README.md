# file

## Description

## How to run
From the main directory, run:
### Windows
```bash
go build -o file.exe
./file.exe
```

### Linux
```bash
go build -o file
./file
```

Then upload a file to localhost:8080/upload as a multipart form with the key "file".

## How to check metrics
Install Prometheus, then run the following commands:
```bash
prometheus --config.file=prometheus.yml
```

### Available metrics

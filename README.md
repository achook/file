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

### Docker
```bash
docker build -t file .
docker run -p 8080:8080 file
```

Then upload a file to localhost:8080/upload as a multipart form with the key "file".

## How to check metrics
Install Prometheus, then run the following commands:
```bash
prometheus --config.file=prometheus.yml
```

### Available metrics
- csv_file_count: Number of CSV files uploaded
- json_file_count: Number of JSON files uploaded
- text_file_count: Number of text files uploaded
- response_time: Response time of the server
- file_size: Size of the uploaded file

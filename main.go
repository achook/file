package main

import (
	"bufio"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a new file on the server
	filename := handler.Filename
	fileType := getFileType(filename)

	var stats Stats
	rdr := bufio.NewReader(file)

	switch fileType {
	case TypeCsv:
		// Parse the CSV file
		stats, err = parseCsv(rdr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case TypeJson:
		// Parse the JSON file
		stats, err = parseJSON(rdr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		// Parse the text file
		stats, err = parseText(rdr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Write the stats to the response
	_, err = w.Write([]byte(fmt.Sprintf("Stats: %+v", stats)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	http.HandleFunc("/upload", uploadFileHandler)
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

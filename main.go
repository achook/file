package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		csvFilesCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "csv_files_count",
			Help: "Number of uploaded CSV files",
		}, []string{"status"}),
		jsonFilesCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "json_files_count",
			Help: "Number of uploaded JSON files",
		}, []string{"status"}),
		textFilesCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "text_files_count",
			Help: "Number of uploaded text files",
		}, []string{"status"}),
		responseTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "response_time",
			Help:    "Response time of the server",
			Buckets: prometheus.DefBuckets,
		}, []string{"status"}),
		fileSize: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "file_size",
			Help:    "Size of the uploaded file",
			Buckets: prometheus.DefBuckets,
		}, []string{"status"}),
	}
	reg.MustRegister(m.csvFilesCount)
	reg.MustRegister(m.jsonFilesCount)
	reg.MustRegister(m.textFilesCount)
	reg.MustRegister(m.responseTime)
	reg.MustRegister(m.fileSize)

	return m
}

type Server struct {
	m *Metrics
}

func (s *Server) uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form data
	start := time.Now()
	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		end := time.Now()
		elapsed := end.Sub(start)
		s.m.responseTime.WithLabelValues("error").Observe(float64(elapsed.Milliseconds()))
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		end := time.Now()
		elapsed := end.Sub(start)
		s.m.responseTime.WithLabelValues("error").Observe(float64(elapsed.Milliseconds()))
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
			end := time.Now()
			elapsed := end.Sub(start)
			s.m.responseTime.WithLabelValues("error").Observe(float64(elapsed.Milliseconds()))
			return
		}
		s.m.csvFilesCount.WithLabelValues("success").Inc()
	case TypeJson:
		// Parse the JSON file
		stats, err = parseJSON(rdr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			end := time.Now()
			elapsed := end.Sub(start)
			s.m.responseTime.WithLabelValues("error").Observe(float64(elapsed.Milliseconds()))
			return
		}
		s.m.jsonFilesCount.WithLabelValues("success").Inc()
	default:
		// Parse the text file
		stats, err = parseText(rdr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			end := time.Now()
			elapsed := end.Sub(start)
			s.m.responseTime.WithLabelValues("error").Observe(float64(elapsed.Milliseconds()))
			return
		}
		s.m.textFilesCount.WithLabelValues("success").Inc()
	}

	// Write the stats to the response as JSON
	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		end := time.Now()
		elapsed := end.Sub(start)
		s.m.responseTime.WithLabelValues("error").Observe(float64(elapsed.Milliseconds()))
		return
	}

	end := time.Now()
	elapsed := end.Sub(start)
	s.m.responseTime.WithLabelValues("success").Observe(float64(elapsed.Milliseconds()))

}

func main() {
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)
	s := &Server{m: m}

	http.HandleFunc("/upload", s.uploadFileHandler)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

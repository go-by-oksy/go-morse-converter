package handlers

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-by-oksy/go-morse-converter/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "failed to read uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read file contents", http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".txt"
	}

	filename := "converted-" +
		time.Now().UTC().Format("20060102-150405.000000000") +
		ext

	if err := os.WriteFile(filename, []byte(result), 0644); err != nil {
		http.Error(w, "failed to save converted file", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("result.html")
	if err != nil {
		http.Error(w, "failed to load result page", http.StatusInternalServerError)
		return
	}

	pageData := struct {
		Result   string
		Filename string
	}{
		Result:   result,
		Filename: filename,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tmpl.Execute(w, pageData); err != nil {
		http.Error(w, "failed to render result page", http.StatusInternalServerError)
	}
}

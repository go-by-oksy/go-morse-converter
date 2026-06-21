func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var file multipart.File
	var header *multipart.FileHeader

	for _, files := range r.MultipartForm.File {
		if len(files) > 0 {
			header = files[0]
			file, err = header.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			break
		}
	}

	if file == nil {
		http.Error(w, "file not found", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	filename := time.Now().UTC().Format("20060102150405") + ext

	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
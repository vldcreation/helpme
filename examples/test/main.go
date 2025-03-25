package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var rootDir string

func init() {
	var err error
	rootDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	rootDir += "/Documents/Work"
}

func main() {
	// Check if the root directory exists
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		fmt.Printf("Directory %s does not exist\n", rootDir)
		os.Exit(1)
	}

	http.HandleFunc("/", fileHandler)

	port := ":8000"
	fmt.Printf("Serving %s on http://%s%s\n", rootDir, getLocalIP(), port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

// Function to handle file requests
func fileHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(rootDir, r.URL.Path)
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	if fileInfo.IsDir() {
		if r.URL.Query().Get("download") == "true" {
			compressAndDownloadDir(w, path)
		} else {
			dirList(w, path)
		}
	} else {
		http.ServeFile(w, r, path)
	}
}

// Function to list files in a directory
func dirList(w http.ResponseWriter, dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	relativePath, err := filepath.Rel(rootDir, dirPath)
	if err != nil {
		http.Error(w, "Error creating relative path", http.StatusInternalServerError)
		return
	}
	if relativePath == "." {
		relativePath = ""
	}

	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			body { font-family: Arial, sans-serif; margin: 20px; }
			h1 { color: #333; }
			.nav-links { margin: 10px 0; }
			.nav-links a { margin-right: 10px; }
			.file-list { list-style: none; padding: 0; }
			.file-item { display: flex; align-items: center; padding: 5px 0; }
			.file-link { text-decoration: none; color: #0066cc; margin-right: 15px; }
			.download-btn { 
				padding: 3px 10px;
				background-color: #4CAF50;
				color: white;
				border: none;
				border-radius: 3px;
				text-decoration: none;
				font-size: 0.9em;
			}
			.download-btn:hover { background-color: #45a049; }
		</style>
	</head>
	<body>
	`)

	fmt.Fprintf(w, "<h1>Directory: /%s</h1>", relativePath)

	fmt.Fprintf(w, "<div class=\"nav-links\">")
	if dirPath != rootDir {
		parentPath := "/" + filepath.ToSlash(filepath.Dir(relativePath))
		if parentPath == "/"+"." {
			parentPath = "/"
		}
		fmt.Fprintf(w, `<a href="%s">Back</a>`, parentPath)
	}
	fmt.Fprintf(w, `<a href="/">Root</a></div>`)

	fmt.Fprintf(w, "<ul class=\"file-list\">")
	for _, file := range files {
		name := file.Name()
		isDir := file.IsDir()
		if isDir {
			name += "/"
		}
		relativePath, _ := filepath.Rel(rootDir, filepath.Join(dirPath, name))
		slashedPath := filepath.ToSlash(relativePath)

		fmt.Fprintf(w, "<li class=\"file-item\">")
		fmt.Fprintf(w, `<a class="file-link" href="/%s">%s</a>`, slashedPath, name)
		downloadParam := ""
		if isDir {
			downloadParam = "?download=true"
		}
		fmt.Fprintf(w, `<a class="download-btn" href="/%s%s" download>Download</a>`, slashedPath, downloadParam)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</ul></body></html>")
}

// Function to compress a directory and send it as a zip file
func compressAndDownloadDir(w http.ResponseWriter, dirPath string) {
	relativePath, err := filepath.Rel(rootDir, dirPath)
	if err != nil {
		http.Error(w, "Error creating relative path", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", relativePath))

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if it's the root directory itself
		if path == dirPath {
			return nil
		}

		// Create a relative path for the zip file
		relPath, err := filepath.Rel(dirPath, path)
		if err != nil {
			return err
		}

		// Create zip header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		http.Error(w, "Error creating zip file", http.StatusInternalServerError)
		return
	}
}

// Function to get the local IP address using net.InterfaceAddrs()
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost" // Fallback if there's an error
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String() // Return the first valid non-loopback IPv4 address
		}
	}
	return "localhost" // Fallback if no valid IP found
}

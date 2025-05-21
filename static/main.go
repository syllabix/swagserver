package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a zip reader from the byte slice
	reader, err := zip.NewReader(bytes.NewReader([]byte(data)), int64(len(data)))
	if err != nil {
		panic(err)
	}

	// Iterate through each file in the archive
	for _, file := range reader.File {
		fmt.Printf("Extracting: %s\n", file.Name)
		rc, err := file.Open()
		if err != nil {
			panic(err)
		}
		defer rc.Close()

		// Create the output file (or directory)
		if file.FileInfo().IsDir() {
			os.MkdirAll(file.Name, file.Mode())
		} else {
			os.MkdirAll(getDir(file.Name), 0755) // Ensure parent directories exist
			outFile, err := os.OpenFile(file.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				panic(err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				panic(err)
			}
		}
	}
}

// Helper to get directory part of a path
func getDir(path string) string {
	if idx := len(path) - 1; idx >= 0 && (path[idx] == '/' || path[idx] == '\\') {
		return path
	}
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			return path[:i]
		}
	}
	return ""
}

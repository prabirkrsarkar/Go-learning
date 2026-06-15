package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// To start incrementally we will first add a function that will download
// a file from the internet. Then we will add concurrency on top of it.

func createDirs(destDir string) (string, error) {

	// Define the full filepath as join of internet filename and destDir
	// Base(url) returns the last parth of the url, which is the filename

	// Unlike Shell, GO doesn't understand ~ so we have to get the homedir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fulldestDir := filepath.Join(homeDir, destDir)

	// Let's create the entire dir path if it doesn't exist. os.Create() can't create Parent dirs.
	if err := os.MkdirAll(fulldestDir, 0755); err != nil {
		return "", fmt.Errorf("Create Destination dir failed!")
	}
	return fulldestDir, nil
}

func downloadFile(url, destDir string) error {

	fileName := filepath.Base(url)
	filePath := filepath.Join(destDir, fileName)

	// Create the file, open it for writing and return a file handle.
	fileHandle, err := os.Create(filePath)
	if err != nil {
		return err
	}

	// If we call fileHandle.Close() before checking for error and if error happens then fileHandle
	// will be nil, calling a Close() will then result in nil-pointer panic. Hence check error first.
	defer fileHandle.Close()

	fmt.Printf("Downloading ... %s", url)
	start := time.Now()

	// Do a GET to the specified URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status not OK! %s", resp.Status)
	}

	_, err = io.Copy(fileHandle, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Download %s took %s\n", fileName, time.Since(start))
	return nil
}

//func Sequentialdownloader(url []string, destDir string) error {}

func main() {
	url := "https://images.unsplash.com/photo-1780995175154-a8c5ae56a230"
	fulldestDir, _ := createDirs("GoDownloads")

	err := downloadFile(url, fulldestDir)

	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}
	log.Println("Done!")
}

package main

import (
	"fmt"
	"os"
	"time"
)

// Filesystem
func getLatestFile(dirpath string) (string, error) {

	if len(dirpath) == 0 {
		return "", fmt.Errorf("Path can not be empty")
	}

	dirs, err := os.ReadDir(dirpath)
	if err != nil {
		return "", err
	}

	var newestFile os.DirEntry
	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}

		if newestFile == nil {
			newestFile = dir
		}

		currentInfo, _ := dir.Info()
		newestInfo, _ := newestFile.Info()
		if currentInfo.ModTime().After(newestInfo.ModTime()) {
			newestFile = dir
		}
	}
	return newestFile.Name(), nil
}

func GetFileContent(filename string, startAt int) (string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if len(f) < startAt+1 {
		return "", nil
	}

	toReturn := f[startAt:]

	return string(toReturn), nil
}

func loopPrintFile(State *state, ch chan string) error {
	fileLength := 0
	lastFilepath := State.currentFilePath
	for {
		if lastFilepath != State.currentFilePath {
			fileLength = 0
			lastFilepath = State.currentFilePath
		}

		content, _ := GetFileContent(State.currentFilePath, fileLength)
		fileLength += len(content)
		if len(content) > 0 {
			fmt.Println("Content sent to channel")
			ch <- content
		}
		ticker := time.NewTicker(100 * time.Millisecond)
		<-ticker.C
	}
}

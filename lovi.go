package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args
	amountArgs := len(args)
	if amountArgs < 2 {
		fmt.Println("usage: lovi <Name from config file>")
		os.Exit(1)
	}

	dirpath, err := GetpathFromConfig(args[1])
	if err != nil {
		fmt.Printf("Config file does not contain %v", args[1])
		os.Exit(1)
	}

	fileName, err := GetLatestFile(dirpath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	loopPrintFile(dirpath + fileName)

}

// Config file
type config struct {
	Folders []struct {
		Name     string `json:"name"`
		Filepath string `json:"filepath"`
	} `json:"folders"`
}

func GetpathFromConfig(name string) (string, error) {
	configFileName := "/lovi.config"
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	file := path + configFileName
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	var Config config

	if err = json.Unmarshal(content, &Config); err != nil {
		return "", err
	}

	for _, c := range Config.Folders {
		if c.Name == name {
			return c.Filepath, nil
		}
	}

	return "", fmt.Errorf("filepath not found in config")
}

// Filesystem
func GetLatestFile(dirpath string) (string, error) {

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

func printFile(filename string, startAt int) (int, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	if len(f) < startAt+1 {
		return len(f), nil
	}

	toPrint := f[startAt:]

	fmt.Print(string(toPrint))

	return len(f), nil
}

func loopPrintFile(filename string) error {
	fileLength := 0
	for {
		fileLength, _ = printFile(filename, fileLength)
		ticker := time.NewTicker(100 * time.Millisecond)
		<-ticker.C
	}
}

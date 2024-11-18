package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	filepath, err := GetpathFromConfig("flex")
	if err != nil {
		fmt.Println("Filepath not found in config file")
		os.Exit(1)
	}

	fmt.Println(filepath)
	fmt.Println(GetLatestFile(filepath))

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

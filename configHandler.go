package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config file
type config struct {
	Folders []struct {
		Name     string `json:"name"`
		Filepath string `json:"filepath"`
		Hotkey   string `json:"hotkey,omitempty"`
	} `json:"folders"`
}

func setConfigFromFile(conf *config, configFileName string) error {
	path, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	file := filepath.Join(path, configFileName)
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Could not read file")
		return err
	}

	if err = json.Unmarshal(content, conf); err != nil {
		fmt.Println("Could not unmarshal json")
		return err
	}

	return nil
}

func (conf *config) getPath(name string) (string, error) {
	for _, c := range conf.Folders {
		if c.Name == name {
			return c.Filepath, nil
		}
		fmt.Println(c)
	}

	return "", fmt.Errorf("Config file does not contain %v", name)
}

func (conf *config) getPathFromHotkey(key rune) (string, error) {
	for _, c := range conf.Folders {
		if c.Hotkey == string(key) {
			return c.Filepath, nil
		}
	}
	return "", fmt.Errorf("Hotkey not found")
}

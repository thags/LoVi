package main

import (
	"fmt"
	"github.com/mattn/go-tty"
	"os"
	"path/filepath"
)

type state struct {
	conf            *config
	currentFilePath string
}

func main() {
	configFileName := "lovi.config"
	args := os.Args
	amountArgs := len(args)
	if amountArgs < 2 {
		fmt.Println("usage: lovi <Name from config file>")
		os.Exit(1)
	}

	var Config config
	err := setConfigFromFile(&Config, configFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dirpath, err := Config.getPath(args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileName, err := getLatestFile(dirpath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	State := state{
		conf:            &Config,
		currentFilePath: filepath.Join(dirpath, fileName),
	}

	go handleUserKeyInput(&State)

	fileChan := make(chan string)
	go loopPrintFile(&State, fileChan)

	for {
		content := <-fileChan
		fmt.Print(content)
	}

}

func handleUserKeyInput(State *state) {
	tty, err := tty.Open()
	if err != nil {
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			continue

		}
		switch r {
		case 'q':
			os.Exit(0)
		default:
			path, err := State.conf.getPathFromHotkey(r)
			if err != nil {
				continue
			}
			file, err := getLatestFile(path)
			if err != nil {
				continue
			}
			State.currentFilePath = filepath.Join(path, file)
		}
	}
}

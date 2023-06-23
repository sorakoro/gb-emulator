package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ExitCodeOK = iota
	ExitCodeErr
)

func main() {
	os.Exit(Run())
}

func Run() int {
	flag.Parse()
	romPath := flag.Arg(0)
	romData, err := readROM(romPath)
	fmt.Println(romData)
	if err != nil {
		return ExitCodeErr
	}
	return ExitCodeOK
}

func readROM(path string) ([]byte, error) {
	if path == "" {
		return []byte{}, errors.New("")
	}
	if filepath.Ext(path) != ".GB" && filepath.Ext(path) != ".GBC" {
		return []byte{}, errors.New("")
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, errors.New("")
	}
	return bytes, err
}

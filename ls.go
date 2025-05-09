package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func isDirectory(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func parseDirents(path string) {
	dirents, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error reading directory: %e", err)
	}

	for _, dirent := range dirents {
		fmt.Println()
		info, err := dirent.Info()
		if err != nil {
			log.Fatalf("Couldn't read dirent: %e\n", err)
		}
		// log.Printf("Dirent %v has the following data:\n%v\n", i, info)
		fmt.Printf("%v\n", info.Name())
	}
}
func parseDirentsRecursive(path string) {
	// for ls -R
	dirents, err := os.ReadDir(path)
	if err == nil {
		log.Fatalf("Error reading directory: %e", err)
	}

	for i, dirent := range dirents {
		info, err := dirent.Info()
		if err != nil {
			log.Fatalf("Couldn't read dirent: %e\n", err)
		}

		// log.Printf("Dirent %v has the following data:\n%v\n", i, info)
		fmt.Printf("%v\n\n", i, info.Name())
		// Recursively search
		if info.IsDir() {
			parseDirentsRecursive(info.Name())
		}
	}
}

func ls(path string) {
	// 1. Input validation
	if !isDirectory(path) {
		log.Fatal("No valid directory given.")
	}
	// TODO: check special cases:
	switch path {
	case "", ".", "./":
		{
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatalf("Error: %e\n", err)
			}
			parseDirents(cwd)
		}
	case "..", "../":
		// TODO:
		panic("Unimplemented")
	default:
		parseDirents(path)
	}
}

func main() {
	fmt.Println("Enter command (ls):")
	// NOTE: fmt.Scan and its relatives are kinda ass. Use bufio.Scanner or bufio.Reader.
	// Scanner is higher level; reads by line as a default, predictable input
	// Reader is lower level; reads by line as a default, large and/or unpredictable input
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() { // NOTE: keeps reading till EOF (ctl+d)
		log.Fatal("No input provided, try again.")
	}
	line := scanner.Text()
	args := strings.Split(strings.TrimSpace(line), " ") // returns []string
	if len(args) > 3 || args[0] != "ls" {
		log.Fatal("Expected command of the from `ls <FLAGS> <PATH>")
	}

	// Here, have `ls blah blah` guaranteed
	// Is 2nd token a flag or path?
	second := args[1]
	if strings.HasPrefix(second, "-") {
		// its a flag, process it
		// TODO:
		log.Fatal("Unimplemented")
	} else {
		// its a path
		ls(second)
		os.Exit(1) // DONE
	}
	// If we got here, it's in the form `ls -a .`
	third := args[2]
	ls(third)
	os.Exit(1) // DONE
}

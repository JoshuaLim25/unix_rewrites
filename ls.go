package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func isDir(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
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

	for i, dirent := range dirents {
		log.Println()
		info, err := dirent.Info()
		if err != nil {
			log.Fatalf("Couldn't read dirent: %e\n", err)
		}

		// log.Printf("Dirent %v has the following data:\n%v\n", i, info)
		log.Printf("Dirent %d: \n%v\n\n", i, info.Name())
	}
}
func parseDirentsRecursive(path string) {
	dirents, err := os.ReadDir(path)
	if err == nil {
		log.Fatalf("Error reading directory: %e", err)
	}

	for i, dirent := range dirents {
		log.Println()
		info, err := dirent.Info()
		if err != nil {
			log.Fatalf("Couldn't read dirent: %e\n", err)
		}

		// log.Printf("Dirent %v has the following data:\n%v\n", i, info)
		log.Printf("Dirent %d: \n%v\n\n", i, info.Name())
		// Recursively search
		if info.IsDir() {
			parseDirentsRecursive(info.Name())
		}
	}
}

func ls(path string) {
	// 1. Input validation
	if !isDir(path) {
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
	args := make([]string, 3)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		args = append(args, scanner.Text())
		log.Printf("args: %v\n", args)
		if len(args) <= 0 {
			log.Fatal("Error: expected more than one argument.")
		} else if len(args) == 1 {
			ls(".")
		} else if len(args) == 2 { // flag given
			// TODO:
			panic("Unimplemented")
		} else if len(args) == 3 {
			// TODO:
			panic("Unimplemented")
		} else {
			log.Fatal("Error: received more than 3 arguments.")
		}
	}
}

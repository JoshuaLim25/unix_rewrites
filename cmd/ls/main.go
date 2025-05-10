package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// NOTE: this predates my knowedge of os.Args and `flags`. Won't do this bufio scanning business for a script just trollin around
// NOTE: this only covers
/*
ls
ls ./
ls ../
ls dir/
flags: -a or -R
*/

func dirExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func parseDirents(path string) {
	dirents, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	fmt.Println()
	for _, dirent := range dirents {
		info, err := dirent.Info()
		if err != nil {
			log.Fatalf("Couldn't read dirent: %v\n", err)
		}
		if strings.HasPrefix(info.Name(), ".") {
			continue
		} else {
			fmt.Printf("%v\n", info.Name())
		}
	}
}

func parseDirentsAll(path string) { // -a
	dirents, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	fmt.Println()
	for _, dirent := range dirents {
		info, err := dirent.Info()
		if err != nil {
			log.Fatalf("Couldn't read dirent: %v\n", err)
		}
		fmt.Printf("%v\n", info.Name())
	}
}

func parseDirentsRecursive(path string) { // -R
	dirents, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	fmt.Printf("%s:\n", path)
	for _, dirent := range dirents {
		info, err := dirent.Info()
		if err != nil {
			log.Fatalf("Couldn't read dirent: %v\n", err)
		}
		if strings.HasPrefix(info.Name(), ".") {
			continue
		} else if !info.Mode().IsDir() {
			fmt.Printf("%v\n\n", filepath.Join(path, info.Name()))
		} else {
			deeperPath := filepath.Join(path, info.Name())
			fmt.Printf("%v\n\n", deeperPath)
			parseDirentsRecursive(deeperPath)
		}
	}
}

func ls(path string) {
	// 1. Input validation
	if !dirExists(path) {
		log.Fatal("No valid directory given.")
	}
	// TODO: check special cases:
	switch path {
	case "", ".", "./":
		{
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatalf("Error: %v\n", err)
			}
			parseDirents(cwd)
		}
	case "..", "../":
		{
			// TODO:
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatalf("Err: %v\n", err)
			}
			parentDir := filepath.Dir(cwd)
			parseDirents(parentDir)
		}
	case "/":
		log.Fatal("Unimplemented")
	default:
		parseDirents(path)
	}
}

func lsAll(path string) {
	// 1. Input validation
	if !dirExists(path) {
		log.Fatal("No valid directory given.")
	}
	// TODO: check special cases:
	switch path {
	case "", ".", "./":
		{
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatalf("Error: %v\n", err)
			}
			parseDirentsAll(cwd)
		}
	case "..", "../":
		{
			// TODO:
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatalf("Err: %v\n", err)
			}
			parentDir := filepath.Dir(cwd)
			parseDirentsAll(parentDir)
		}
	case "/":
		log.Fatal("Unimplemented")
	default:
		parseDirentsAll(path)
	}
}

func lsRecursive(path string) {
	// 1. Input validation
	if !dirExists(path) {
		log.Fatal("No valid directory given.")
	}
	// TODO: check special cases:
	switch path {
	case "", ".", "./":
		{
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatalf("Error: %v\n", err)
			}
			parseDirentsRecursive(cwd)
		}
	case "..", "../":
		{
			// TODO:
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatalf("Err: %v\n", err)
			}
			parentDir := filepath.Dir(cwd)
			parseDirentsRecursive(parentDir)
		}
	case "/":
		log.Fatal("Unimplemented")
	default:
		parseDirentsRecursive(path)
	}
}

func main() {
	fmt.Println("Enter an ls command:")
	// NOTE: fmt.Scan and its relatives are kinda ass. Use bufio.Scanner or bufio.Reader.
	// Scanner is higher level; reads by line as a default, predictable input
	// Reader is lower level; reads by line as a default, large and/or unpredictable input
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() { // NOTE: keeps reading till EOF (ctl+d)
		log.Fatal("No input provided, try again.")
	}
	line := scanner.Text()
	args := strings.Split(strings.TrimSpace(line), " ") // returns []string
	if len(args) == 0 || len(args) > 3 || args[0] != "ls" {
		log.Fatal("Expected command of the from `ls <FLAGS> <PATH>")
	}

	// len(args) guaranteed in (0, 3]
	switch l := len(args); l {
	case 1: // plain old ls
		ls(".")
	case 2: // I have to duplicate lol
		{
			second := args[1]
			if strings.HasPrefix(second, "-") {
				// It's a flag, process it
				switch second {
				case "-a":
					lsAll(".")
				case "-R":
					lsRecursive(".")
				default:
					ls(".")
				}
				// TODO:
			} else {
				// It's a path
				ls(second)
			}
		}
	case 3: // 2 or 3 args supplied. Need switch to access `second` safely.
		{
			second := args[1]
			if strings.HasPrefix(second, "-") {
				// It's a flag, process it
				dir := args[2]
				switch second {
				case "-a":
					lsAll(dir)
				case "-R":
					lsRecursive(dir)
				default:
					ls(dir)
				}
				// TODO:
			} else {
				// It's a path
				ls(second)
			}
		}
	}
}

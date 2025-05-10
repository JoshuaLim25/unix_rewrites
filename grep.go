package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
grep -iRw ++ .
*/

func patternMatch(path, pattern string) {
	fp, err := os.OpenFile(path, os.O_RDONLY, 444)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			fmt.Println(line)
		}
	}
}

func main() {
	numArgs := len(os.Args)
	fmt.Println(os.Args)
	fmt.Printf("Num args given: %d\n", numArgs)
	if numArgs <= 1 || numArgs > 4 || os.Args[0] != "./g" {
		log.Fatal("Expected argument in the form `grep <flags> <pattern> <path>`")
	}

	if numArgs == 2 {
		// TODO: flag or path?
		if strings.HasPrefix(os.Args[1], "-") {
		}
		log.Fatal("Unimplemented")
	}

	patternMatch("test.txt", "Alice")
	fmt.Println("Hi i still work :)")
}

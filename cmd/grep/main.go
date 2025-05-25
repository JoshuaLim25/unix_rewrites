package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
	// "flag"
	"fmt"
	// "regexp"
)

// Helpers
func PrintUsage() {
	log.Fatal("Usage: grep [OPTION]... PATTERNS [FILE]...")
}

// Typedefs
type Stack []rune // could make generic, also no need to tho
type Pattern string

// Interface impls and methods
// NOTE: errors.New() should be used instead
// func (s *Stack) Error() string {
// 	return fmt.Sprintf("Error popping from empty stack")
// }

func (s *Stack) Push(elem rune) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() (rune, error) {
	if len(*s) == 0 {
		// return 0, s.Error() // WARN: can't do this, would need to return 0, s
		return 0, errors.New("Error popping from empty stack")
	}
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last, nil
}


// Search thru file for a match and return matched line
func PatternMatch(pat Pattern, file string) ([]string, error) {
	// if pat.IsValid() {
	//   return nil, errors.New("Invalid regex provided.")
	// }
	fp, err := os.Open(file)
	if err != nil {
	  return nil, err
	}
	matchedLines := []string{}
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, string(pat)) {
			matchedLines = append(matchedLines, line)
		}
	}
	return matchedLines, nil
}

func ProcessFiles(pat Pattern, files []string) {
	for _, file := range files {
		matches, _ := PatternMatch(pat, file)
		fmt.Println(matches)
	}
}


func main() {
	files := []string{}
	testPattern := Pattern("Bob")
	for i := 1; i <= 10; i++ {
		files = append(files, fmt.Sprintf("test-%d.txt", i))
	}
	fmt.Printf("DEBUGPRINT[72]: main.go:77: files=%+v\n", files)
	ProcessFiles(testPattern, files)
}

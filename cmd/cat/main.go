package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func printUsage() {
	log.Fatal("Usage: cat [OPTION]... [FILE]...\n")
}

func readFiles1(files []string) {
	for _, file := range files {
		fp, err := os.Open(file)
		if err != nil {
			log.Fatalf("Error reading file: %v\n", err)
		}
		defer fp.Close()
		scanner := bufio.NewScanner(fp)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Fprintln(os.Stdout, line)
		}
	}
}

func readFiles2(files []string) {
	for _, file := range files {
		fp, err := os.Open(file)
		defer fp.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		}
		if _, err := io.Copy(os.Stdout, fp); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		}
	}
}

func main() {
	switch len(os.Args) {
	case 0:
		printUsage()
	case 1:
		{
			// no args (cat)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Fprintln(os.Stdout, line)
			}
		}
	case 2:
		{
			arg2 := os.Args[1]
			if strings.HasPrefix(arg2, "-") {
				printUsage()
			}
			if _, err := os.Stat(arg2); err != nil {
				log.Fatalf("Error reading file: %v\n", err)
			}
			file := []string{arg2}
			readFiles2(file)
		}
	default:
		// Two opts: of form `cat -c files` or `cat files`
		arg2 := os.Args[1]
		if strings.HasPrefix(arg2, "-") {
			log.Fatal("Unimplemented: handle flag")
		}
		readFiles2(os.Args[1:])
	}
}

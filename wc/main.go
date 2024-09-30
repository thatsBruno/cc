package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filePathPtr := flag.String("c", "", "file path")

	flag.Parse()

	file, err := os.Stat(*filePathPtr)
	if err != nil {
		log.Fatal(err)
	}

	fileLines, err := os.Open(*filePathPtr) // replace with your file name
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileLines.Close()

	scanner := bufio.NewScanner(fileLines)

	lines := 0
	for scanner.Scan() {
		lines++
	}

	fmt.Println("File name: ", *filePathPtr)
	fmt.Println("File size: ", file.Size())
	fmt.Println("File lines: ", lines)
}

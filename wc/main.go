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

	fileStats, err := os.Stat(*filePathPtr)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(*filePathPtr) // replace with your file name
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := 0
	words := ""
	for scanner.Scan() {
		lines++
		words += scanner.Text() + "\n"
	}

	wordCount := 0
	for i := 0; i < len(words); i++ {
		wordCount = i
	}

	fmt.Println("File name: ", *filePathPtr)
	fmt.Println("File size: ", fileStats.Size())
	fmt.Println("File lines: ", lines)
	fmt.Println("File words: ", words)
	fmt.Println("File words: ", wordCount)
}

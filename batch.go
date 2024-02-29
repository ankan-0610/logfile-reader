package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	processLogsInBatches()
}

func processLogsInBatches() {
	filePath := "example.log"
	batchSize := 100

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]string, batchSize)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Read lines in batches
		for i := 0; i < batchSize && scanner.Scan(); i++ {
			lines[i] = scanner.Text()
		}
		processBatch(lines)
	}
}

func processBatch(lines []string) {
	// Process batch of lines
	for _, line := range lines {
		fmt.Println(line)
	}
}

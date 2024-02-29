package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func main(){
	concurrent()
}

func processLine(line string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Process each line
	fmt.Println(line)
}

func concurrent() {
	filePath := "example.log"
	numWorkers := 5

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	linesChannel := make(chan string, numWorkers)
	var wg sync.WaitGroup

	// Launch goroutines to concurrently process lines
	for i := 0; i < numWorkers; i++ {
		go func() {
			for line := range linesChannel {
				processLine(line, &wg)
			}
		}()
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		linesChannel <- line
	}

	close(linesChannel)
	wg.Wait()
}
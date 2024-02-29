package main

import(
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	buffered()
}

func buffered(){
	filePath := "example.log"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
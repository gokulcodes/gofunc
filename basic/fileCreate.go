package main

import (
	"os"
)

func main() {
	file, _ := os.Create("./defer.pdf")
	defer file.Close()
	for {
		break
	}
}

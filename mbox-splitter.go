package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fp := flag.String("f", "./test_file", "path to the mbox file you wish to split")
	flag.Parse()
	f, e := os.Open(*fp)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if e := scanner.Err(); e != nil {
		log.Fatal(e)
	}
}

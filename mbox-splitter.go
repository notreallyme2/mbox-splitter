package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file := flag.String("f", "./test_file", "path to the mbox file you wish to split")
	flag.Parse()
	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var b strings.Builder
	var l, currentFirstl, nextFirstl string
	s := bufio.NewScanner(f)
	s.Scan()
	currentFirstl = s.Text()
	addLine(&b, currentFirstl)

	for s.Scan() {
		l = s.Text()
		if strings.HasPrefix(l, "From ") {
			nextFirstl = l
			writeEmail(b.String())
			b.Reset()
			currentFirstl = nextFirstl
			addLine(&b, currentFirstl)
		} else {
			addLine(&b, l)
		}

	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func addLine(b *strings.Builder, l string) {
	b.WriteString(fmt.Sprintf("%s\n", l))
}

func writeEmail(email string) {
	fmt.Println("*********************")
	fmt.Println("WRITING EMAIL TO FILE")
	fmt.Println("*********************")
	fmt.Println(email)
	fmt.Println("***")
	fmt.Println("END")
	fmt.Println("***")
}

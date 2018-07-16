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

	// grab the first line of the mbox file to get us going, add it to the current email
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

// addLine adds the next line to the email.
// For some reason, Builder.WriteString strips new lines (!?), hence this function.
func addLine(b *strings.Builder, l string) {
	b.WriteString(fmt.Sprintf("%s\n", l))
}

// writeEmail writes a single email to a new archive, based on year.
// It checks whether YEAR.mbox exists, if not it creates it in the same directory as the original mbox file.
// It then adds a single email to the YEAR.mbox file
func writeEmail(email string) {
	fmt.Println("*********************")
	fmt.Println("WRITING EMAIL TO FILE")
	fmt.Println("*********************")
	fmt.Println(email)
	fmt.Println("***")
	fmt.Println("END")
	fmt.Println("***")
}

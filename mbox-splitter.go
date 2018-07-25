package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Splitting your mbox file...")
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
// It checks whether YEAR.mbox exists.
// If not it creates it in the directory from which mbox-splitter was run.
// It then adds a single email to the YEAR.mbox file
func writeEmail(email string) {
	// declare file handle
	var f *os.File
	// extract the year
	re, _ := regexp.Compile(`(^From.*)`)
	fl := re.FindString(email)
	if len(fl) == 0 {
		log.Fatal(`Could not identify "From:" line, and therefore unable to extract year.`)
	}
	y := fl[len(fl)-4 : len(fl)]
	file := y + ".mbox"
	// open file, create new file if needed
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		if os.IsNotExist(err) {
			// create file
			fmt.Printf("Creating %s\n", file)
			f, err = os.Create(file)
		} else {
			log.Fatal(err)
		}
	}
	// write the email to the new .mbox file
	_, err = f.Write([]byte(email))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

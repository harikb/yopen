package main

import (
	"bufio"
	"github.com/harikb/yopen"
	"log"
)

func main() {

	rdr, err := yopen.NewReader("/usr/share/man/man1/less.1.gz")
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}

	wtr1, err := yopen.NewWriter("/tmp/x1.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}

	wtr2, err := yopen.NewWriter("/tmp/x1.gz")
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}

	in := bufio.NewScanner(rdr)
	out1 := bufio.NewWriter(wtr1)
	out2 := bufio.NewWriter(wtr2)

	for in.Scan() {

		line := in.Bytes()

		_, err = out1.Write(line)
		if err != nil {
			log.Fatalf("Unable to write to file: %v", err)
		}

		_, err = out2.Write(line)
		if err != nil {
			log.Fatalf("Unable to write to file: %v", err)
		}
	}

	err = rdr.Close()
	if err != nil {
		log.Fatalf("Unable to write to file: %v", err)
	}
	err = out1.Flush()
	if err != nil {
		log.Fatalf("Unable to write to file: %v", err)
	}
	err = out2.Flush()
	if err != nil {
		log.Fatalf("Unable to write to file: %v", err)
	}
	err = wtr1.Close()
	if err != nil {
		log.Fatalf("Unable to write to file: %v", err)
	}
	err = wtr2.Close()
	if err != nil {
		log.Fatalf("Unable to write to file: %v", err)
	}
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/tomsteele/go-nmap"
)

func main() {
	prettyPrint := flag.Bool("p", false, "Pretty-print JSON output")
	outFile := flag.String("o", "", "Write output to `file` instead of standard output")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTION]... [FILE]...\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Convert nmap XML FILE(s) to JSON")
		fmt.Fprintln(flag.CommandLine.Output(), "\nWith no FILE, or when FILE is -, read standard input.\n")
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "\nExamples:")
		fmt.Fprintln(flag.CommandLine.Output(), `  nmap -sn -oX network_data.xml 192.168.2.0/24 && nmap2json network_data.xml
  nmap -sn -oX - 192.168.2.0/24 | nmap2json`)
	}
	flag.Parse()

	w := os.Stdout
	if *outFile != "" {
		f, err := os.Create(*outFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	}
	enc := json.NewEncoder(w)
	if *prettyPrint {
		enc.SetIndent("", "  ")
	}

	readers := make([]io.Reader, 0)

	if len(flag.Args()) == 0 {
		// Read from standard input
		readers = append(readers, os.Stdin)
	}

	for _, s := range flag.Args() {
		if s == "-" {
			// Read from standard input
			readers = append(readers, os.Stdin)
			continue
		}
		f, err := os.Open(s)
		if err != nil {
			log.Println(err)
			continue
		}
		defer f.Close()
		readers = append(readers, f)
	}

	for _, r := range readers {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Println(err)
			continue
		}
		nr, err := nmap.Parse(b)
		if err != nil {
			log.Println(err)
			continue
		}
		if err := enc.Encode(nr); err != nil {
			log.Println(err)
		}
	}
}

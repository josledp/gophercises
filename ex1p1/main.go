package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	problemsFile := flag.String("problems", "problems.csv", "problems file to load")
	pf, err := os.Open(*problemsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer pf.Close()

	pfcsv := csv.NewReader(pf)

	var records []string
	reader := bufio.NewReader(os.Stdin)
	total := 0
	correct := 0
	for records, err = pfcsv.Read(); err == nil; records, err = pfcsv.Read() {
		question := records[0]
		answer := records[1]
		fmt.Println(question)
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n\r")
		if text == answer {
			correct++
		}
		total++
	}
	if err != io.EOF {
		log.Fatal(err)
	}
	fmt.Printf("Total: %d, correct: %d\n", total, correct)
}

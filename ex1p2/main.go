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
	"time"
)

func main() {

	problemsFile := flag.String("problems", "problems.csv", "problems file to load")
	timer := flag.Int64("timer", 30, "max time in seconds")
	flag.Parse()

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
	fmt.Println("please press enter to start the quiz")
	reader.ReadString('\n')

	ticker := time.NewTicker(time.Duration(*timer) * time.Second)
	answerCh := make(chan string)
mainLoop:
	for records, err = pfcsv.Read(); err == nil; records, err = pfcsv.Read() {
		question := records[0]
		answer := records[1]
		fmt.Println(question)
		go func() {
			text, _ := reader.ReadString('\n')
			text = strings.Trim(text, "\n\r")
			answerCh <- text

		}()
		select {
		case <-ticker.C:
			fmt.Println("oooooooohhhhh timed out!")
			break mainLoop
		case text := <-answerCh:
			if text == answer {
				correct++
			}
			total++
		}
	}
	if err != io.EOF && err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total: %d, correct: %d\n", total, correct)
}

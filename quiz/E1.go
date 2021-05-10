package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	countPtr := flag.Int("count", 30, "timer value")
	filePtr := flag.String("file", "problems.csv", "input file")
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		return
	}
	defer file.Close()

	timer := time.NewTimer(time.Duration(*countPtr) * time.Second)
	done := make(chan bool)

	csvr := csv.NewReader(file)

	correct, total := 0, 0
	go func() {
		for {
			record, err := csvr.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return
			}
			fmt.Println(record[0])
			var input string
			fmt.Scanln(&input)
			if input == record[1] {
				correct++
			}
			total++
		}
		<-done
	}()

	select {
	case <-done:
	case <-timer.C:
		fmt.Println("time's up")
	}

	fmt.Printf("You scored %d / %d.\n", correct, total)
}

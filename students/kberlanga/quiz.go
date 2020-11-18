package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	reader2 := bufio.NewReader(os.Stdin)
	var contentFile []byte
	var err error = nil
	// fmt.Print("Enter file name of Quiz: ")
	fileName := flag.String("file", "problems.csv", "file name of quiz")
	flag.Parse()
	contentFile, err = ioutil.ReadFile(strings.TrimSpace(*fileName))
	if err != nil {
		fmt.Println("File not exists")
	}

	// for {
	// 	fileName, _ := reader2.ReadString('\n')
	// 	contentFile, err = ioutil.ReadFile(strings.TrimSpace(fileName))
	// 	if err != nil {
	// 		fmt.Println("File not exists")
	// 		fmt.Print("Enter correct file name of Quiz: ")
	// 		continue
	// 	}
	// 	break
	// }

	totalQuestions := 0
	correctAnswer := 0
	reader := csv.NewReader(strings.NewReader(string(contentFile)))

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		totalQuestions++
		fmt.Println("Question: ", row[0])
		fmt.Print("Enter your answer: ")
		userAnswer, _ := reader2.ReadString('\n')
		if strings.TrimSpace(userAnswer) == row[1] {
			correctAnswer++
		}
	}
	fmt.Println("Total questions of Quiz:", totalQuestions)
	fmt.Println("Total correct answers:", correctAnswer)

	// fmt.Printf("File contents: %s", csv_content)
}

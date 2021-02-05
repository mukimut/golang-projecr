package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
)

type CliStreamerRecord struct {
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

var wg = sync.WaitGroup{}
var ch = make(chan string, 50)

func main() {
	wg.Add(2)
	args := os.Args[1]
	go processArgs(args)
	messages := ""
	file, _ := os.Create("./messages.txt")
	writer := bufio.NewWriter(file)

	go func(ch <-chan string) {
		for message := range ch {
			fmt.Println(message)
			messages += message + "\n"
		}
		wg.Done()
		writer.WriteString(messages)
		writer.Flush()
	}(ch)

	wg.Wait()

}

func processArgs(arg string) {
	var cliStreamers []CliStreamerRecord
	gocsv.UnmarshalString(arg, &cliStreamers)
	fmt.Println(cliStreamers)

	for i := 0; i < cliStreamers[0].RunTimes; i++ {
		message := "CLI Invoke " + strconv.Itoa(i+1) + " " + cliStreamers[0].Message1
		fmt.Println(message)
		ch <- message
		time.Sleep(time.Duration(cliStreamers[0].StreamDelay * int(time.Second)))

		message = "CLI Invoke " + strconv.Itoa(i+1) + " " + cliStreamers[0].Message2
		fmt.Println(message)
		ch <- message
		time.Sleep(time.Duration(cliStreamers[0].StreamDelay * int(time.Second)))
	}

	close(ch)
	wg.Done()
}

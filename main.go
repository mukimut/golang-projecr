package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	message1 := "hello1"
	message2 := "hola2"
	delay := 1
	repeat := 3
	command := "Title,Message 1,Message 2,Stream Delay,Run Times \nCLI Invoker Name," + message1 + "," + message2 + "," + strconv.Itoa(delay) + "," + strconv.Itoa(repeat) + " "

	readCommand(command, repeat, message1, message2)

	time.Sleep(100 * time.Millisecond)

}

func readCommand(args string, repeat int, message1, message2 string) {
	cmd := exec.Command("./golang-project", args)
	cmdReader, err := cmd.StdoutPipe()
	var messages []string

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(cmdReader)

	go func() {
		for scanner.Scan() {
			messages = append(messages, scanner.Text())
		}
		test(messages, message1, message2, repeat)
	}()

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
}

func test(allMessages []string, message1 string, message2 string, repeats int) {
	pass := true
	for i := 1; i <= repeats; i++ {
		currentPass := false

		for _, line := range allMessages {
			if checkMessage(line, message1, i) {
				currentPass = true
				break
			}
		}

		if !currentPass {
			pass = false
			fmt.Println("Faile at iteration " + strconv.Itoa(i) + " of " + message1)
		}
		currentPass = false

		for _, line := range allMessages {
			if checkMessage(line, message2, i) {
				currentPass = true
				break
			}
		}

		if !currentPass {
			pass = false
			fmt.Println("Faile at iteration " + strconv.Itoa(i) + " of " + message1)
		}

		if !pass {
			break
		}

		if pass {
			fmt.Print("PASSED")
		} else {
			fmt.Print("FAILED")
		}

	}

}

func checkMessage(line, testMessage string, iteration int) bool {
	return line == "CLI Invoke "+strconv.Itoa(iteration)+" "+testMessage
}

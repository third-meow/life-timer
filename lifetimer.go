package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func errCheck(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

//just prints to terminal for now
func output(s string) {
	fmt.Print(s)
}

func timekeeper(name string, mins int) {
	for i := 0; i < mins; i++ {
		minTimer := time.NewTimer(time.Second)
		<-minTimer.C
	}
	output(fmt.Sprintf("\nTimer %s has finished\n", name))
	promptForNewTimer()
}

func processUserInput(name *string, mins *int) {
	//read in string from user
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	errCheck(err)

	//split input string into
	input_arr := strings.Fields(input)
	*name = input_arr[0]
	*mins, _ = strconv.Atoi(input_arr[1])
}

func promptForNewTimer() {
	output(">>> ")
	var name string
	var mins int
	processUserInput(&name, &mins)

	go timekeeper(name, mins)
	promptForNewTimer()
}

func main() {
	promptForNewTimer()
}

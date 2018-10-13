package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func errCheck(e error) {
	if e != nil {
		fmt.Print("\n")
		fmt.Print(e)
		fmt.Print("\n")
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

func processUserInput() {
	//setup regexs
	help, err := regexp.Compile("[hH][eE][lL][pP]")
	errCheck(err)
	quit, err := regexp.Compile("[qQeE][uUxX][iI][tT]")
	errCheck(err)
	timerDetails, err := regexp.Compile("([A-z]*)\\s+\\d+")
	errCheck(err)

	//read in string from user
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	errCheck(err)

	if quit.MatchString(input) {
		os.Exit(3)
	} else if help.MatchString(input) {
		displayHelp(true)
	} else if timerDetails.MatchString(input) {
		//split input string into
		input_arr := strings.Fields(input)
		name := input_arr[0]
		mins, err := strconv.Atoi(input_arr[1])
		errCheck(err)
		go timekeeper(name, mins)
	} else {
		output("Command not found, type 'help' for options\n")
	}
}

func displayHelp(full bool) {
	output("Enter timer details in format 'timer-name minutes'\n")
	if full {
		output("Other commands are:\n\texit - exit program\n\tquit /\n\thelp - display this help texti\n")
	}
}

func promptForNewTimer() {
	output(">>> ")
	processUserInput()
	promptForNewTimer()
}

func main() {

	displayHelp(false)
	promptForNewTimer()
}

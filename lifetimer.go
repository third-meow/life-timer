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

//if error exists, print it
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

//main timer function
func timekeeper(name string, mins int) {
	//wait for the duration specifed by mins
	for i := 0; i < mins; i++ {
		minTimer := time.NewTimer(time.Second)
		<-minTimer.C
	}

	//print timer finished message
	output(fmt.Sprintf("\nTimer %s has finished\n", name))

	//prompt user for new tiemer
	promptForCommand()
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

	//classify user input
	if quit.MatchString(input) {
		//exit program with status 3
		os.Exit(3)
	} else if help.MatchString(input) {
		//display (full) help text
		displayHelp(true)
	} else if timerDetails.MatchString(input) {
		//split input string into name and mins
		input_arr := strings.Fields(input)
		name := input_arr[0]
		mins, err := strconv.Atoi(input_arr[1])
		errCheck(err)

		//start new timer
		go timekeeper(name, mins)
	} else {
		//if user input does not match any command:
		output("Command not found, type 'help' for options\n")
	}
}

//print help text
func displayHelp(full bool) {
	output("Enter timer details in format 'timer-name minutes'\n")
	if full {
		output("Other commands are:\n\texit - exit program\n\tquit /\n\thelp - display this help texti\n")
	}
}

//give prompt to user for command
func promptForCommand() {
	//print prompt
	output(">>> ")

	//process user's input
	processUserInput()

	//prompt again
	promptForCommand()
}

func main() {
	//start by displaying (minimal) help text
	displayHelp(false)

	//prompt for timer
	promptForCommand()
}

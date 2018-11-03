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

type timekeeper struct {
	name               string
	duration, minCount int
}

type commandRegexSet struct {
	help, quit, timerDetails *regexp.Regexp
}

var commands commandRegexSet


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

func promptAndProcessInput() {
	//print prompt
	output(">>> ")

	//read text
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	errCheck(err)

	commandType := commands.type(input)

	//classify user input
	if commandType == "quit" {
		//exit program with status 3
		os.Exit(3)
	} else if commandType == "help" {
		//display (full) help text
		displayHelp(true)
	} else if commandType == "timerDetails" {
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


func setupRegexs() {
	var err error
	commands.help, err = regexp.Compile("[hH][eE][lL][pP]")
	errCheck(err)
	commands.quit, err = regexp.Compile("[qQeE][uUxX][iI][tT]")
	errCheck(err)
	commands.timerDetails, err = regexp.Compile("([A-z]*)\\s+\\d+")
	errCheck(err)
}



func main() {
	//read in string from user
	//start by displaying (minimal) help text
	//displayHelp(false)
	//while true {
	//	promptAndProcessInput()
	//}
}

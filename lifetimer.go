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

var timekeepers []timekeeper

func (tk *timekeeper) run() {
	//wait for the duration
	for tk.minCount = 0; tk.minCount < tk.duration; tk.minCount++ {
		minTimer := time.NewTimer(time.Second)
		<-minTimer.C
	}

	//print timer finished message
	output(fmt.Sprintf("\nTimer %s has finished\n", tk.name))
}

func (tk *timekeeper) status() [2]int {
	status := [2]int{tk.minCount, tk.duration}
	return status
}
type commandRegexSet struct {
	help, quit, timerDetails *regexp.Regexp
}

var commands commandRegexSet

func (crs *commandRegexSet) commandType(in string) string {
	if crs.quit.MatchString(in) {
		return "quit"
	} else if crs.help.MatchString(in) {
		return "help"
	} else if crs.timerDetails.MatchString(in) {
		return "timerDetails"
	} else {
		return ""
	}
}

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

func newTimekeeper(details string) {
	detailsArr := strings.Fields(details)
	name := detailsArr[0]
	mins, err := strconv.Atoi(detailsArr[1])
	errCheck(err)

	tempTimekeeper := timekeeper{name: name, duration: mins}
	timekeepers = append(timekeepers, tempTimekeeper)
	timekeepers[len(timekeepers)-1].run()
}

//prompt user for input, then process any input given
func promptAndProcessInput() {
	//print prompt
	output(">>> ")

	//read text
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	errCheck(err)

	commandType := commands.commandType(input)

	//classify user input
	if commandType == "quit" {
		//exit program with status 3
		os.Exit(3)

	} else if commandType == "help" {
		//display (full) help text
		displayHelp(true)

	} else if commandType == "timerDetails" {
		//start timer
		newTimekeeper(input)

	} else {
		//if user input does not match any command:
		output("Command not found, type 'help' for options\n")
	}
}

//print a status bar
func printStatusBar(status [2]int) {
	bar := "|"
	for i := 0; i < status[0]; i++ {
		bar = bar + "#"
	}
	remaining := status[1] - status[0]
	for i := 0; i < remaining; i++ {
		bar = bar + "~"
	}
	bar = bar + "|\n"
	output(bar)
}

//display status bars
func displayTimerStates() {
	for _, tk := range timekeepers {
		status := tk.status()
		printStatusBar(status)
	}
}

//remove any inactive/finsihed timers
func removeInactiveTimers() {
	for i, tk := range timekeepers {
		if tk.done {
			timekeepers[i] = timekeepers[len(timekeepers)-1]
			timekeepers = timekeepers[:len(timekeepers)-1]
		}
	}
}

func main() {
	setupRegexs()
	displayHelp(false)
	for {
		displayTimerStates()
		removeInactiveTimers()
	}
}

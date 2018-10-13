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

func timekeeper(name string, mins int) {
	for i := 0; i < mins; i++ {
		minTimer := time.NewTimer(time.Second)
		<-minTimer.C
	}
	fmt.Printf("Timer %s has finished\n", name)
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

func main() {
	fmt.Println("Enter timer details \t\t\t\t Format: 'name length'")
	var name string
	var mins int
	processUserInput(&name, &mins)

	done_reports := make(chan string)

	go timekeeper(name, mins, done_reports)
	dt := <-done_reports
	fmt.Println(dt)

}

package main

import (
	"fmt"
	"time"
	ui "uilive"
)

func main() {
	writer := ui.New()
	writer.Start()
	for i := 0; i < 14; i++ {
		fmt.Fprintf(writer, "\n%d", i)
		fmt.Fprintf(writer, "\n>>>")
		if i%4 == 3 {
			fmt.Printf("\ntimer %d is done\n\n", i)
		}
		time.Sleep(time.Second)
	}
	writer.Stop()
}

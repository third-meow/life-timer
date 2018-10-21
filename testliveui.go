package main

import (
	"fmt"
	"time"
	ui "uilive"
)

func main() {
	instrings := make([]string, 0)

	writer := ui.New()
	writer.Start()
	for i := 0; i < 14; i++ {
		fmt.Fprintf(writer, "\n%d ", i)
		fmt.Fprintf(writer, ">>>")
		go func() {
			var s string
			_, _ = fmt.Scanf("%s\n", &s)
			instrings = append(instrings, s)
		}()
		if i%4 == 3 {
			fmt.Fprintf(writer.Bypass(), "\ntimer %d is done", i)
		}
		time.Sleep(time.Second)
	}
	writer.Stop()

	fmt.Println("")
	for _, e := range instrings {
		fmt.Println(e)
	}
}

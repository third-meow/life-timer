package main

import (
	"fmt"
	"time"
	ui "uilive"
)

func dispAll(w *ui.Writer, i int) {
	fmt.Fprintf(w, "\n%d ", i)
	fmt.Fprintf(w, ">>>")
	if i%4 == 3 {
		fmt.Fprintf(w.Bypass(), "\ntimer %d is done", i)
	}
}

func takeInput(w *ui.Writer, i int, out *[]string) {
	var s string
	fmt.Scanf("%s\n", &s)
	*out = append(*out, s)
	dispAll(w, i)
}

func main() {
	instrings := make([]string, 0)

	writer := ui.New()
	writer.Start()
	for i := 0; i < 14; i++ {
		dispAll(writer, i)
		go takeInput(writer, i, &instrings)
		time.Sleep(time.Second * 8)
	}
	writer.Stop()

	fmt.Println("")
	for _, e := range instrings {
		fmt.Println(e)
	}
	for i := 0; i < 3; i++ {
		var nothing string
		fmt.Scanln(&nothing)
	}
}

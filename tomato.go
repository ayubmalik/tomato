package tomato

import (
	"fmt"
	"io"
	"time"
)

// New starts timer and returns a channel which will be closed after duration d.
// After d the screen fill flash.
func New(w io.Writer, d time.Duration) chan int {
	quit := make(chan int)
	ticker := time.NewTicker(1 * time.Second)
	remaining := int(d / time.Second)
	write(w, remaining)
	go func() {
		time.AfterFunc(d, func() {
			ticker.Stop()
			bell(w)
			flash(w)
			close(quit)
		})

		for {
			select {
			case <-quit:
				return
			case <-ticker.C:
				remaining--
				write(w, remaining)
			}
		}
	}()
	return quit
}

// Reset terminal ANSI sequences
func Reset(w io.Writer) {
	io.WriteString(w, "\033[?5l")
	io.WriteString(w, "\033[0m")
}

func flash(w io.Writer) {
	for i := 0; i < 10; i++ {
		io.WriteString(w, "\033]0;!!! TOMATO FINISHED !!!\033\\")
		io.WriteString(w, "\033[?5h")
		time.Sleep(time.Millisecond * 500)
		io.WriteString(w, "\033]0;\033\\")
		io.WriteString(w, "\033[?5l")
		time.Sleep(time.Millisecond * 500)
	}
}

func bell(w io.Writer) {
	for i := 0; i < 10; i++ {
		io.WriteString(w, "\a")
		time.Sleep(time.Millisecond * 500)
	}
}

func write(w io.Writer, remaining int) {
	d := time.Duration(remaining) * time.Second
	s := fmt.Sprintf("\033[1K\r\033[38;5;226;7;1m%s\033[0m", d)
	io.WriteString(w, s)
}

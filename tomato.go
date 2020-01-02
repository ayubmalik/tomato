package tomato

import (
	"fmt"
	"io"
	"time"
)

// New starts timer and returns a channel which will be closed after remaining secs
func New(w io.Writer, d time.Duration) chan int {
	quit := make(chan int)
	ticker := time.NewTicker(1 * time.Second)
	remaining := int(d / time.Second)
	write(w, remaining)
	go func() {
		time.AfterFunc(d, func() {
			ticker.Stop()
			io.WriteString(w, "\033[?25h\033[0m")
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

func write(w io.Writer, remaining int) {
	d := time.Duration(remaining) * time.Second
	s := fmt.Sprintf("\033[?25l\033[1K\r\033[38;5;226;7;1m%s\033[0m", d)
	io.WriteString(w, s)
}

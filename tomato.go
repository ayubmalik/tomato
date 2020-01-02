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
	io.WriteString(w, fmt.Sprintf("%d", remaining))
}

package tomato

import (
	"fmt"
	"io"
	"time"
)

// New starts timer and returns a channel which will be closed after n secs
func New(w io.Writer, n int) chan int {
	quit := make(chan int)
	ticker := time.NewTicker(1 * time.Second)
	d := time.Duration(n) * time.Second
	go func() {
		time.AfterFunc(d, func() {
			defer ticker.Stop()
			close(quit)
		})

		for {
			select {
			case <-quit:
				return
			case <-ticker.C:
				n--
				io.WriteString(w, fmt.Sprintf("%d", n))
			}
		}
	}()
	return quit
}

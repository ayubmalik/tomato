package main

import (
	"os"

	"github.com/ayubmalik/tomato"
)

func main() {
	tom := tomato.New(os.Stdout, 6)
	<-tom
}

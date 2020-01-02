package main

import (
	"os"
	"time"

	"github.com/ayubmalik/tomato"
)

func main() {
	d := time.Second * 5
	tom := tomato.New(os.Stdout, d)
	<-tom
}

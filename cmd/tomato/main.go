package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ayubmalik/tomato"
)

func main() {
	tFlag := flag.String("time", "30m", "time tomato runs for, e.g 30m 1h 45s")
	flag.Parse()
	d, err := time.ParseDuration(*tFlag)
	if err != nil {
		fmt.Println((err))
		os.Exit(1)
	}

	tom := tomato.New(os.Stdout, d)
	<-tom
}

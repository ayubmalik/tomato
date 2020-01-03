package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ayubmalik/tomato"
)

func main() {
	tFlag := flag.String("time", "30m", "the amount of time tomato runs for, e.g 30m 1h 45s")
	flag.Parse()
	d, err := time.ParseDuration(*tFlag)
	if err != nil {
		fmt.Println((err))
		os.Exit(1)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		reset()
	}()

	tom := tomato.New(os.Stdout, d)
	<-tom
	reset()
}

func reset() {
	tomato.Reset(os.Stdout)
	os.Exit(0)
}

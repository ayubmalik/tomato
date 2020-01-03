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
	tFlag := flag.String("time", "30m", "the time tomato runs for e.g 30m 1h 45s")
	flag.Parse()
	d, err := time.ParseDuration(*tFlag)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		reset(1)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		reset(1)
	}()

	tom := tomato.New(os.Stdout, d)
	<-tom
	reset(0)
}

func reset(n int) {
	tomato.Reset(os.Stdout)
	os.Exit(n)
}

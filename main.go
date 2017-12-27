package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"
)

const (
	defaultPubSubTopic = "tweets"
)

var (
	// ctx
	appContext context.Context

	// google
	projectID         string
	sourceID          string
	frequencyInMilSec int
)

func main() {

	// FLAGS
	flag.StringVar(&projectID, "project", os.Getenv("GCLOUD_PROJECT"), "GCP Project ID")
	flag.IntVar(&frequencyInMilSec, "frequency", 1000, "Duration in milliseconds")
	flag.Parse()

	if projectID == "" {
		log.Fatalf("Missing required app configs: project:%v", projectID)
	}
	// END FLAGAS

	// HOST
	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error while parsing hostname: %v", err)
	}
	sourceID = name
	// END HOST

	// FREQUENCY
	frequency := time.Duration(frequencyInMilSec) * time.Millisecond
	// END FREQUENCY

	// CONTEXT
	ctx, cancel := context.WithCancel(context.Background())
	appContext = ctx
	messages := make(chan int64)
	// END CONTEXT

	go func() {
		// Wait for SIGINT and SIGTERM (HIT CTRL-C)
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		log.Println(<-ch)
		cancel()
		os.Exit(0)
	}()

	// initialize publisher
	initPublisher()

	// start provider
	go provide(frequency, messages)

	// LOOP
	for {
		select {
		case <-appContext.Done():
			break
		case m := <-messages:
			publish(m)
		}
	}
	// END LOOP

}

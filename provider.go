package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	mockedMetricName = "MockedMetric"
	minMetricVal     = 0
	maxMetricVal     = 100
)

func provide(d time.Duration, ch chan<- int64) {

	log.Printf("Frequency: %v", d)
	ticker := time.NewTicker(d)
	for t := range ticker.C {
		m := getMetric()
		log.Printf("Metric %v - %v", t, m)
		ch <- m
	}

}

func getMetric() int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return int64(minMetricVal + rand.Intn(maxMetricVal-minMetricVal))
}

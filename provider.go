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

func provide(ch chan<- int64) {

	//TODO: Parametirize
	frequency := 1000 * time.Millisecond
	ticker := time.NewTicker(frequency)
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

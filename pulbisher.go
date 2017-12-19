package main

import (
	"log"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3"
	googlepb "github.com/golang/protobuf/ptypes/timestamp"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
	monitoredrespb "google.golang.org/genproto/googleapis/api/monitoredres"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

const (
	customMetricType = "custom.googleapis.com/partner/mocked"
)

var (
	monitorClient *monitoring.MetricClient
)

func initPublisher() {

	client, err := monitoring.NewMetricClient(appContext)
	if err != nil {
		log.Fatalf("Failed to create monitor client: %v", err)
	}

	monitorClient = client
}

func publish(m int64) {

	dataPoint := &monitoringpb.Point{
		Interval: &monitoringpb.TimeInterval{
			EndTime: &googlepb.Timestamp{Seconds: time.Now().Unix()},
		},
		Value: &monitoringpb.TypedValue{
			Value: &monitoringpb.TypedValue_Int64Value{Int64Value: m},
		},
	}

	tsRequest := &monitoringpb.CreateTimeSeriesRequest{
		Name: monitoring.MetricProjectPath(projectID),
		TimeSeries: []*monitoringpb.TimeSeries{
			{
				Metric: &metricpb.Metric{
					Type:   customMetricType,
					Labels: map[string]string{"instance_id": sourceID},
				},
				Resource: &monitoredrespb.MonitoredResource{
					Type:   "global",
					Labels: map[string]string{"project_id": projectID},
				},
				Points: []*monitoringpb.Point{dataPoint},
			},
		},
	}

	if err := monitorClient.CreateTimeSeries(appContext, tsRequest); err != nil {
		log.Printf("Failed to write time series data: %v", err)
	}

}

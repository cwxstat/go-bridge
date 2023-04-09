package metric

import (
	"fmt"
	"testing"
	"time"
)

func populateMetric(m MetricPopulator) {
	m.Name("Example Metric")
	m.Start(time.Now().UTC().Add(-1 * time.Hour))
	m.End(time.Now().UTC())

	// Always UTC
	m.AddDataPoint(time.Now().Add(-30*time.Minute), &Value{Float64: 42.078})
	m.AddDataPoint(time.Now().Add(-15*time.Minute), &Value{Int32: 25})
	m.AddDataPoint(time.Now().Add(-10*time.Minute), &Value{String: "offline"})
	m.AddDataPoint(time.Now().Add(-9*time.Minute), &Value{String: "restarting"})
	m.AddDataPoint(time.Now().Add(-8*time.Minute), &Value{String: "online"})
	m.AddDataPoint(time.Now().Add(-7*time.Minute), &Value{Float64: 0.0})
	m.AddDataPoint(time.Now().Add(-6*time.Minute), &Value{Float64: 71.067888888})

	m.AddMetadata("source", "sensor-001")
	m.AddMetadata("location", "building-1")
	m.AddLabel("account", "001")
	m.AddLabel("region", "us-east-1")
}

func TestMetric_AddDataPoint(t *testing.T) {
	m := NewMetric()
	populateMetric(m)

	// Create a MetricReader for the metric
	metricReader := NewMetricReader(m)

	// Use the MetricReader interface to access the data
	fmt.Println("Metric Name:", metricReader.GetName())
	fmt.Println("Start Time:", metricReader.GetStart())
	fmt.Println("End Time:", metricReader.GetEnd())

	fmt.Println("Data Points:")
	for _, data := range metricReader.GetData() {
		fmt.Printf("  %v: %v\n", data.DateTime, data.Value.GetValidValue())
	}

	fmt.Println("Metadata:")

	for key, value := range metricReader.GetMetadata() {
		fmt.Printf("  %v: %v\n", key, value)
	}
	for key, value := range metricReader.GetLabels() {
		fmt.Printf("  %v: %v\n", key, value)
	}
}

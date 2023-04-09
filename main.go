package main

import (
	"fmt"
	"gobridge/metric"
	"time"
)

type Data struct {
	Data []data
}

type data struct {
	Time time.Time
	Data string
}

func (d *Data) Add(time time.Time, s string) {
	d.Data = append(d.Data, data{Time: time, Data: s})
}

func (d *Data) Get() []data {
	return d.Data
}

// CollectorInterface defines the interface for the collector implementation
type CollectorInterface interface {
	Collect() *Data
}

// ResolverInterface defines the interface for the resolver abstraction
type ResolverInterface interface {
	Resolve() *metric.Metric
	SetCollector(CollectorInterface)
}

// DispatcherInterface defines the interface for the dispatcher abstraction
type DispatcherInterface interface {
	Dispatch()
	SetResolver(ResolverInterface)
}

// ConcreteCollector is a concrete implementation of the CollectorInterface
type ConcreteCollector struct{}

// Collect method for ConcreteCollector
func (c *ConcreteCollector) Collect() *Data {
	fmt.Println("Collector: Collecting data")
	data := &Data{}
	data.Add(time.Now(), "Collecting data...")
	data.Add(time.Now(), "2.3")
	data.Add(time.Now(), "4.3")
	data.Add(time.Now(), "....done collecting data")
	return data
}

// ConcreteResolver is a concrete implementation of the ResolverInterface
type ConcreteResolver struct {
	collector CollectorInterface
	metric    *metric.Metric
}

// Resolve method for ConcreteResolver
//
//	we could do analysis here, but for now we just return the count of data points
func (r *ConcreteResolver) Resolve() *metric.Metric {
	r.metric = metric.NewMetric()

	populateMetric(r.metric, r.collector.Collect())
	return r.metric
}

func populateMetric(m metric.MetricPopulator, data *Data) {
	m.Name("Resolver Metric")
	m.Start(time.Now().UTC().Add(-1 * time.Hour))
	m.End(time.Now().UTC())

	// Simulate error
	m.AddDataPoint(time.Now().Add(-10*time.Minute), &metric.Value{String: "offline"})
	m.AddDataPoint(time.Now().Add(-9*time.Minute), &metric.Value{String: "restarting"})
	m.AddDataPoint(time.Now().Add(-8*time.Minute), &metric.Value{String: "online"})

	for _, v := range data.Get() {
		m.AddDataPoint(v.Time, &metric.Value{String: v.Data})
	}

	m.AddMetadata("source", "sensor-001")
	m.AddMetadata("location", "building-1")
	m.AddLabel("account", "001")
	m.AddLabel("region", "us-east-1")
}

// SetCollector sets the collector for ConcreteResolver
func (r *ConcreteResolver) SetCollector(c CollectorInterface) {
	r.collector = c
}

// ConcreteDispatcher is a concrete implementation of the DispatcherInterface
type ConcreteDispatcher struct {
	resolver ResolverInterface
	reader   metric.MetricReader
}

// Dispatch method for ConcreteDispatcher
func (d *ConcreteDispatcher) Dispatch() {
	d.reader = metric.NewMetricReader(d.resolver.Resolve())
	fmt.Println("Dispatching data")

	for _, data := range d.reader.GetData() {
		fmt.Printf("  %v: %v\n", data.DateTime, data.Value.GetValidValue())
	}

}

// SetResolver sets the resolver for ConcreteDispatcher
func (d *ConcreteDispatcher) SetResolver(r ResolverInterface) {
	d.resolver = r
}

func main() {
	collector := &ConcreteCollector{}
	resolver := &ConcreteResolver{}
	resolver.SetCollector(collector)
	dispatcher := &ConcreteDispatcher{}
	dispatcher.SetResolver(resolver)

	dispatcher.Dispatch()
}

package metric

import "time"

type MetricReader interface {
	GetName() string
	GetStart() time.Time
	GetEnd() time.Time
	GetData() []MetricData
	GetMetadata() map[string]interface{}

	GetLabels() map[string]string
}

// MetricReader implementation
type MetricReaderImpl struct {
	metric *Metric
}

func NewMetricReader(metric *Metric) MetricReader {
	return &MetricReaderImpl{metric: metric}
}

func (mr *MetricReaderImpl) GetName() string {
	return mr.metric.GetName()
}

func (mr *MetricReaderImpl) GetStart() time.Time {
	return mr.metric.start
}

func (mr *MetricReaderImpl) GetEnd() time.Time {
	return mr.metric.end
}

func (mr *MetricReaderImpl) GetData() []MetricData {
	return *mr.metric.Data
}

func (mr *MetricReaderImpl) GetMetadata() map[string]interface{} {
	return mr.metric.Meta.Meta

}

func (mr *MetricReaderImpl) GetLabels() map[string]string {
	return mr.metric.Meta.Labels
}

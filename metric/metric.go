package metric

import "time"

type MetricPopulator interface {
	Name(name string)
	Start(time.Time)
	End(time.Time)
	AddDataPoint(dateTime time.Time, value *Value)
	AddMetadata(key string, value interface{})

	AddLabel(key string, value string)
}

type Metric struct {
	name  string
	start time.Time
	end   time.Time
	Data  *[]MetricData
	Meta  *MetricMeta
}
type MetricData struct {
	DateTime time.Time `required:"true"`
	Value    *Value
}

type MetricMeta struct {
	Labels map[string]string
	Meta   map[string]interface{}
}

type Value struct {
	Float64 float64
	Bool    bool
	String  string
	Int32   int32
	Int64   int64
}

func (v *Value) GetValidValue() interface{} {
	if v.Float64 != 0 {
		return v.Float64
	}
	if v.Bool != false {
		return v.Bool
	}
	if v.String != "" {
		return v.String
	}
	if v.Int32 != 0 {
		return v.Int32
	}
	if v.Int64 != 0 {
		return v.Int64
	}
	return v.Int32
}

func NewMetric() *Metric {
	return &Metric{
		Data: &[]MetricData{},
		Meta: &MetricMeta{},
	}
}

func (m *Metric) addData(data MetricData) {
	*m.Data = append(*m.Data, data)
}

func (m *Metric) addMeta(key string, value interface{}) {
	if m.Meta.Meta == nil {
		m.Meta.Meta = make(map[string]interface{})
	}
	m.Meta.Meta[key] = value
}

func (m *Metric) addLabel(key string, value string) {
	if m.Meta.Labels == nil {
		m.Meta.Labels = make(map[string]string)
	}
	m.Meta.Labels[key] = value
}

func (m *Metric) Start(t time.Time) {
	m.start = t
}
func (m *Metric) End(t time.Time) {
	m.end = t
}

func (m *Metric) Name(n string) {
	m.name = n
}

func (m *Metric) GetName() string {
	return m.name
}

func (m *Metric) GetStart() string {
	return m.start.Format(time.RFC3339)
}

func (m *Metric) GetEnd() string {
	return m.start.Format(time.RFC3339)
}

func (m *Metric) AddDataPoint(dateTime time.Time, value *Value) {
	m.addData(MetricData{DateTime: dateTime, Value: value})
}

func (m *Metric) AddMetadata(key string, value interface{}) {
	m.addMeta(key, value)
}

func (m *Metric) AddLabel(key string, value string) {
	m.addLabel(key, value)
}

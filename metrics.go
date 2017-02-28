package metser

//go:generate msgp -o msgp_gen.go -io=false -tests=false

// Type is a metric type
type Type int8

// List of supported metric types
const (
	UnknownType Type = iota
	CounterType
	BatchTimerType
	GaugeType
)

// ID is the metric id
type ID []byte

// Counter is a counter containing the counter ID and the counter value
type Counter struct {
	ID    ID
	Value int64
}

// BatchTimer is a timer containing the timer ID and a list of timer values
type BatchTimer struct {
	ID     ID
	Values []float64
}

// Gauge is a gauge containing the gauge ID and the value at certain time
type Gauge struct {
	ID    ID
	Value float64
}

// MetricUnion is a union of different types of metrics, only one of which is valid
// at any given time. The actual type of the metric depends on the type field,
// which determines which value field is valid.
type MetricUnion struct {
	Type          Type
	ID            ID
	CounterVal    int64
	BatchTimerVal []float64
	GaugeVal      float64
}

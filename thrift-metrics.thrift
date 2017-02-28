struct ThriftCounter {
  1: binary ID
  2: i64 Value
}

struct ThriftBatchTimer {
  1: binary ID
  2: list<double> Values
}

struct ThriftGauge {
  1: binary ID
  2: double Value
}

struct ThriftMetricUnion {
  1: byte Type
  2: binary ID
  3: i64 CounterVal
  4: list<double> BatchTimerVal
  5: double GaugeVal
}
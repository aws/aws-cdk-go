package awscloudwatch


// Unit for metric.
//
// Example:
//   var logGroup logGroup
//
//   mf := logs.NewMetricFilter(this, jsii.String("MetricFilter"), &MetricFilterProps{
//   	LogGroup: LogGroup,
//   	MetricNamespace: jsii.String("MyApp"),
//   	MetricName: jsii.String("Latency"),
//   	FilterPattern: logs.FilterPattern_Exists(jsii.String("$.latency")),
//   	MetricValue: jsii.String("$.latency"),
//   	Dimensions: map[string]*string{
//   		"ErrorCode": jsii.String("$.errorCode"),
//   	},
//   	Unit: cloudwatch.Unit_MILLISECONDS,
//   })
//
//   //expose a metric from the metric filter
//   metric := mf.Metric()
//
//   //you can use the metric to create a new alarm
//   //you can use the metric to create a new alarm
//   cloudwatch.NewAlarm(this, jsii.String("alarm from metric filter"), &AlarmProps{
//   	Metric: Metric,
//   	Threshold: jsii.Number(100),
//   	EvaluationPeriods: jsii.Number(2),
//   })
//
type Unit string

const (
	// Seconds.
	Unit_SECONDS Unit = "SECONDS"
	// Microseconds.
	Unit_MICROSECONDS Unit = "MICROSECONDS"
	// Milliseconds.
	Unit_MILLISECONDS Unit = "MILLISECONDS"
	// Bytes.
	Unit_BYTES Unit = "BYTES"
	// Kilobytes.
	Unit_KILOBYTES Unit = "KILOBYTES"
	// Megabytes.
	Unit_MEGABYTES Unit = "MEGABYTES"
	// Gigabytes.
	Unit_GIGABYTES Unit = "GIGABYTES"
	// Terabytes.
	Unit_TERABYTES Unit = "TERABYTES"
	// Bits.
	Unit_BITS Unit = "BITS"
	// Kilobits.
	Unit_KILOBITS Unit = "KILOBITS"
	// Megabits.
	Unit_MEGABITS Unit = "MEGABITS"
	// Gigabits.
	Unit_GIGABITS Unit = "GIGABITS"
	// Terabits.
	Unit_TERABITS Unit = "TERABITS"
	// Percent.
	Unit_PERCENT Unit = "PERCENT"
	// Count.
	Unit_COUNT Unit = "COUNT"
	// Bytes/second (B/s).
	Unit_BYTES_PER_SECOND Unit = "BYTES_PER_SECOND"
	// Kilobytes/second (kB/s).
	Unit_KILOBYTES_PER_SECOND Unit = "KILOBYTES_PER_SECOND"
	// Megabytes/second (MB/s).
	Unit_MEGABYTES_PER_SECOND Unit = "MEGABYTES_PER_SECOND"
	// Gigabytes/second (GB/s).
	Unit_GIGABYTES_PER_SECOND Unit = "GIGABYTES_PER_SECOND"
	// Terabytes/second (TB/s).
	Unit_TERABYTES_PER_SECOND Unit = "TERABYTES_PER_SECOND"
	// Bits/second (b/s).
	Unit_BITS_PER_SECOND Unit = "BITS_PER_SECOND"
	// Kilobits/second (kb/s).
	Unit_KILOBITS_PER_SECOND Unit = "KILOBITS_PER_SECOND"
	// Megabits/second (Mb/s).
	Unit_MEGABITS_PER_SECOND Unit = "MEGABITS_PER_SECOND"
	// Gigabits/second (Gb/s).
	Unit_GIGABITS_PER_SECOND Unit = "GIGABITS_PER_SECOND"
	// Terabits/second (Tb/s).
	Unit_TERABITS_PER_SECOND Unit = "TERABITS_PER_SECOND"
	// Count/second.
	Unit_COUNT_PER_SECOND Unit = "COUNT_PER_SECOND"
	// None.
	Unit_NONE Unit = "NONE"
)


package awscdkkinesisanalyticsflinkalpha


// Available log levels for Flink applications.
//
// Example:
//   var bucket Bucket
//
//   flinkApp := flink.NewApplication(this, jsii.String("Application"), &ApplicationProps{
//   	Code: flink.ApplicationCode_FromBucket(bucket, jsii.String("my-app.jar")),
//   	Runtime: flink.Runtime_FLINK_1_20(),
//   	CheckpointingEnabled: jsii.Boolean(true),
//   	 // default is true
//   	CheckpointInterval: awscdk.Duration_Seconds(jsii.Number(30)),
//   	 // default is 1 minute
//   	MinPauseBetweenCheckpoints: awscdk.Duration_*Seconds(jsii.Number(10)),
//   	 // default is 5 seconds
//   	LogLevel: flink.LogLevel_ERROR,
//   	 // default is INFO
//   	MetricsLevel: flink.MetricsLevel_PARALLELISM,
//   	 // default is APPLICATION
//   	AutoScalingEnabled: jsii.Boolean(false),
//   	 // default is true
//   	Parallelism: jsii.Number(32),
//   	 // default is 1
//   	ParallelismPerKpu: jsii.Number(2),
//   	 // default is 1
//   	SnapshotsEnabled: jsii.Boolean(false),
//   	 // default is true
//   	LogGroup: logs.NewLogGroup(this, jsii.String("LogGroup")),
//   })
//
// Experimental.
type LogLevel string

const (
	// Debug level logging.
	// Experimental.
	LogLevel_DEBUG LogLevel = "DEBUG"
	// Info level logging.
	// Experimental.
	LogLevel_INFO LogLevel = "INFO"
	// Warn level logging.
	// Experimental.
	LogLevel_WARN LogLevel = "WARN"
	// Error level logging.
	// Experimental.
	LogLevel_ERROR LogLevel = "ERROR"
)


package awskinesisanalyticsflink


// Available log levels for Flink applications.
//
// Example:
//   var bucket bucket
//
//   flinkApp := flink.NewApplication(this, jsii.String("Application"), &applicationProps{
//   	code: flink.applicationCode.fromBucket(bucket, jsii.String("my-app.jar")),
//   	runtime: flink.runtime_FLINK_1_13(),
//   	checkpointingEnabled: jsii.Boolean(true),
//   	 // default is true
//   	checkpointInterval: awscdk.Duration.seconds(jsii.Number(30)),
//   	 // default is 1 minute
//   	minPauseBetweenCheckpoints: awscdk.Duration.seconds(jsii.Number(10)),
//   	 // default is 5 seconds
//   	logLevel: flink.logLevel_ERROR,
//   	 // default is INFO
//   	metricsLevel: flink.metricsLevel_PARALLELISM,
//   	 // default is APPLICATION
//   	autoScalingEnabled: jsii.Boolean(false),
//   	 // default is true
//   	parallelism: jsii.Number(32),
//   	 // default is 1
//   	parallelismPerKpu: jsii.Number(2),
//   	 // default is 1
//   	snapshotsEnabled: jsii.Boolean(false),
//   	 // default is true
//   	logGroup: logs.NewLogGroup(this, jsii.String("LogGroup")),
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


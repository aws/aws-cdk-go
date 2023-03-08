// A CDK Construct Library for Kinesis Analytics Flink applications
package awscdkkinesisanalyticsflinkalpha


// Granularity of metrics sent to CloudWatch.
//
// Example:
//   var bucket bucket
//
//   flinkApp := flink.NewApplication(this, jsii.String("Application"), &ApplicationProps{
//   	Code: flink.ApplicationCode_FromBucket(bucket, jsii.String("my-app.jar")),
//   	Runtime: flink.Runtime_FLINK_1_13(),
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
type MetricsLevel string

const (
	// Application sends the least metrics to CloudWatch.
	// Experimental.
	MetricsLevel_APPLICATION MetricsLevel = "APPLICATION"
	// Task includes task-level metrics sent to CloudWatch.
	// Experimental.
	MetricsLevel_TASK MetricsLevel = "TASK"
	// Operator includes task-level and operator-level metrics sent to CloudWatch.
	// Experimental.
	MetricsLevel_OPERATOR MetricsLevel = "OPERATOR"
	// Send all metrics including metrics per task thread.
	// Experimental.
	MetricsLevel_PARALLELISM MetricsLevel = "PARALLELISM"
)


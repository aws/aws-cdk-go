package awskinesisanalyticsflink


// Interface for building AWS::KinesisAnalyticsV2::Application PropertyGroup configuration.
//
// Example:
//   var bucket bucket
//
//   flinkApp := flink.NewApplication(this, jsii.String("Application"), &ApplicationProps{
//   	PropertyGroups: &PropertyGroups{
//   		FlinkApplicationProperties: map[string]*string{
//   			"inputStreamName": jsii.String("my-input-kinesis-stream"),
//   			"outputStreamName": jsii.String("my-output-kinesis-stream"),
//   		},
//   	},
//   	// ...
//   	Runtime: flink.Runtime_FLINK_1_13(),
//   	Code: flink.ApplicationCode_FromBucket(bucket, jsii.String("my-app.jar")),
//   })
//
// Experimental.
type PropertyGroups struct {
}


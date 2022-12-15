package awskinesisanalyticsflink


// Interface for building AWS::KinesisAnalyticsV2::Application PropertyGroup configuration.
//
// Example:
//   var bucket bucket
//
//   flinkApp := flink.NewApplication(this, jsii.String("Application"), &applicationProps{
//   	propertyGroups: &propertyGroups{
//   		flinkApplicationProperties: map[string]*string{
//   			"inputStreamName": jsii.String("my-input-kinesis-stream"),
//   			"outputStreamName": jsii.String("my-output-kinesis-stream"),
//   		},
//   	},
//   	// ...
//   	runtime: flink.runtime_FLINK_1_13(),
//   	code: flink.applicationCode.fromBucket(bucket, jsii.String("my-app.jar")),
//   })
//
// Experimental.
type PropertyGroups struct {
}


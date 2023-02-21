package awscloudfront


// Contains information about the Amazon Kinesis data stream where you are sending real-time log data in a real-time log configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endPointProperty := &endPointProperty{
//   	kinesisStreamConfig: &kinesisStreamConfigProperty{
//   		roleArn: jsii.String("roleArn"),
//   		streamArn: jsii.String("streamArn"),
//   	},
//   	streamType: jsii.String("streamType"),
//   }
//
type CfnRealtimeLogConfig_EndPointProperty struct {
	// Contains information about the Amazon Kinesis data stream where you are sending real-time log data.
	KinesisStreamConfig interface{} `field:"required" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// The type of data stream where you are sending real-time log data.
	//
	// The only valid value is `Kinesis` .
	StreamType *string `field:"required" json:"streamType" yaml:"streamType"`
}


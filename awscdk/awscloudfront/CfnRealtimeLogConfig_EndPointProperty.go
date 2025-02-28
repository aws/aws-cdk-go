package awscloudfront


// Contains information about the Amazon Kinesis data stream where you are sending real-time log data in a real-time log configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endPointProperty := &EndPointProperty{
//   	KinesisStreamConfig: &KinesisStreamConfigProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		StreamArn: jsii.String("streamArn"),
//   	},
//   	StreamType: jsii.String("streamType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-realtimelogconfig-endpoint.html
//
type CfnRealtimeLogConfig_EndPointProperty struct {
	// Contains information about the Amazon Kinesis data stream where you are sending real-time log data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-realtimelogconfig-endpoint.html#cfn-cloudfront-realtimelogconfig-endpoint-kinesisstreamconfig
	//
	KinesisStreamConfig interface{} `field:"required" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// The type of data stream where you are sending real-time log data.
	//
	// The only valid value is `Kinesis` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-realtimelogconfig-endpoint.html#cfn-cloudfront-realtimelogconfig-endpoint-streamtype
	//
	StreamType *string `field:"required" json:"streamType" yaml:"streamType"`
}


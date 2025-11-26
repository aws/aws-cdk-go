package previewawscloudfrontmixins


// Contains information about the Amazon Kinesis data stream where you are sending real-time log data for this real-time log configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnRealtimeLogConfigPropsMixin_EndPointProperty struct {
	// Contains information about the Amazon Kinesis data stream where you are sending real-time log data in a real-time log configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-realtimelogconfig-endpoint.html#cfn-cloudfront-realtimelogconfig-endpoint-kinesisstreamconfig
	//
	KinesisStreamConfig interface{} `field:"optional" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// The type of data stream where you are sending real-time log data.
	//
	// The only valid value is `Kinesis` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-realtimelogconfig-endpoint.html#cfn-cloudfront-realtimelogconfig-endpoint-streamtype
	//
	StreamType *string `field:"optional" json:"streamType" yaml:"streamType"`
}


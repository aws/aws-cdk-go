package mixinsawscloudfront


// Properties for CfnRealtimeLogConfigPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRealtimeLogConfigMixinProps := &CfnRealtimeLogConfigMixinProps{
//   	EndPoints: []interface{}{
//   		&EndPointProperty{
//   			KinesisStreamConfig: &KinesisStreamConfigProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				StreamArn: jsii.String("streamArn"),
//   			},
//   			StreamType: jsii.String("streamType"),
//   		},
//   	},
//   	Fields: []*string{
//   		jsii.String("fields"),
//   	},
//   	Name: jsii.String("name"),
//   	SamplingRate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html
//
type CfnRealtimeLogConfigMixinProps struct {
	// Contains information about the Amazon Kinesis data stream where you are sending real-time log data for this real-time log configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-endpoints
	//
	EndPoints interface{} `field:"optional" json:"endPoints" yaml:"endPoints"`
	// A list of fields that are included in each real-time log record.
	//
	// In an API response, the fields are provided in the same order in which they are sent to the Amazon Kinesis data stream.
	//
	// For more information about fields, see [Real-time log configuration fields](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-fields
	//
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// The unique name of this real-time log configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sampling rate for this real-time log configuration.
	//
	// The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. The sampling rate is an integer between 1 and 100, inclusive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-realtimelogconfig.html#cfn-cloudfront-realtimelogconfig-samplingrate
	//
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
}


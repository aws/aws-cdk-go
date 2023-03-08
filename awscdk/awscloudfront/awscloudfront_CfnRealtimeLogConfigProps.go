package awscloudfront


// Properties for defining a `CfnRealtimeLogConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRealtimeLogConfigProps := &cfnRealtimeLogConfigProps{
//   	endPoints: []interface{}{
//   		&endPointProperty{
//   			kinesisStreamConfig: &kinesisStreamConfigProperty{
//   				roleArn: jsii.String("roleArn"),
//   				streamArn: jsii.String("streamArn"),
//   			},
//   			streamType: jsii.String("streamType"),
//   		},
//   	},
//   	fields: []*string{
//   		jsii.String("fields"),
//   	},
//   	name: jsii.String("name"),
//   	samplingRate: jsii.Number(123),
//   }
//
type CfnRealtimeLogConfigProps struct {
	// Contains information about the Amazon Kinesis data stream where you are sending real-time log data for this real-time log configuration.
	EndPoints interface{} `field:"required" json:"endPoints" yaml:"endPoints"`
	// A list of fields that are included in each real-time log record.
	//
	// In an API response, the fields are provided in the same order in which they are sent to the Amazon Kinesis data stream.
	//
	// For more information about fields, see [Real-time log configuration fields](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) in the *Amazon CloudFront Developer Guide* .
	Fields *[]*string `field:"required" json:"fields" yaml:"fields"`
	// The unique name of this real-time log configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The sampling rate for this real-time log configuration.
	//
	// The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. The sampling rate is an integer between 1 and 100, inclusive.
	SamplingRate *float64 `field:"required" json:"samplingRate" yaml:"samplingRate"`
}


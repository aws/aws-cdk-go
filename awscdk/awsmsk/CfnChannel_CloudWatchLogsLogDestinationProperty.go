package awsmsk


// CloudWatch Logs log destination details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsLogDestinationProperty := &CloudWatchLogsLogDestinationProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	LogGroup: jsii.String("logGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-cloudwatchlogslogdestination.html
//
type CfnChannel_CloudWatchLogsLogDestinationProperty struct {
	// Whether CloudWatch Logs logging is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-cloudwatchlogslogdestination.html#cfn-msk-channel-cloudwatchlogslogdestination-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The CloudWatch log group for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-cloudwatchlogslogdestination.html#cfn-msk-channel-cloudwatchlogslogdestination-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}


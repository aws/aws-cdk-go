package awskinesisanalyticsv2


// Provides a description of Amazon CloudWatch logging options, including the log stream Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLoggingOptionProperty := &CloudWatchLoggingOptionProperty{
//   	LogStreamArn: jsii.String("logStreamArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationcloudwatchloggingoption-cloudwatchloggingoption.html
//
type CfnApplicationCloudWatchLoggingOption_CloudWatchLoggingOptionProperty struct {
	// The ARN of the CloudWatch log to receive application messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationcloudwatchloggingoption-cloudwatchloggingoption.html#cfn-kinesisanalyticsv2-applicationcloudwatchloggingoption-cloudwatchloggingoption-logstreamarn
	//
	LogStreamArn *string `field:"required" json:"logStreamArn" yaml:"logStreamArn"`
}


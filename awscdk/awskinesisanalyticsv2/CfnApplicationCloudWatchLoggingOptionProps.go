package awskinesisanalyticsv2


// Properties for defining a `CfnApplicationCloudWatchLoggingOption`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationCloudWatchLoggingOptionProps := &CfnApplicationCloudWatchLoggingOptionProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	CloudWatchLoggingOption: &CloudWatchLoggingOptionProperty{
//   		LogStreamArn: jsii.String("logStreamArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html
//
type CfnApplicationCloudWatchLoggingOptionProps struct {
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html#cfn-kinesisanalyticsv2-applicationcloudwatchloggingoption-applicationname
	//
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Provides a description of Amazon CloudWatch logging options, including the log stream Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html#cfn-kinesisanalyticsv2-applicationcloudwatchloggingoption-cloudwatchloggingoption
	//
	CloudWatchLoggingOption interface{} `field:"required" json:"cloudWatchLoggingOption" yaml:"cloudWatchLoggingOption"`
}


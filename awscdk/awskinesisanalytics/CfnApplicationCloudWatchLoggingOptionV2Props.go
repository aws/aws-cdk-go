package awskinesisanalytics


// Properties for defining a `CfnApplicationCloudWatchLoggingOption`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationCloudWatchLoggingOptionV2Props := &CfnApplicationCloudWatchLoggingOptionV2Props{
//   	ApplicationName: jsii.String("applicationName"),
//   	CloudWatchLoggingOption: &CloudWatchLoggingOptionProperty{
//   		LogStreamArn: jsii.String("logStreamArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html
//
// Deprecated: use `aws-kinesisanalyticsv2` instead.
type CfnApplicationCloudWatchLoggingOptionV2Props struct {
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html#cfn-kinesisanalyticsv2-applicationcloudwatchloggingoption-applicationname
	//
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Provides a description of Amazon CloudWatch logging options, including the log stream Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html#cfn-kinesisanalyticsv2-applicationcloudwatchloggingoption-cloudwatchloggingoption
	//
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	CloudWatchLoggingOption interface{} `field:"required" json:"cloudWatchLoggingOption" yaml:"cloudWatchLoggingOption"`
}


package awskinesisanalyticsv2


// Properties for defining a `CfnApplicationCloudWatchLoggingOption`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationCloudWatchLoggingOptionProps := &cfnApplicationCloudWatchLoggingOptionProps{
//   	applicationName: jsii.String("applicationName"),
//   	cloudWatchLoggingOption: &cloudWatchLoggingOptionProperty{
//   		logStreamArn: jsii.String("logStreamArn"),
//   	},
//   }
//
type CfnApplicationCloudWatchLoggingOptionProps struct {
	// The name of the application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Provides a description of Amazon CloudWatch logging options, including the log stream Amazon Resource Name (ARN).
	CloudWatchLoggingOption interface{} `field:"required" json:"cloudWatchLoggingOption" yaml:"cloudWatchLoggingOption"`
}


package awskinesisanalytics


// Properties for defining a `CfnApplicationCloudWatchLoggingOptionV2`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationCloudWatchLoggingOptionV2Props := &cfnApplicationCloudWatchLoggingOptionV2Props{
//   	applicationName: jsii.String("applicationName"),
//   	cloudWatchLoggingOption: &cloudWatchLoggingOptionProperty{
//   		logStreamArn: jsii.String("logStreamArn"),
//   	},
//   }
//
type CfnApplicationCloudWatchLoggingOptionV2Props struct {
	// The name of the application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Provides a description of Amazon CloudWatch logging options, including the log stream Amazon Resource Name (ARN).
	CloudWatchLoggingOption interface{} `field:"required" json:"cloudWatchLoggingOption" yaml:"cloudWatchLoggingOption"`
}


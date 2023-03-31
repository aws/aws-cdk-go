package awsivschat


// The CloudWatchLogsDestinationConfiguration property type specifies a CloudWatch Logs location where chat logs will be stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsDestinationConfigurationProperty := &cloudWatchLogsDestinationConfigurationProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   }
//
type CfnLoggingConfiguration_CloudWatchLogsDestinationConfigurationProperty struct {
	// Name of the Amazon Cloudwatch Logs destination where chat activity will be logged.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}


package awskinesisanalytics


// Provides a description of Amazon CloudWatch logging options, including the log stream Amazon Resource Name (ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLoggingOptionProperty := &cloudWatchLoggingOptionProperty{
//   	logStreamArn: jsii.String("logStreamArn"),
//   }
//
type CfnApplicationCloudWatchLoggingOptionV2_CloudWatchLoggingOptionProperty struct {
	// The ARN of the CloudWatch log to receive application messages.
	LogStreamArn *string `field:"required" json:"logStreamArn" yaml:"logStreamArn"`
}


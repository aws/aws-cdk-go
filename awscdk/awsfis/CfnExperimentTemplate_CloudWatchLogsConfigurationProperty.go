package awsfis


// Specifies the configuration for experiment logging to CloudWatch Logs .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsConfigurationProperty := &CloudWatchLogsConfigurationProperty{
//   	LogGroupArn: jsii.String("logGroupArn"),
//   }
//
type CfnExperimentTemplate_CloudWatchLogsConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the destination Amazon CloudWatch Logs log group.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}


package awsaps


// The LoggingConfiguration attribute sets the logging configuration for the workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationProperty := &LoggingConfigurationProperty{
//   	LogGroupArn: jsii.String("logGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingconfiguration.html
//
type CfnWorkspace_LoggingConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the CloudWatch log group the logs are emitted to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingconfiguration.html#cfn-aps-workspace-loggingconfiguration-loggrouparn
	//
	LogGroupArn *string `field:"optional" json:"logGroupArn" yaml:"logGroupArn"`
}


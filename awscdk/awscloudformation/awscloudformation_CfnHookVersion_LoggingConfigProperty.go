package awscloudformation


// The `LoggingConfig` property type specifies logging configuration information for an extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigProperty := &loggingConfigProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   	logRoleArn: jsii.String("logRoleArn"),
//   }
//
type CfnHookVersion_LoggingConfigProperty struct {
	// The Amazon CloudWatch Logs group to which CloudFormation sends error logging information when invoking the extension's handlers.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The Amazon Resource Name (ARN) of the role that CloudFormation should assume when sending log entries to CloudWatch Logs.
	LogRoleArn *string `field:"optional" json:"logRoleArn" yaml:"logRoleArn"`
}


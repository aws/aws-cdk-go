// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Contains logging configuration information for an extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigProperty := &loggingConfigProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   	logRoleArn: jsii.String("logRoleArn"),
//   }
//
type CfnTypeActivation_LoggingConfigProperty struct {
	// The Amazon CloudWatch Logs group to which CloudFormation sends error logging information when invoking the extension's handlers.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The Amazon Resource Name (ARN) of the role that CloudFormation should assume when sending log entries to CloudWatch Logs.
	LogRoleArn *string `field:"optional" json:"logRoleArn" yaml:"logRoleArn"`
}


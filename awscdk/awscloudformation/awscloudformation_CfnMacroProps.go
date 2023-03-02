package awscloudformation


// Properties for defining a `CfnMacro`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMacroProps := &cfnMacroProps{
//   	functionName: jsii.String("functionName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	logGroupName: jsii.String("logGroupName"),
//   	logRoleArn: jsii.String("logRoleArn"),
//   }
//
type CfnMacroProps struct {
	// The Amazon Resource Name (ARN) of the underlying AWS Lambda function that you want AWS CloudFormation to invoke when the macro is run.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The name of the macro.
	//
	// The name of the macro must be unique across all macros in the account.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the macro.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The CloudWatch Logs group to which AWS CloudFormation sends error logging information when invoking the macro's underlying AWS Lambda function.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The ARN of the role AWS CloudFormation should assume when sending log entries to CloudWatch Logs .
	LogRoleArn *string `field:"optional" json:"logRoleArn" yaml:"logRoleArn"`
}


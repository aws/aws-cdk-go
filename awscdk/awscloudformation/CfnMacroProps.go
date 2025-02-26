package awscloudformation


// Properties for defining a `CfnMacro`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMacroProps := &CfnMacroProps{
//   	FunctionName: jsii.String("functionName"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	LogRoleArn: jsii.String("logRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html
//
type CfnMacroProps struct {
	// The Amazon Resource Name (ARN) of the underlying Lambda function that you want CloudFormation to invoke when the macro is run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html#cfn-cloudformation-macro-functionname
	//
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The name of the macro.
	//
	// The name of the macro must be unique across all macros in the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html#cfn-cloudformation-macro-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the macro.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html#cfn-cloudformation-macro-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The CloudWatch Logs group to which CloudFormation sends error logging information when invoking the macro's underlying Lambda function.
	//
	// This will be an existing CloudWatch Logs LogGroup. Neither CloudFormation or Lambda will create the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html#cfn-cloudformation-macro-loggroupname
	//
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The ARN of the role CloudFormation should assume when sending log entries to CloudWatch Logs .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-macro.html#cfn-cloudformation-macro-logrolearn
	//
	LogRoleArn *string `field:"optional" json:"logRoleArn" yaml:"logRoleArn"`
}


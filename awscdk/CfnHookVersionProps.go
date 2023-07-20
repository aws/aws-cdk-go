package awscdk


// Properties for defining a `CfnHookVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHookVersionProps := &CfnHookVersionProps{
//   	SchemaHandlerPackage: jsii.String("schemaHandlerPackage"),
//   	TypeName: jsii.String("typeName"),
//
//   	// the properties below are optional
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	LoggingConfig: &LoggingConfigProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogRoleArn: jsii.String("logRoleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hookversion.html
//
type CfnHookVersionProps struct {
	// A URL to the Amazon S3 bucket containing the hook project package that contains the necessary files for the hook you want to register.
	//
	// For information on generating a schema handler package for the resource you want to register, see [submit](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-cli-submit.html) in the *CloudFormation CLI User Guide for Extension Development* .
	//
	// > The user registering the resource must be able to access the package in the S3 bucket. That's, the user must have [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html) permissions for the schema handler package. For more information, see [Actions, Resources, and Condition Keys for Amazon S3](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazons3.html) in the *AWS Identity and Access Management User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hookversion.html#cfn-cloudformation-hookversion-schemahandlerpackage
	//
	SchemaHandlerPackage *string `field:"required" json:"schemaHandlerPackage" yaml:"schemaHandlerPackage"`
	// The unique name for your hook.
	//
	// Specifies a three-part namespace for your hook, with a recommended pattern of `Organization::Service::Hook` .
	//
	// > The following organization namespaces are reserved and can't be used in your hook type names:
	// >
	// > - `Alexa`
	// > - `AMZN`
	// > - `Amazon`
	// > - `ASK`
	// > - `AWS`
	// > - `Custom`
	// > - `Dev`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hookversion.html#cfn-cloudformation-hookversion-typename
	//
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The Amazon Resource Name (ARN) of the task execution role that grants the hook permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hookversion.html#cfn-cloudformation-hookversion-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Contains logging configuration information for an extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hookversion.html#cfn-cloudformation-hookversion-loggingconfig
	//
	LoggingConfig interface{} `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
}


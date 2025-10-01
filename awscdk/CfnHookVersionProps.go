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
	// A URL to the Amazon S3 bucket for the Hook project package that contains the necessary files for the Hook you want to register.
	//
	// For information on generating a schema handler package, see [Modeling custom CloudFormation Hooks](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-model.html) in the *AWS CloudFormation Hooks User Guide* .
	//
	// > To register the Hook, you must have `s3:GetObject` permissions to access the S3 objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hookversion.html#cfn-cloudformation-hookversion-schemahandlerpackage
	//
	SchemaHandlerPackage *string `field:"required" json:"schemaHandlerPackage" yaml:"schemaHandlerPackage"`
	// The unique name for your Hook.
	//
	// Specifies a three-part namespace for your Hook, with a recommended pattern of `Organization::Service::Hook` .
	//
	// > The following organization namespaces are reserved and can't be used in your Hook type names:
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
	// The Amazon Resource Name (ARN) of the task execution role that grants the Hook permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hookversion.html#cfn-cloudformation-hookversion-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Contains logging configuration information for an extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-hookversion.html#cfn-cloudformation-hookversion-loggingconfig
	//
	LoggingConfig interface{} `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
}


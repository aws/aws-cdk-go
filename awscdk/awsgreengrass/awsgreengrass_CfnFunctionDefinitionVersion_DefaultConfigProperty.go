package awsgreengrass


// The default configuration that applies to all Lambda functions in the function definition version.
//
// Individual Lambda functions can override these settings.
//
// In an AWS CloudFormation template, `DefaultConfig` is a property of the [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultConfigProperty := &defaultConfigProperty{
//   	execution: &executionProperty{
//   		isolationMode: jsii.String("isolationMode"),
//   		runAs: &runAsProperty{
//   			gid: jsii.Number(123),
//   			uid: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnFunctionDefinitionVersion_DefaultConfigProperty struct {
	// Configuration settings for the Lambda execution environment on the AWS IoT Greengrass core.
	Execution interface{} `field:"required" json:"execution" yaml:"execution"`
}


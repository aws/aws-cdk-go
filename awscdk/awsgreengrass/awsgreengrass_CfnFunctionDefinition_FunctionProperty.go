package awsgreengrass


// A function is a Lambda function that's referenced from an AWS IoT Greengrass group.
//
// The function is deployed to a Greengrass core where it runs locally. For more information, see [Run Lambda Functions on the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-functions.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, the `Functions` property of the [`FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.html) property type contains a list of `Function` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   functionProperty := &functionProperty{
//   	functionArn: jsii.String("functionArn"),
//   	functionConfiguration: &functionConfigurationProperty{
//   		encodingType: jsii.String("encodingType"),
//   		environment: &environmentProperty{
//   			accessSysfs: jsii.Boolean(false),
//   			execution: &executionProperty{
//   				isolationMode: jsii.String("isolationMode"),
//   				runAs: &runAsProperty{
//   					gid: jsii.Number(123),
//   					uid: jsii.Number(123),
//   				},
//   			},
//   			resourceAccessPolicies: []interface{}{
//   				&resourceAccessPolicyProperty{
//   					resourceId: jsii.String("resourceId"),
//
//   					// the properties below are optional
//   					permission: jsii.String("permission"),
//   				},
//   			},
//   			variables: variables,
//   		},
//   		execArgs: jsii.String("execArgs"),
//   		executable: jsii.String("executable"),
//   		memorySize: jsii.Number(123),
//   		pinned: jsii.Boolean(false),
//   		timeout: jsii.Number(123),
//   	},
//   	id: jsii.String("id"),
//   }
//
type CfnFunctionDefinition_FunctionProperty struct {
	// The Amazon Resource Name (ARN) of the alias (recommended) or version of the referenced Lambda function.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The group-specific settings of the Lambda function.
	//
	// These settings configure the function's behavior in the Greengrass group.
	FunctionConfiguration interface{} `field:"required" json:"functionConfiguration" yaml:"functionConfiguration"`
	// A descriptive or arbitrary ID for the function.
	//
	// This value must be unique within the function definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	Id *string `field:"required" json:"id" yaml:"id"`
}


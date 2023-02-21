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
//   functionProperty := &FunctionProperty{
//   	FunctionArn: jsii.String("functionArn"),
//   	FunctionConfiguration: &FunctionConfigurationProperty{
//   		EncodingType: jsii.String("encodingType"),
//   		Environment: &EnvironmentProperty{
//   			AccessSysfs: jsii.Boolean(false),
//   			Execution: &ExecutionProperty{
//   				IsolationMode: jsii.String("isolationMode"),
//   				RunAs: &RunAsProperty{
//   					Gid: jsii.Number(123),
//   					Uid: jsii.Number(123),
//   				},
//   			},
//   			ResourceAccessPolicies: []interface{}{
//   				&ResourceAccessPolicyProperty{
//   					ResourceId: jsii.String("resourceId"),
//
//   					// the properties below are optional
//   					Permission: jsii.String("permission"),
//   				},
//   			},
//   			Variables: variables,
//   		},
//   		ExecArgs: jsii.String("execArgs"),
//   		Executable: jsii.String("executable"),
//   		MemorySize: jsii.Number(123),
//   		Pinned: jsii.Boolean(false),
//   		Timeout: jsii.Number(123),
//   	},
//   	Id: jsii.String("id"),
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


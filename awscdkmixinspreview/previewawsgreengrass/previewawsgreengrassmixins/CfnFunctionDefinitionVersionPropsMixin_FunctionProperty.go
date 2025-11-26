package previewawsgreengrassmixins


// A function is a Lambda function that's referenced from an AWS IoT Greengrass group.
//
// The function is deployed to a Greengrass core where it runs locally. For more information, see [Run Lambda Functions on the AWS IoT Greengrass Core](https://docs.aws.amazon.com/greengrass/v1/developerguide/lambda-functions.html) in the *Developer Guide* .
//
// In an CloudFormation template, the `Functions` property of the [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource contains a list of `Function` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   					Permission: jsii.String("permission"),
//   					ResourceId: jsii.String("resourceId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-function.html
//
type CfnFunctionDefinitionVersionPropsMixin_FunctionProperty struct {
	// The Amazon Resource Name (ARN) of the alias (recommended) or version of the referenced Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-function.html#cfn-greengrass-functiondefinitionversion-function-functionarn
	//
	FunctionArn *string `field:"optional" json:"functionArn" yaml:"functionArn"`
	// The group-specific settings of the Lambda function.
	//
	// These settings configure the function's behavior in the Greengrass group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-function.html#cfn-greengrass-functiondefinitionversion-function-functionconfiguration
	//
	FunctionConfiguration interface{} `field:"optional" json:"functionConfiguration" yaml:"functionConfiguration"`
	// A descriptive or arbitrary ID for the function.
	//
	// This value must be unique within the function definition version. Maximum length is 128 characters with pattern `[a-zA-Z0-9:_-]+` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-function.html#cfn-greengrass-functiondefinitionversion-function-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}


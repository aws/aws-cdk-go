package awsgreengrass


// A function definition version contains a list of functions.
//
// > After you create a function definition version that contains the functions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an CloudFormation template, `FunctionDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::FunctionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   functionDefinitionVersionProperty := &FunctionDefinitionVersionProperty{
//   	Functions: []interface{}{
//   		&FunctionProperty{
//   			FunctionArn: jsii.String("functionArn"),
//   			FunctionConfiguration: &FunctionConfigurationProperty{
//   				EncodingType: jsii.String("encodingType"),
//   				Environment: &EnvironmentProperty{
//   					AccessSysfs: jsii.Boolean(false),
//   					Execution: &ExecutionProperty{
//   						IsolationMode: jsii.String("isolationMode"),
//   						RunAs: &RunAsProperty{
//   							Gid: jsii.Number(123),
//   							Uid: jsii.Number(123),
//   						},
//   					},
//   					ResourceAccessPolicies: []interface{}{
//   						&ResourceAccessPolicyProperty{
//   							ResourceId: jsii.String("resourceId"),
//
//   							// the properties below are optional
//   							Permission: jsii.String("permission"),
//   						},
//   					},
//   					Variables: variables,
//   				},
//   				ExecArgs: jsii.String("execArgs"),
//   				Executable: jsii.String("executable"),
//   				MemorySize: jsii.Number(123),
//   				Pinned: jsii.Boolean(false),
//   				Timeout: jsii.Number(123),
//   			},
//   			Id: jsii.String("id"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	DefaultConfig: &DefaultConfigProperty{
//   		Execution: &ExecutionProperty{
//   			IsolationMode: jsii.String("isolationMode"),
//   			RunAs: &RunAsProperty{
//   				Gid: jsii.Number(123),
//   				Uid: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.html
//
type CfnFunctionDefinition_FunctionDefinitionVersionProperty struct {
	// The functions in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.html#cfn-greengrass-functiondefinition-functiondefinitionversion-functions
	//
	Functions interface{} `field:"required" json:"functions" yaml:"functions"`
	// The default configuration that applies to all Lambda functions in the group.
	//
	// Individual Lambda functions can override these settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.html#cfn-greengrass-functiondefinition-functiondefinitionversion-defaultconfig
	//
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
}


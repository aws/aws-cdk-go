package awsgreengrass


// Properties for defining a `CfnFunctionDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//   var variables interface{}
//
//   cfnFunctionDefinitionProps := &CfnFunctionDefinitionProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	InitialVersion: &FunctionDefinitionVersionProperty{
//   		Functions: []interface{}{
//   			&FunctionProperty{
//   				FunctionArn: jsii.String("functionArn"),
//   				FunctionConfiguration: &FunctionConfigurationProperty{
//   					EncodingType: jsii.String("encodingType"),
//   					Environment: &EnvironmentProperty{
//   						AccessSysfs: jsii.Boolean(false),
//   						Execution: &ExecutionProperty{
//   							IsolationMode: jsii.String("isolationMode"),
//   							RunAs: &RunAsProperty{
//   								Gid: jsii.Number(123),
//   								Uid: jsii.Number(123),
//   							},
//   						},
//   						ResourceAccessPolicies: []interface{}{
//   							&ResourceAccessPolicyProperty{
//   								ResourceId: jsii.String("resourceId"),
//
//   								// the properties below are optional
//   								Permission: jsii.String("permission"),
//   							},
//   						},
//   						Variables: variables,
//   					},
//   					ExecArgs: jsii.String("execArgs"),
//   					Executable: jsii.String("executable"),
//   					MemorySize: jsii.Number(123),
//   					Pinned: jsii.Boolean(false),
//   					Timeout: jsii.Number(123),
//   				},
//   				Id: jsii.String("id"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		DefaultConfig: &DefaultConfigProperty{
//   			Execution: &ExecutionProperty{
//   				IsolationMode: jsii.String("isolationMode"),
//   				RunAs: &RunAsProperty{
//   					Gid: jsii.Number(123),
//   					Uid: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Tags: tags,
//   }
//
type CfnFunctionDefinitionProps struct {
	// The name of the function definition.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The function definition version to include when the function definition is created.
	//
	// A function definition version contains a list of [`function`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-function.html) property types.
	//
	// > To associate a function definition version after the function definition is created, create an [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource and specify the ID of this function definition.
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// Application-specific metadata to attach to the function definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/latest/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


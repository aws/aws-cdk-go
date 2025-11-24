package mixinsawsgreengrass


// Properties for CfnFunctionDefinitionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//   var variables interface{}
//
//   cfnFunctionDefinitionMixinProps := &CfnFunctionDefinitionMixinProps{
//   	InitialVersion: &FunctionDefinitionVersionProperty{
//   		DefaultConfig: &DefaultConfigProperty{
//   			Execution: &ExecutionProperty{
//   				IsolationMode: jsii.String("isolationMode"),
//   				RunAs: &RunAsProperty{
//   					Gid: jsii.Number(123),
//   					Uid: jsii.Number(123),
//   				},
//   			},
//   		},
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
//   								Permission: jsii.String("permission"),
//   								ResourceId: jsii.String("resourceId"),
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
//   	},
//   	Name: jsii.String("name"),
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html
//
type CfnFunctionDefinitionMixinProps struct {
	// The function definition version to include when the function definition is created.
	//
	// A function definition version contains a list of [`function`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-function.html) property types.
	//
	// > To associate a function definition version after the function definition is created, create an [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource and specify the ID of this function definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html#cfn-greengrass-functiondefinition-initialversion
	//
	InitialVersion interface{} `field:"optional" json:"initialVersion" yaml:"initialVersion"`
	// The name of the function definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html#cfn-greengrass-functiondefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Application-specific metadata to attach to the function definition.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tagging Your AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/v1/developerguide/tagging.html) in the *Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html#cfn-greengrass-functiondefinition-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


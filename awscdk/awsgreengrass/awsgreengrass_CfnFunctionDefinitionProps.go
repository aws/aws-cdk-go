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
//   cfnFunctionDefinitionProps := &cfnFunctionDefinitionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	initialVersion: &functionDefinitionVersionProperty{
//   		functions: []interface{}{
//   			&functionProperty{
//   				functionArn: jsii.String("functionArn"),
//   				functionConfiguration: &functionConfigurationProperty{
//   					encodingType: jsii.String("encodingType"),
//   					environment: &environmentProperty{
//   						accessSysfs: jsii.Boolean(false),
//   						execution: &executionProperty{
//   							isolationMode: jsii.String("isolationMode"),
//   							runAs: &runAsProperty{
//   								gid: jsii.Number(123),
//   								uid: jsii.Number(123),
//   							},
//   						},
//   						resourceAccessPolicies: []interface{}{
//   							&resourceAccessPolicyProperty{
//   								resourceId: jsii.String("resourceId"),
//
//   								// the properties below are optional
//   								permission: jsii.String("permission"),
//   							},
//   						},
//   						variables: variables,
//   					},
//   					execArgs: jsii.String("execArgs"),
//   					executable: jsii.String("executable"),
//   					memorySize: jsii.Number(123),
//   					pinned: jsii.Boolean(false),
//   					timeout: jsii.Number(123),
//   				},
//   				id: jsii.String("id"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		defaultConfig: &defaultConfigProperty{
//   			execution: &executionProperty{
//   				isolationMode: jsii.String("isolationMode"),
//   				runAs: &runAsProperty{
//   					gid: jsii.Number(123),
//   					uid: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	tags: tags,
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


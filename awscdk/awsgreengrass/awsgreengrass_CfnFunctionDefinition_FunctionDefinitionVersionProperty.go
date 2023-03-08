package awsgreengrass


// A function definition version contains a list of functions.
//
// > After you create a function definition version that contains the functions you want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) .
//
// In an AWS CloudFormation template, `FunctionDefinitionVersion` is the property type of the `InitialVersion` property in the [`AWS::Greengrass::FunctionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   functionDefinitionVersionProperty := &functionDefinitionVersionProperty{
//   	functions: []interface{}{
//   		&functionProperty{
//   			functionArn: jsii.String("functionArn"),
//   			functionConfiguration: &functionConfigurationProperty{
//   				encodingType: jsii.String("encodingType"),
//   				environment: &environmentProperty{
//   					accessSysfs: jsii.Boolean(false),
//   					execution: &executionProperty{
//   						isolationMode: jsii.String("isolationMode"),
//   						runAs: &runAsProperty{
//   							gid: jsii.Number(123),
//   							uid: jsii.Number(123),
//   						},
//   					},
//   					resourceAccessPolicies: []interface{}{
//   						&resourceAccessPolicyProperty{
//   							resourceId: jsii.String("resourceId"),
//
//   							// the properties below are optional
//   							permission: jsii.String("permission"),
//   						},
//   					},
//   					variables: variables,
//   				},
//   				execArgs: jsii.String("execArgs"),
//   				executable: jsii.String("executable"),
//   				memorySize: jsii.Number(123),
//   				pinned: jsii.Boolean(false),
//   				timeout: jsii.Number(123),
//   			},
//   			id: jsii.String("id"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	defaultConfig: &defaultConfigProperty{
//   		execution: &executionProperty{
//   			isolationMode: jsii.String("isolationMode"),
//   			runAs: &runAsProperty{
//   				gid: jsii.Number(123),
//   				uid: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnFunctionDefinition_FunctionDefinitionVersionProperty struct {
	// The functions in this version.
	Functions interface{} `field:"required" json:"functions" yaml:"functions"`
	// The default configuration that applies to all Lambda functions in the group.
	//
	// Individual Lambda functions can override these settings.
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
}


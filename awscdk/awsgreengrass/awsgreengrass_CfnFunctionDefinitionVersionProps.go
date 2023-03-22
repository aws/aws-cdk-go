package awsgreengrass


// Properties for defining a `CfnFunctionDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   cfnFunctionDefinitionVersionProps := &cfnFunctionDefinitionVersionProps{
//   	functionDefinitionId: jsii.String("functionDefinitionId"),
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
type CfnFunctionDefinitionVersionProps struct {
	// The ID of the function definition associated with this version.
	//
	// This value is a GUID.
	FunctionDefinitionId *string `field:"required" json:"functionDefinitionId" yaml:"functionDefinitionId"`
	// The functions in this version.
	Functions interface{} `field:"required" json:"functions" yaml:"functions"`
	// The default configuration that applies to all Lambda functions in the group.
	//
	// Individual Lambda functions can override these settings.
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
}


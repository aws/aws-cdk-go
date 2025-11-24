package mixinsawsgreengrass


// Properties for CfnFunctionDefinitionVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var variables interface{}
//
//   cfnFunctionDefinitionVersionMixinProps := &CfnFunctionDefinitionVersionMixinProps{
//   	DefaultConfig: &DefaultConfigProperty{
//   		Execution: &ExecutionProperty{
//   			IsolationMode: jsii.String("isolationMode"),
//   			RunAs: &RunAsProperty{
//   				Gid: jsii.Number(123),
//   				Uid: jsii.Number(123),
//   			},
//   		},
//   	},
//   	FunctionDefinitionId: jsii.String("functionDefinitionId"),
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
//   							Permission: jsii.String("permission"),
//   							ResourceId: jsii.String("resourceId"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html
//
type CfnFunctionDefinitionVersionMixinProps struct {
	// The default configuration that applies to all Lambda functions in the group.
	//
	// Individual Lambda functions can override these settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html#cfn-greengrass-functiondefinitionversion-defaultconfig
	//
	DefaultConfig interface{} `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// The ID of the function definition associated with this version.
	//
	// This value is a GUID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html#cfn-greengrass-functiondefinitionversion-functiondefinitionid
	//
	FunctionDefinitionId *string `field:"optional" json:"functionDefinitionId" yaml:"functionDefinitionId"`
	// The functions in this version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html#cfn-greengrass-functiondefinitionversion-functions
	//
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
}


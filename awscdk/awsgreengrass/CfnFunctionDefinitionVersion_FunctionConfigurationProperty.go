package awsgreengrass


// The group-specific configuration settings for a Lambda function.
//
// These settings configure the function's behavior in the Greengrass group. For more information, see [Controlling Execution of Greengrass Lambda Functions by Using Group-Specific Configuration](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `FunctionConfiguration` is a property of the [`Function`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-function.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   functionConfigurationProperty := &FunctionConfigurationProperty{
//   	EncodingType: jsii.String("encodingType"),
//   	Environment: &EnvironmentProperty{
//   		AccessSysfs: jsii.Boolean(false),
//   		Execution: &ExecutionProperty{
//   			IsolationMode: jsii.String("isolationMode"),
//   			RunAs: &RunAsProperty{
//   				Gid: jsii.Number(123),
//   				Uid: jsii.Number(123),
//   			},
//   		},
//   		ResourceAccessPolicies: []interface{}{
//   			&ResourceAccessPolicyProperty{
//   				ResourceId: jsii.String("resourceId"),
//
//   				// the properties below are optional
//   				Permission: jsii.String("permission"),
//   			},
//   		},
//   		Variables: variables,
//   	},
//   	ExecArgs: jsii.String("execArgs"),
//   	Executable: jsii.String("executable"),
//   	MemorySize: jsii.Number(123),
//   	Pinned: jsii.Boolean(false),
//   	Timeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-functionconfiguration.html
//
type CfnFunctionDefinitionVersion_FunctionConfigurationProperty struct {
	// The expected encoding type of the input payload for the function.
	//
	// Valid values are `json` (default) and `binary` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-functionconfiguration.html#cfn-greengrass-functiondefinitionversion-functionconfiguration-encodingtype
	//
	EncodingType *string `field:"optional" json:"encodingType" yaml:"encodingType"`
	// The environment configuration of the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-functionconfiguration.html#cfn-greengrass-functiondefinitionversion-functionconfiguration-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The execution arguments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-functionconfiguration.html#cfn-greengrass-functiondefinitionversion-functionconfiguration-execargs
	//
	ExecArgs *string `field:"optional" json:"execArgs" yaml:"execArgs"`
	// The name of the function executable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-functionconfiguration.html#cfn-greengrass-functiondefinitionversion-functionconfiguration-executable
	//
	Executable *string `field:"optional" json:"executable" yaml:"executable"`
	// The memory size (in KB) required by the function.
	//
	// > This property applies only to Lambda functions that run in a Greengrass container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-functionconfiguration.html#cfn-greengrass-functiondefinitionversion-functionconfiguration-memorysize
	//
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Indicates whether the function is pinned (or *long-lived* ).
	//
	// Pinned functions start when the core starts and process all requests in the same container. The default value is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-functionconfiguration.html#cfn-greengrass-functiondefinitionversion-functionconfiguration-pinned
	//
	Pinned interface{} `field:"optional" json:"pinned" yaml:"pinned"`
	// The allowed execution time (in seconds) after which the function should terminate.
	//
	// For pinned functions, this timeout applies for each request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-functionconfiguration.html#cfn-greengrass-functiondefinitionversion-functionconfiguration-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}


package awsgreengrass


// The environment configuration for a Lambda function on the AWS IoT Greengrass core.
//
// In an AWS CloudFormation template, `Environment` is a property of the [`FunctionConfiguration`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-functionconfiguration.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   environmentProperty := &EnvironmentProperty{
//   	AccessSysfs: jsii.Boolean(false),
//   	Execution: &ExecutionProperty{
//   		IsolationMode: jsii.String("isolationMode"),
//   		RunAs: &RunAsProperty{
//   			Gid: jsii.Number(123),
//   			Uid: jsii.Number(123),
//   		},
//   	},
//   	ResourceAccessPolicies: []interface{}{
//   		&ResourceAccessPolicyProperty{
//   			ResourceId: jsii.String("resourceId"),
//
//   			// the properties below are optional
//   			Permission: jsii.String("permission"),
//   		},
//   	},
//   	Variables: variables,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html
//
type CfnFunctionDefinitionVersion_EnvironmentProperty struct {
	// Indicates whether the function is allowed to access the `/sys` directory on the core device, which allows the read device information from `/sys` .
	//
	// > This property applies only to Lambda functions that run in a Greengrass container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html#cfn-greengrass-functiondefinitionversion-environment-accesssysfs
	//
	AccessSysfs interface{} `field:"optional" json:"accessSysfs" yaml:"accessSysfs"`
	// Settings for the Lambda execution environment in AWS IoT Greengrass .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html#cfn-greengrass-functiondefinitionversion-environment-execution
	//
	Execution interface{} `field:"optional" json:"execution" yaml:"execution"`
	// A list of the [resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.html) in the group that the function can access, with the corresponding read-only or read-write permissions. The maximum is 10 resources.
	//
	// > This property applies only to Lambda functions that run in a Greengrass container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html#cfn-greengrass-functiondefinitionversion-environment-resourceaccesspolicies
	//
	ResourceAccessPolicies interface{} `field:"optional" json:"resourceAccessPolicies" yaml:"resourceAccessPolicies"`
	// Environment variables for the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-environment.html#cfn-greengrass-functiondefinitionversion-environment-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}


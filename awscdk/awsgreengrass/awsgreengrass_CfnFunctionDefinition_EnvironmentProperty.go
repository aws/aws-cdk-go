package awsgreengrass


// The environment configuration for a Lambda function on the AWS IoT Greengrass core.
//
// In an AWS CloudFormation template, `Environment` is a property of the [`FunctionConfiguration`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-functionconfiguration.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var variables interface{}
//
//   environmentProperty := &environmentProperty{
//   	accessSysfs: jsii.Boolean(false),
//   	execution: &executionProperty{
//   		isolationMode: jsii.String("isolationMode"),
//   		runAs: &runAsProperty{
//   			gid: jsii.Number(123),
//   			uid: jsii.Number(123),
//   		},
//   	},
//   	resourceAccessPolicies: []interface{}{
//   		&resourceAccessPolicyProperty{
//   			resourceId: jsii.String("resourceId"),
//
//   			// the properties below are optional
//   			permission: jsii.String("permission"),
//   		},
//   	},
//   	variables: variables,
//   }
//
type CfnFunctionDefinition_EnvironmentProperty struct {
	// Indicates whether the function is allowed to access the `/sys` directory on the core device, which allows the read device information from `/sys` .
	//
	// > This property applies only to Lambda functions that run in a Greengrass container.
	AccessSysfs interface{} `field:"optional" json:"accessSysfs" yaml:"accessSysfs"`
	// Settings for the Lambda execution environment in AWS IoT Greengrass .
	Execution interface{} `field:"optional" json:"execution" yaml:"execution"`
	// A list of the [resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.html) in the group that the function can access, with the corresponding read-only or read-write permissions. The maximum is 10 resources.
	//
	// > This property applies only for Lambda functions that run in a Greengrass container.
	ResourceAccessPolicies interface{} `field:"optional" json:"resourceAccessPolicies" yaml:"resourceAccessPolicies"`
	// Environment variables for the Lambda function.
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}


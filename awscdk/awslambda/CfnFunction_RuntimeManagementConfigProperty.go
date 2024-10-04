package awslambda


// Sets the runtime management configuration for a function's version.
//
// For more information, see [Runtime updates](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeManagementConfigProperty := &RuntimeManagementConfigProperty{
//   	UpdateRuntimeOn: jsii.String("updateRuntimeOn"),
//
//   	// the properties below are optional
//   	RuntimeVersionArn: jsii.String("runtimeVersionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-runtimemanagementconfig.html
//
type CfnFunction_RuntimeManagementConfigProperty struct {
	// Specify the runtime update mode.
	//
	// - *Auto (default)* - Automatically update to the most recent and secure runtime version using a [Two-phase runtime version rollout](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html#runtime-management-two-phase) . This is the best choice for most customers to ensure they always benefit from runtime updates.
	// - *FunctionUpdate* - Lambda updates the runtime of you function to the most recent and secure runtime version when you update your function. This approach synchronizes runtime updates with function deployments, giving you control over when runtime updates are applied and allowing you to detect and mitigate rare runtime update incompatibilities early. When using this setting, you need to regularly update your functions to keep their runtime up-to-date.
	// - *Manual* - You specify a runtime version in your function configuration. The function will use this runtime version indefinitely. In the rare case where a new runtime version is incompatible with an existing function, this allows you to roll back your function to an earlier runtime version. For more information, see [Roll back a runtime version](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html#runtime-management-rollback) .
	//
	// *Valid Values* : `Auto` | `FunctionUpdate` | `Manual`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-runtimemanagementconfig.html#cfn-lambda-function-runtimemanagementconfig-updateruntimeon
	//
	UpdateRuntimeOn *string `field:"required" json:"updateRuntimeOn" yaml:"updateRuntimeOn"`
	// The ARN of the runtime version you want the function to use.
	//
	// > This is only required if you're using the *Manual* runtime update mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-runtimemanagementconfig.html#cfn-lambda-function-runtimemanagementconfig-runtimeversionarn
	//
	RuntimeVersionArn *string `field:"optional" json:"runtimeVersionArn" yaml:"runtimeVersionArn"`
}


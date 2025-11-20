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
	// Trigger for runtime update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-runtimemanagementconfig.html#cfn-lambda-function-runtimemanagementconfig-updateruntimeon
	//
	UpdateRuntimeOn *string `field:"required" json:"updateRuntimeOn" yaml:"updateRuntimeOn"`
	// Unique identifier for a runtime version arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-runtimemanagementconfig.html#cfn-lambda-function-runtimemanagementconfig-runtimeversionarn
	//
	RuntimeVersionArn *string `field:"optional" json:"runtimeVersionArn" yaml:"runtimeVersionArn"`
}


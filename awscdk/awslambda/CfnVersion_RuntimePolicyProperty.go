package awslambda


// Runtime Management Config of a function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimePolicyProperty := &RuntimePolicyProperty{
//   	UpdateRuntimeOn: jsii.String("updateRuntimeOn"),
//
//   	// the properties below are optional
//   	RuntimeVersionArn: jsii.String("runtimeVersionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-version-runtimepolicy.html
//
type CfnVersion_RuntimePolicyProperty struct {
	// The runtime update mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-version-runtimepolicy.html#cfn-lambda-version-runtimepolicy-updateruntimeon
	//
	UpdateRuntimeOn *string `field:"required" json:"updateRuntimeOn" yaml:"updateRuntimeOn"`
	// The ARN of the runtime the function is configured to use.
	//
	// If the runtime update mode is manual, the ARN is returned, otherwise null is returned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-version-runtimepolicy.html#cfn-lambda-version-runtimepolicy-runtimeversionarn
	//
	RuntimeVersionArn *string `field:"optional" json:"runtimeVersionArn" yaml:"runtimeVersionArn"`
}


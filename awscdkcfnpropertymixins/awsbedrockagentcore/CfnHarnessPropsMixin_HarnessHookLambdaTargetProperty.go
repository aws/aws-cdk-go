package awsbedrockagentcore


// A Lambda function invoked synchronously for a lifecycle hook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessHookLambdaTargetProperty := &HarnessHookLambdaTargetProperty{
//   	Arn: jsii.String("arn"),
//   	FailureMode: jsii.String("failureMode"),
//   	TimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshooklambdatarget.html
//
type CfnHarnessPropsMixin_HarnessHookLambdaTargetProperty struct {
	// The ARN of the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshooklambdatarget.html#cfn-bedrockagentcore-harness-harnesshooklambdatarget-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Whether the agent loop continues or stops when the Lambda invocation fails.
	//
	// Defaults to deny.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshooklambdatarget.html#cfn-bedrockagentcore-harness-harnesshooklambdatarget-failuremode
	//
	FailureMode *string `field:"optional" json:"failureMode" yaml:"failureMode"`
	// The maximum number of seconds to wait for the Lambda response.
	//
	// Defaults to 60.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshooklambdatarget.html#cfn-bedrockagentcore-harness-harnesshooklambdatarget-timeoutseconds
	//
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}


package awsbedrockagentcore


// A hook that runs before an agent invocation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessBeforeInvocationHookProperty := &HarnessBeforeInvocationHookProperty{
//   	Name: jsii.String("name"),
//   	Target: &HarnessHookTargetProperty{
//   		EventBridge: &HarnessHookEventBridgeTargetProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		Lambda: &HarnessHookLambdaTargetProperty{
//   			Arn: jsii.String("arn"),
//   			FailureMode: jsii.String("failureMode"),
//   			TimeoutSeconds: jsii.Number(123),
//   		},
//   		Sns: &HarnessHookSnsTargetProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbeforeinvocationhook.html
//
type CfnHarnessPropsMixin_HarnessBeforeInvocationHookProperty struct {
	// The unique name of the hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbeforeinvocationhook.html#cfn-bedrockagentcore-harness-harnessbeforeinvocationhook-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The destination that receives lifecycle hook events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbeforeinvocationhook.html#cfn-bedrockagentcore-harness-harnessbeforeinvocationhook-target
	//
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}


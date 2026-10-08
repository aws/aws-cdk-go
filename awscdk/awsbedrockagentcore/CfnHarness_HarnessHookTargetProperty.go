package awsbedrockagentcore


// The destination that receives lifecycle hook events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessHookTargetProperty := &HarnessHookTargetProperty{
//   	EventBridge: &HarnessHookEventBridgeTargetProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	Lambda: &HarnessHookLambdaTargetProperty{
//   		Arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		FailureMode: jsii.String("failureMode"),
//   		TimeoutSeconds: jsii.Number(123),
//   	},
//   	Sns: &HarnessHookSnsTargetProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshooktarget.html
//
type CfnHarness_HarnessHookTargetProperty struct {
	// An EventBridge event bus that receives lifecycle hook events asynchronously.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshooktarget.html#cfn-bedrockagentcore-harness-harnesshooktarget-eventbridge
	//
	EventBridge interface{} `field:"optional" json:"eventBridge" yaml:"eventBridge"`
	// A Lambda function invoked synchronously for a lifecycle hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshooktarget.html#cfn-bedrockagentcore-harness-harnesshooktarget-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
	// An SNS topic that receives lifecycle hook events asynchronously.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshooktarget.html#cfn-bedrockagentcore-harness-harnesshooktarget-sns
	//
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
}


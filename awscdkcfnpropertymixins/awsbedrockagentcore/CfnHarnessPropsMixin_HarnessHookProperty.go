package awsbedrockagentcore


// A lifecycle hook configured for one phase of the agent loop.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessHookProperty := &HarnessHookProperty{
//   	AfterInvocation: &HarnessAfterInvocationHookProperty{
//   		Name: jsii.String("name"),
//   		Target: &HarnessHookTargetProperty{
//   			EventBridge: &HarnessHookEventBridgeTargetProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			Lambda: &HarnessHookLambdaTargetProperty{
//   				Arn: jsii.String("arn"),
//   				FailureMode: jsii.String("failureMode"),
//   				TimeoutSeconds: jsii.Number(123),
//   			},
//   			Sns: &HarnessHookSnsTargetProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   		},
//   	},
//   	AfterToolCall: &HarnessAfterToolCallHookProperty{
//   		Name: jsii.String("name"),
//   		Target: &HarnessHookTargetProperty{
//   			EventBridge: &HarnessHookEventBridgeTargetProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			Lambda: &HarnessHookLambdaTargetProperty{
//   				Arn: jsii.String("arn"),
//   				FailureMode: jsii.String("failureMode"),
//   				TimeoutSeconds: jsii.Number(123),
//   			},
//   			Sns: &HarnessHookSnsTargetProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   		},
//   	},
//   	BeforeInvocation: &HarnessBeforeInvocationHookProperty{
//   		Name: jsii.String("name"),
//   		Target: &HarnessHookTargetProperty{
//   			EventBridge: &HarnessHookEventBridgeTargetProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			Lambda: &HarnessHookLambdaTargetProperty{
//   				Arn: jsii.String("arn"),
//   				FailureMode: jsii.String("failureMode"),
//   				TimeoutSeconds: jsii.Number(123),
//   			},
//   			Sns: &HarnessHookSnsTargetProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   		},
//   	},
//   	BeforeToolCall: &HarnessBeforeToolCallHookProperty{
//   		Name: jsii.String("name"),
//   		Target: &HarnessHookTargetProperty{
//   			EventBridge: &HarnessHookEventBridgeTargetProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			Lambda: &HarnessHookLambdaTargetProperty{
//   				Arn: jsii.String("arn"),
//   				FailureMode: jsii.String("failureMode"),
//   				TimeoutSeconds: jsii.Number(123),
//   			},
//   			Sns: &HarnessHookSnsTargetProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshook.html
//
type CfnHarnessPropsMixin_HarnessHookProperty struct {
	// A hook that runs after an agent invocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshook.html#cfn-bedrockagentcore-harness-harnesshook-afterinvocation
	//
	AfterInvocation interface{} `field:"optional" json:"afterInvocation" yaml:"afterInvocation"`
	// A hook that runs after each tool call.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshook.html#cfn-bedrockagentcore-harness-harnesshook-aftertoolcall
	//
	AfterToolCall interface{} `field:"optional" json:"afterToolCall" yaml:"afterToolCall"`
	// A hook that runs before an agent invocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshook.html#cfn-bedrockagentcore-harness-harnesshook-beforeinvocation
	//
	BeforeInvocation interface{} `field:"optional" json:"beforeInvocation" yaml:"beforeInvocation"`
	// A hook that runs before each tool call.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshook.html#cfn-bedrockagentcore-harness-harnesshook-beforetoolcall
	//
	BeforeToolCall interface{} `field:"optional" json:"beforeToolCall" yaml:"beforeToolCall"`
}


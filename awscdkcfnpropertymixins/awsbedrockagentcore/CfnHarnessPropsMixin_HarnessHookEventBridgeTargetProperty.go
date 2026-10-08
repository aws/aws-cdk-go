package awsbedrockagentcore


// An EventBridge event bus that receives lifecycle hook events asynchronously.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessHookEventBridgeTargetProperty := &HarnessHookEventBridgeTargetProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshookeventbridgetarget.html
//
type CfnHarnessPropsMixin_HarnessHookEventBridgeTargetProperty struct {
	// The ARN of the EventBridge event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshookeventbridgetarget.html#cfn-bedrockagentcore-harness-harnesshookeventbridgetarget-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}


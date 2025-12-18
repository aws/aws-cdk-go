package previewawsautoscalingmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceLifecyclePolicyProperty := &InstanceLifecyclePolicyProperty{
//   	RetentionTriggers: &RetentionTriggersProperty{
//   		TerminateHookAbandon: jsii.String("terminateHookAbandon"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-instancelifecyclepolicy.html
//
type CfnAutoScalingGroupPropsMixin_InstanceLifecyclePolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-instancelifecyclepolicy.html#cfn-autoscaling-autoscalinggroup-instancelifecyclepolicy-retentiontriggers
	//
	RetentionTriggers interface{} `field:"optional" json:"retentionTriggers" yaml:"retentionTriggers"`
}


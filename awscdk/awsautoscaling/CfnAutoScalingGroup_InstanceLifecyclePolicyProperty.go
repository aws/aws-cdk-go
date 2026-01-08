package awsautoscaling


// The instance lifecycle policy for the Auto Scaling group.
//
// This policy controls instance behavior when an instance transitions through its lifecycle states. Configure retention triggers to specify when instances should move to a `Retained` state instead of automatic termination.
//
// For more information, see [Control instance retention with instance lifecycle policies](https://docs.aws.amazon.com/autoscaling/ec2/userguide/instance-lifecycle-policy.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceLifecyclePolicyProperty := &InstanceLifecyclePolicyProperty{
//   	RetentionTriggers: &RetentionTriggersProperty{
//   		TerminateHookAbandon: jsii.String("terminateHookAbandon"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-instancelifecyclepolicy.html
//
type CfnAutoScalingGroup_InstanceLifecyclePolicyProperty struct {
	// Specifies the conditions that trigger instance retention behavior.
	//
	// These triggers determine when instances should move to a `Retained` state instead of automatic termination. This allows you to maintain control over instance management when lifecycles transition and operations fail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-instancelifecyclepolicy.html#cfn-autoscaling-autoscalinggroup-instancelifecyclepolicy-retentiontriggers
	//
	RetentionTriggers interface{} `field:"optional" json:"retentionTriggers" yaml:"retentionTriggers"`
}


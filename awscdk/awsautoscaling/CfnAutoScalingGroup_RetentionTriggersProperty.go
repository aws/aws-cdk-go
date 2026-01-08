package awsautoscaling


// Defines the specific triggers that cause instances to be retained in a Retained state rather than terminated.
//
// Each trigger corresponds to a different failure scenario during the instance lifecycle. This allows fine-grained control over when to preserve instances for manual intervention.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retentionTriggersProperty := &RetentionTriggersProperty{
//   	TerminateHookAbandon: jsii.String("terminateHookAbandon"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-retentiontriggers.html
//
type CfnAutoScalingGroup_RetentionTriggersProperty struct {
	// Specifies the action when a termination lifecycle hook is abandoned due to failure, timeout, or explicit abandonment (calling CompleteLifecycleAction).
	//
	// Set to `Retain` to move instances to a `Retained` state. Set to `Terminate` for default termination behavior.
	//
	// Retained instances don't count toward desired capacity and remain until you call `TerminateInstanceInAutoScalingGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-retentiontriggers.html#cfn-autoscaling-autoscalinggroup-retentiontriggers-terminatehookabandon
	//
	TerminateHookAbandon *string `field:"optional" json:"terminateHookAbandon" yaml:"terminateHookAbandon"`
}


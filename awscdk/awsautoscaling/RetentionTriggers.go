package awsautoscaling


// Defines the specific triggers that cause instances to be retained in a Retained state rather than terminated.
//
// Each trigger corresponds to a different failure scenario during
// the instance lifecycle. This allows fine-grained control over when to preserve instances
// for manual intervention.
//
// Example:
//   var vpc Vpc
//   var instanceType InstanceType
//   var machineImage IMachineImage
//
//
//   asg := autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	// Configure instance lifecycle policy
//   	InstanceLifecyclePolicy: &InstanceLifecyclePolicy{
//   		RetentionTriggers: &RetentionTriggers{
//   			TerminateHookAbandon: autoscaling.TerminateHookAbandonAction_RETAIN,
//   		},
//   	},
//   })
//
//   // Add termination lifecycle hook (required for the policy to take effect)
//   asg.addLifecycleHook(jsii.String("TerminationHook"), &BasicLifecycleHookProps{
//   	LifecycleTransition: autoscaling.LifecycleTransition_INSTANCE_TERMINATING,
//   })
//
type RetentionTriggers struct {
	// Specifies the action when a termination lifecycle hook is abandoned due to failure, timeout, or explicit abandonment (calling CompleteLifecycleAction).
	//
	// Set to Retain to move instances to a Retained state. Set to Terminate for default termination behavior.
	//
	// Retained instances don't count toward desired capacity and remain
	// until you call TerminateInstanceInAutoScalingGroup.
	// Default: - No action specified.
	//
	TerminateHookAbandon TerminateHookAbandonAction `field:"optional" json:"terminateHookAbandon" yaml:"terminateHookAbandon"`
}


package awsautoscaling


// Actions for when a termination lifecycle hook is abandoned.
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
type TerminateHookAbandonAction string

const (
	// Move instances to a Retained state when termination hook is abandoned.
	TerminateHookAbandonAction_RETAIN TerminateHookAbandonAction = "RETAIN"
	// Terminate instances normally when termination hook is abandoned (default behavior).
	TerminateHookAbandonAction_TERMINATE TerminateHookAbandonAction = "TERMINATE"
)


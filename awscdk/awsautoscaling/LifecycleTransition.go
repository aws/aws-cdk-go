package awsautoscaling


// What instance transition to attach the hook to.
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
type LifecycleTransition string

const (
	// Execute the hook when an instance is about to be added.
	LifecycleTransition_INSTANCE_LAUNCHING LifecycleTransition = "INSTANCE_LAUNCHING"
	// Execute the hook when an instance is about to be terminated.
	LifecycleTransition_INSTANCE_TERMINATING LifecycleTransition = "INSTANCE_TERMINATING"
)


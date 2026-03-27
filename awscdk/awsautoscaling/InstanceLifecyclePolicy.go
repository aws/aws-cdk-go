package awsautoscaling


// Instance lifecycle policy for an Auto Scaling group.
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
type InstanceLifecyclePolicy struct {
	// Retention triggers for the instance lifecycle policy.
	// Default: - No retention triggers configured.
	//
	RetentionTriggers *RetentionTriggers `field:"optional" json:"retentionTriggers" yaml:"retentionTriggers"`
}


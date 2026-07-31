package awsautoscaling


// The strategy to use for the instance refresh.
//
// Example:
//   var vpc Vpc
//   var instanceType InstanceType
//   var machineImage IMachineImage
//   var alarm Alarm
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	UpdatePolicy: autoscaling.UpdatePolicy_InstanceRefresh(&InstanceRefreshOptions{
//   		Strategy: autoscaling.InstanceRefreshStrategy_ROLLING,
//   		MinHealthyPercentage: jsii.Number(90),
//   		MaxHealthyPercentage: jsii.Number(100),
//   		InstanceWarmup: awscdk.Duration_Seconds(jsii.Number(300)),
//   		CheckpointPercentages: []*f64{
//   			jsii.Number(50),
//   			jsii.Number(100),
//   		},
//   		CheckpointDelay: awscdk.Duration_Minutes(jsii.Number(10)),
//   		Alarms: []IAlarmRef{
//   			alarm,
//   		},
//   	}),
//   })
//
type InstanceRefreshStrategy string

const (
	// Terminate instances and launch new ones to replace them.
	InstanceRefreshStrategy_ROLLING InstanceRefreshStrategy = "ROLLING"
	// Replace the root volume of instances in place.
	//
	// When this strategy is used, only `ImageId` changes within the launch template or mixed
	// instances policy are allowed. Other property changes may cause the stack update to fail.
	InstanceRefreshStrategy_REPLACE_ROOT_VOLUME InstanceRefreshStrategy = "REPLACE_ROOT_VOLUME"
)


package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
)

// Options for customizing the instance refresh update policy.
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
type InstanceRefreshOptions struct {
	// The strategy to use for the instance refresh.
	Strategy InstanceRefreshStrategy `field:"required" json:"strategy" yaml:"strategy"`
	// The CloudWatch alarms to monitor during the instance refresh.
	//
	// If any of the alarms goes into
	// ALARM state, the instance refresh fails. You can specify up to 10 alarms.
	// Default: - no alarms.
	//
	Alarms *[]interfacesawscloudwatch.IAlarmRef `field:"optional" json:"alarms" yaml:"alarms"`
	// The amount of time after an instance refresh completes successfully before CloudFormation considers the update successful.
	// Default: Duration.seconds(0)
	//
	BakeTime awscdk.Duration `field:"optional" json:"bakeTime" yaml:"bakeTime"`
	// The amount of time to wait after a checkpoint is reached before continuing.
	// Default: - 3600 seconds (1 hour), applied only when checkpointPercentages is set.
	//
	CheckpointDelay awscdk.Duration `field:"optional" json:"checkpointDelay" yaml:"checkpointDelay"`
	// Threshold values for each checkpoint in ascending order.
	//
	// Each number must be unique.
	// To replace all instances in the group, the last number in the array must be 100.
	// Default: - no checkpoints.
	//
	CheckpointPercentages *[]*float64 `field:"optional" json:"checkpointPercentages" yaml:"checkpointPercentages"`
	// The amount of time to wait after a new instance enters the InService state before moving on to refresh the next instance.
	// Default: - the group's DefaultInstanceWarmup if defined; otherwise the group's HealthCheckGracePeriod.
	//
	InstanceWarmup awscdk.Duration `field:"optional" json:"instanceWarmup" yaml:"instanceWarmup"`
	// The maximum percentage of the group that can be in service and healthy, or pending, to support your workload during the instance refresh.
	//
	// The difference between
	// `maxHealthyPercentage` and `minHealthyPercentage` cannot be greater than 100. A larger
	// range increases the number of instances that can be replaced at the same time.
	// Default: - the value set in the Auto Scaling group's instance maintenance policy, if defined; otherwise 110 when the strategy is Rolling, or 100 when the strategy is ReplaceRootVolume.
	//
	MaxHealthyPercentage *float64 `field:"optional" json:"maxHealthyPercentage" yaml:"maxHealthyPercentage"`
	// The minimum percentage of the group to keep in service, healthy, and ready to use to support your workload during the instance refresh, expressed as a percentage of the group's desired capacity.
	// Default: - the value set in the Auto Scaling group's instance maintenance policy, if defined; otherwise 100 when the strategy is Rolling, or 90 when the strategy is ReplaceRootVolume.
	//
	MinHealthyPercentage *float64 `field:"optional" json:"minHealthyPercentage" yaml:"minHealthyPercentage"`
	// Specifies the behavior of instances that are protected from scale in during an instance refresh.
	// Default: ScaleInProtectedInstances.WAIT
	//
	ScaleInProtectedInstances ScaleInProtectedInstances `field:"optional" json:"scaleInProtectedInstances" yaml:"scaleInProtectedInstances"`
	// Indicates whether skip matching is enabled.
	// Default: true.
	//
	SkipMatching *bool `field:"optional" json:"skipMatching" yaml:"skipMatching"`
	// Specifies the behavior of instances in standby during an instance refresh.
	// Default: StandbyInstances.WAIT
	//
	StandbyInstances StandbyInstances `field:"optional" json:"standbyInstances" yaml:"standbyInstances"`
}


package awscdk


// Preferences for the AutoScalingInstanceRefresh update policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAutoScalingInstanceRefreshPreferences := &CfnAutoScalingInstanceRefreshPreferences{
//   	AlarmSpecification: &CfnAutoScalingInstanceRefreshAlarmSpecification{
//   		Alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   	},
//   	BakeTime: jsii.Number(123),
//   	CheckpointDelay: jsii.Number(123),
//   	CheckpointPercentages: []*f64{
//   		jsii.Number(123),
//   	},
//   	InstanceWarmup: jsii.Number(123),
//   	MaxHealthyPercentage: jsii.Number(123),
//   	MinHealthyPercentage: jsii.Number(123),
//   	ScaleInProtectedInstances: jsii.String("scaleInProtectedInstances"),
//   	SkipMatching: jsii.Boolean(false),
//   	StandbyInstances: jsii.String("standbyInstances"),
//   }
//
type CfnAutoScalingInstanceRefreshPreferences struct {
	// The CloudWatch alarms to monitor during the instance refresh.
	//
	// If any alarm goes into
	// ALARM state, the instance refresh fails.
	AlarmSpecification *CfnAutoScalingInstanceRefreshAlarmSpecification `field:"optional" json:"alarmSpecification" yaml:"alarmSpecification"`
	// The number of seconds after an instance refresh completes successfully before CloudFormation considers the update successful.
	// Default: 0.
	//
	BakeTime *float64 `field:"optional" json:"bakeTime" yaml:"bakeTime"`
	// The number of seconds to wait after a checkpoint is reached before continuing.
	// Default: - 3600 seconds (1 hour), applied only when checkpointPercentages is set.
	//
	CheckpointDelay *float64 `field:"optional" json:"checkpointDelay" yaml:"checkpointDelay"`
	// Threshold values for each checkpoint in ascending order.
	//
	// Each number is a percentage of
	// the total number of instances in the group. When the percentage of instances that have
	// been replaced reaches a checkpoint, the refresh waits for the configured checkpoint delay
	// before continuing.
	// Default: - no checkpoints.
	//
	CheckpointPercentages *[]*float64 `field:"optional" json:"checkpointPercentages" yaml:"checkpointPercentages"`
	// The number of seconds to wait after a new instance enters the InService state before moving on to replacing the next instance.
	// Default: - the group's DefaultInstanceWarmup if defined; otherwise the group's HealthCheckGracePeriod.
	//
	InstanceWarmup *float64 `field:"optional" json:"instanceWarmup" yaml:"instanceWarmup"`
	// The maximum percentage of the group that can be in service and healthy, or pending, to support your workload during the instance refresh.
	//
	// Used to control batch size.
	// Default: - the value set in the Auto Scaling group's instance maintenance policy, if defined; otherwise 110 when the strategy is Rolling, or 100 when the strategy is ReplaceRootVolume.
	//
	MaxHealthyPercentage *float64 `field:"optional" json:"maxHealthyPercentage" yaml:"maxHealthyPercentage"`
	// The minimum percentage of the group to keep in service, healthy, and ready to use to support your workload during the instance refresh.
	// Default: - the value set in the Auto Scaling group's instance maintenance policy, if defined; otherwise 100 when the strategy is Rolling, or 90 when the strategy is ReplaceRootVolume.
	//
	MinHealthyPercentage *float64 `field:"optional" json:"minHealthyPercentage" yaml:"minHealthyPercentage"`
	// Specifies the behavior of instances that are protected from scale in during an instance refresh.
	// Default: 'Wait'.
	//
	ScaleInProtectedInstances *string `field:"optional" json:"scaleInProtectedInstances" yaml:"scaleInProtectedInstances"`
	// Indicates whether skip matching is enabled.
	//
	// If true, instances that already match the
	// desired configuration are not replaced.
	// Default: true.
	//
	SkipMatching *bool `field:"optional" json:"skipMatching" yaml:"skipMatching"`
	// Specifies the behavior of instances in standby during an instance refresh.
	// Default: 'Wait'.
	//
	StandbyInstances *string `field:"optional" json:"standbyInstances" yaml:"standbyInstances"`
}


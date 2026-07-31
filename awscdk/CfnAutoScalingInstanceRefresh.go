package awscdk


// To specify how AWS CloudFormation handles instance refresh for an Auto Scaling group, use the AutoScalingInstanceRefresh update policy.
//
// When properties that trigger an instance refresh change
// (such as LaunchTemplate or MixedInstancesPolicy), CloudFormation starts an instance refresh to
// replace instances gradually.
//
// This policy is mutually exclusive with AutoScalingRollingUpdate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAutoScalingInstanceRefresh := &CfnAutoScalingInstanceRefresh{
//   	Strategy: jsii.String("strategy"),
//
//   	// the properties below are optional
//   	Preferences: &CfnAutoScalingInstanceRefreshPreferences{
//   		AlarmSpecification: &CfnAutoScalingInstanceRefreshAlarmSpecification{
//   			Alarms: []*string{
//   				jsii.String("alarms"),
//   			},
//   		},
//   		BakeTime: jsii.Number(123),
//   		CheckpointDelay: jsii.Number(123),
//   		CheckpointPercentages: []*f64{
//   			jsii.Number(123),
//   		},
//   		InstanceWarmup: jsii.Number(123),
//   		MaxHealthyPercentage: jsii.Number(123),
//   		MinHealthyPercentage: jsii.Number(123),
//   		ScaleInProtectedInstances: jsii.String("scaleInProtectedInstances"),
//   		SkipMatching: jsii.Boolean(false),
//   		StandbyInstances: jsii.String("standbyInstances"),
//   	},
//   }
//
type CfnAutoScalingInstanceRefresh struct {
	// The strategy to use for the instance refresh.
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// The preferences for the instance refresh.
	Preferences *CfnAutoScalingInstanceRefreshPreferences `field:"optional" json:"preferences" yaml:"preferences"`
}


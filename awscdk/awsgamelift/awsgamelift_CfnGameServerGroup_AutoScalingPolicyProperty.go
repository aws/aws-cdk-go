package awsgamelift


// *This data type is used with the GameLift FleetIQ and game server groups.*.
//
// Configuration settings for intelligent automatic scaling that uses target tracking. After the Auto Scaling group is created, all updates to Auto Scaling policies, including changing this policy and adding or removing other policies, is done directly on the Auto Scaling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingPolicyProperty := &autoScalingPolicyProperty{
//   	targetTrackingConfiguration: &targetTrackingConfigurationProperty{
//   		targetValue: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	estimatedInstanceWarmup: jsii.Number(123),
//   }
//
type CfnGameServerGroup_AutoScalingPolicyProperty struct {
	// Settings for a target-based scaling policy applied to Auto Scaling group.
	//
	// These settings are used to create a target-based policy that tracks the GameLift FleetIQ metric `PercentUtilizedGameServers` and specifies a target value for the metric. As player usage changes, the policy triggers to adjust the game server group capacity so that the metric returns to the target value.
	TargetTrackingConfiguration interface{} `field:"required" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
	// Length of time, in seconds, it takes for a new instance to start new game server processes and register with GameLift FleetIQ.
	//
	// Specifying a warm-up time can be useful, particularly with game servers that take a long time to start up, because it avoids prematurely starting new instances.
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
}


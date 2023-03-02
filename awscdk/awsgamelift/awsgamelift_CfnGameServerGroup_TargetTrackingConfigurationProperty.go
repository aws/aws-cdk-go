package awsgamelift


// *This data type is used with the Amazon GameLift FleetIQ and game server groups.*.
//
// Settings for a target-based scaling policy as part of a `GameServerGroupAutoScalingPolicy` . These settings are used to create a target-based policy that tracks the GameLift FleetIQ metric `"PercentUtilizedGameServers"` and specifies a target value for the metric. As player usage changes, the policy triggers to adjust the game server group capacity so that the metric returns to the target value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingConfigurationProperty := &targetTrackingConfigurationProperty{
//   	targetValue: jsii.Number(123),
//   }
//
type CfnGameServerGroup_TargetTrackingConfigurationProperty struct {
	// Desired value to use with a game server group target-based scaling policy.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
}


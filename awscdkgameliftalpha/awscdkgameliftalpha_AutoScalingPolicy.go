// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration settings for intelligent automatic scaling that uses target tracking.
//
// After the Auto Scaling group is created, all updates to Auto Scaling policies, including changing this policy and adding or removing other policies, is done directly on the Auto Scaling group.
//
// Example:
//   var launchTemplate iLaunchTemplate
//   var vpc iVpc
//
//
//   gamelift.NewGameServerGroup(this, jsii.String("Game server group"), &gameServerGroupProps{
//   	gameServerGroupName: jsii.String("sample-gameservergroup-name"),
//   	instanceDefinitions: []instanceDefinition{
//   		&instanceDefinition{
//   			instanceType: ec2.instanceType.of(ec2.instanceClass_C5, ec2.instanceSize_LARGE),
//   		},
//   		&instanceDefinition{
//   			instanceType: ec2.*instanceType.of(ec2.*instanceClass_C4, ec2.*instanceSize_LARGE),
//   		},
//   	},
//   	launchTemplate: launchTemplate,
//   	vpc: vpc,
//   	autoScalingPolicy: &autoScalingPolicy{
//   		estimatedInstanceWarmup: awscdk.Duration.minutes(jsii.Number(5)),
//   		targetTrackingConfiguration: jsii.Number(5),
//   	},
//   })
//
// Experimental.
type AutoScalingPolicy struct {
	// Settings for a target-based scaling policy applied to Auto Scaling group.
	//
	// These settings are used to create a target-based policy that tracks the GameLift FleetIQ metric `PercentUtilizedGameServers` and specifies a target value for the metric.
	//
	// As player usage changes, the policy triggers to adjust the game server group capacity so that the metric returns to the target value.
	// Experimental.
	TargetTrackingConfiguration *float64 `field:"required" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
	// Length of time, it takes for a new instance to start new game server processes and register with GameLift FleetIQ.
	//
	// Specifying a warm-up time can be useful, particularly with game servers that take a long time to start up, because it avoids prematurely starting new instances.
	// Experimental.
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
}


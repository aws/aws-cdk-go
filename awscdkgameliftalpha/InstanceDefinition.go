package awscdkgameliftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// An allowed instance type for a game server group.
//
// All game server groups must have at least two instance types defined for it.
// GameLift FleetIQ periodically evaluates each defined instance type for viability.
// It then updates the Auto Scaling group with the list of viable instance types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instanceType instanceType
//
//   instanceDefinition := &InstanceDefinition{
//   	InstanceType: instanceType,
//
//   	// the properties below are optional
//   	Weight: jsii.Number(123),
//   }
//
// Experimental.
type InstanceDefinition struct {
	// An Amazon EC2 instance type designation.
	// Experimental.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Instance weighting that indicates how much this instance type contributes to the total capacity of a game server group.
	//
	// Instance weights are used by GameLift FleetIQ to calculate the instance type's cost per unit hour and better identify the most cost-effective options.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-weighting.html
	//
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


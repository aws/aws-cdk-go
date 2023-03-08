package awsgamelift


// *This data type is used with the Amazon GameLift FleetIQ and game server groups.*.
//
// An allowed instance type for a `GameServerGroup` . All game server groups must have at least two instance types defined for it. GameLift FleetIQ periodically evaluates each defined instance type for viability. It then updates the Auto Scaling group with the list of viable instance types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceDefinitionProperty := &instanceDefinitionProperty{
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	weightedCapacity: jsii.String("weightedCapacity"),
//   }
//
type CfnGameServerGroup_InstanceDefinitionProperty struct {
	// An Amazon EC2 instance type designation.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Instance weighting that indicates how much this instance type contributes to the total capacity of a game server group.
	//
	// Instance weights are used by GameLift FleetIQ to calculate the instance type's cost per unit hour and better identify the most cost-effective options. For detailed information on weighting instance capacity, see [Instance Weighting](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-weighting.html) in the *Amazon Elastic Compute Cloud Auto Scaling User Guide* . Default value is "1".
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


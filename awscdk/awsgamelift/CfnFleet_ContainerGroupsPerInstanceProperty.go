package awsgamelift


// *This data type is currently not available.
//
// It is under improvement as we respond to customer feedback from the Containers public preview.*
//
// Determines how many replica container groups that Amazon GameLift deploys to each instance in a container fleet.
//
// Amazon GameLift calculates the maximum possible replica groups per instance based on the instance 's CPU and memory resources. When deploying a fleet, Amazon GameLift places replica container groups on each fleet instance based on the following:
//
// - If no desired value is set, Amazon GameLift places the calculated maximum.
// - If a desired number is set to a value higher than the calculated maximum, fleet creation fails..
// - If a desired number is set to a value lower than the calculated maximum, Amazon GameLift places the desired number.
//
// *Part of:* `ContainerGroupsConfiguration` , `ContainerGroupsAttributes`
//
// *Returned by:* `DescribeFleetAttributes` , `CreateFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerGroupsPerInstanceProperty := &ContainerGroupsPerInstanceProperty{
//   	DesiredReplicaContainerGroupsPerInstance: jsii.Number(123),
//   	MaxReplicaContainerGroupsPerInstance: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsperinstance.html
//
type CfnFleet_ContainerGroupsPerInstanceProperty struct {
	// The desired number of replica container groups to place on each fleet instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsperinstance.html#cfn-gamelift-fleet-containergroupsperinstance-desiredreplicacontainergroupsperinstance
	//
	DesiredReplicaContainerGroupsPerInstance *float64 `field:"optional" json:"desiredReplicaContainerGroupsPerInstance" yaml:"desiredReplicaContainerGroupsPerInstance"`
	// The maximum possible number of replica container groups that each fleet instance can have.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsperinstance.html#cfn-gamelift-fleet-containergroupsperinstance-maxreplicacontainergroupsperinstance
	//
	MaxReplicaContainerGroupsPerInstance *float64 `field:"optional" json:"maxReplicaContainerGroupsPerInstance" yaml:"maxReplicaContainerGroupsPerInstance"`
}


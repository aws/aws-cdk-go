package awsgamelift


// The number of container groups per instance.
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
	// Use this parameter to override the number of replica container groups GameLift will launch per instance with a number that is lower than that calculated maximum.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsperinstance.html#cfn-gamelift-fleet-containergroupsperinstance-desiredreplicacontainergroupsperinstance
	//
	DesiredReplicaContainerGroupsPerInstance *float64 `field:"optional" json:"desiredReplicaContainerGroupsPerInstance" yaml:"desiredReplicaContainerGroupsPerInstance"`
	// GameLift calculates the maximum number of replica container groups it can launch per instance based on instance properties such as CPU, memory, and connection ports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsperinstance.html#cfn-gamelift-fleet-containergroupsperinstance-maxreplicacontainergroupsperinstance
	//
	MaxReplicaContainerGroupsPerInstance *float64 `field:"optional" json:"maxReplicaContainerGroupsPerInstance" yaml:"maxReplicaContainerGroupsPerInstance"`
}


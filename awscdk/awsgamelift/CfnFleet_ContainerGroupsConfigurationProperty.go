package awsgamelift


// Specifies container groups that this instance will hold.
//
// You must specify exactly one replica group. Optionally, you may specify exactly one daemon group. You can't change this property after you create the fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerGroupsConfigurationProperty := &ContainerGroupsConfigurationProperty{
//   	ConnectionPortRange: &ConnectionPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	ContainerGroupDefinitionNames: []*string{
//   		jsii.String("containerGroupDefinitionNames"),
//   	},
//
//   	// the properties below are optional
//   	ContainerGroupsPerInstance: &ContainerGroupsPerInstanceProperty{
//   		DesiredReplicaContainerGroupsPerInstance: jsii.Number(123),
//   		MaxReplicaContainerGroupsPerInstance: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsconfiguration.html
//
type CfnFleet_ContainerGroupsConfigurationProperty struct {
	// Defines the range of ports on the instance that allow inbound traffic to connect with containers in a fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsconfiguration.html#cfn-gamelift-fleet-containergroupsconfiguration-connectionportrange
	//
	ConnectionPortRange interface{} `field:"required" json:"connectionPortRange" yaml:"connectionPortRange"`
	// The names of the container group definitions that will be created in an instance.
	//
	// You must specify exactly one REPLICA container group. You have the option to also specify one DAEMON container group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsconfiguration.html#cfn-gamelift-fleet-containergroupsconfiguration-containergroupdefinitionnames
	//
	ContainerGroupDefinitionNames *[]*string `field:"required" json:"containerGroupDefinitionNames" yaml:"containerGroupDefinitionNames"`
	// The number of container groups per instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsconfiguration.html#cfn-gamelift-fleet-containergroupsconfiguration-containergroupsperinstance
	//
	ContainerGroupsPerInstance interface{} `field:"optional" json:"containerGroupsPerInstance" yaml:"containerGroupsPerInstance"`
}


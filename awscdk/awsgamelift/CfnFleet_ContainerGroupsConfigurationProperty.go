package awsgamelift


// *This data type is currently not available.
//
// It is under improvement as we respond to customer feedback from the Containers public preview.*
//
// Configuration details for a set of container groups, for use when creating a fleet with compute type `CONTAINER` .
//
// *Used with:* `CreateFleet`.
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
	// A set of ports to allow inbound traffic, including game clients, to connect to processes running in the container fleet.
	//
	// Connection ports are dynamically mapped to container ports, which are assigned to individual processes running in a container. The connection port range must have enough ports to map to all container ports across a fleet instance. To calculate the minimum connection ports needed, use the following formula:
	//
	// *[Total number of container ports as defined for containers in the replica container group] * [Desired or calculated number of replica container groups per instance] + [Total number of container ports as defined for containers in the daemon container group]*
	//
	// As a best practice, double the minimum number of connection ports.
	//
	// > Use the fleet's `EC2InboundPermissions` property to control external access to connection ports. Set this property to the connection port numbers that you want to open access to. See `IpPermission` for more details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsconfiguration.html#cfn-gamelift-fleet-containergroupsconfiguration-connectionportrange
	//
	ConnectionPortRange interface{} `field:"required" json:"connectionPortRange" yaml:"connectionPortRange"`
	// The list of container group definition names to deploy to a new container fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsconfiguration.html#cfn-gamelift-fleet-containergroupsconfiguration-containergroupdefinitionnames
	//
	ContainerGroupDefinitionNames *[]*string `field:"required" json:"containerGroupDefinitionNames" yaml:"containerGroupDefinitionNames"`
	// The number of container groups per instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-containergroupsconfiguration.html#cfn-gamelift-fleet-containergroupsconfiguration-containergroupsperinstance
	//
	ContainerGroupsPerInstance interface{} `field:"optional" json:"containerGroupsPerInstance" yaml:"containerGroupsPerInstance"`
}


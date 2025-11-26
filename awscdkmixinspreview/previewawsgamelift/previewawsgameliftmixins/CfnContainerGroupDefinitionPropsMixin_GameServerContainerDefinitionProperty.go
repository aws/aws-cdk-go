package previewawsgameliftmixins


// Describes the game server container in an existing game server container group.
//
// A game server container identifies a container image with your game server build. A game server container is automatically considered essential; if an essential container fails, the entire container group restarts.
//
// You can update a container definition and deploy the updates to an existing fleet. When creating or updating a game server container group definition, use the property [](https://docs.aws.amazon.com/gamelift/latest/apireference/API_GameServerContainerDefinitionInput) .
//
// *Part of:* [ContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ContainerGroupDefinition.html)
//
// *Returned by:* [DescribeContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeContainerGroupDefinition.html) , [ListContainerGroupDefinitions](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListContainerGroupDefinitions.html) , [UpdateContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateContainerGroupDefinition.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gameServerContainerDefinitionProperty := &GameServerContainerDefinitionProperty{
//   	ContainerName: jsii.String("containerName"),
//   	DependsOn: []interface{}{
//   		&ContainerDependencyProperty{
//   			Condition: jsii.String("condition"),
//   			ContainerName: jsii.String("containerName"),
//   		},
//   	},
//   	EnvironmentOverride: []interface{}{
//   		&ContainerEnvironmentProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ImageUri: jsii.String("imageUri"),
//   	MountPoints: []interface{}{
//   		&ContainerMountPointProperty{
//   			AccessLevel: jsii.String("accessLevel"),
//   			ContainerPath: jsii.String("containerPath"),
//   			InstancePath: jsii.String("instancePath"),
//   		},
//   	},
//   	PortConfiguration: &PortConfigurationProperty{
//   		ContainerPortRanges: []interface{}{
//   			&ContainerPortRangeProperty{
//   				FromPort: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   				ToPort: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ResolvedImageDigest: jsii.String("resolvedImageDigest"),
//   	ServerSdkVersion: jsii.String("serverSdkVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html
//
type CfnContainerGroupDefinitionPropsMixin_GameServerContainerDefinitionProperty struct {
	// The container definition identifier.
	//
	// Container names are unique within a container group definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-containername
	//
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Indicates that the container relies on the status of other containers in the same container group during startup and shutdown sequences.
	//
	// A container might have dependencies on multiple containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-dependson
	//
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// A set of environment variables that's passed to the container on startup.
	//
	// See the [ContainerDefinition::environment](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html#ECS-Type-ContainerDefinition-environment) parameter in the *Amazon Elastic Container Service API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-environmentoverride
	//
	EnvironmentOverride interface{} `field:"optional" json:"environmentOverride" yaml:"environmentOverride"`
	// The URI to the image that Amazon GameLift Servers uses when deploying this container to a container fleet.
	//
	// For a more specific identifier, see `ResolvedImageDigest` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-imageuri
	//
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// A mount point that binds a path inside the container to a file or directory on the host system and lets it access the file or directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-mountpoints
	//
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// The set of ports that are available to bind to processes in the container.
	//
	// For example, a game server process requires a container port to allow game clients to connect to it. Container ports aren't directly accessed by inbound traffic. Amazon GameLift Servers maps these container ports to externally accessible connection ports, which are assigned as needed from the container fleet's `ConnectionPortRange` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-portconfiguration
	//
	PortConfiguration interface{} `field:"optional" json:"portConfiguration" yaml:"portConfiguration"`
	// A unique and immutable identifier for the container image.
	//
	// The digest is a SHA 256 hash of the container image manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-resolvedimagedigest
	//
	ResolvedImageDigest *string `field:"optional" json:"resolvedImageDigest" yaml:"resolvedImageDigest"`
	// The Amazon GameLift Servers server SDK version that the game server is integrated with.
	//
	// Only game servers using 5.2.0 or higher are compatible with container fleets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-gameservercontainerdefinition.html#cfn-gamelift-containergroupdefinition-gameservercontainerdefinition-serversdkversion
	//
	ServerSdkVersion *string `field:"optional" json:"serverSdkVersion" yaml:"serverSdkVersion"`
}


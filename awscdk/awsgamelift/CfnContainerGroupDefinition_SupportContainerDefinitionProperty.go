package awsgamelift


// Describes a support container in a container group.
//
// A support container might be in a game server container group or a per-instance container group. Support containers don't run game server processes.
//
// You can update a support container definition and deploy the updates to an existing fleet. When creating or updating a game server container group definition, use the property [GameServerContainerDefinitionInput](https://docs.aws.amazon.com/gamelift/latest/apireference/API_GameServerContainerDefinitionInput.html) .
//
// *Part of:* [ContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ContainerGroupDefinition.html)
//
// *Returned by:* [DescribeContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeContainerGroupDefinition.html) , [ListContainerGroupDefinitions](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListContainerGroupDefinitions.html) , [UpdateContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateContainerGroupDefinition.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   supportContainerDefinitionProperty := &SupportContainerDefinitionProperty{
//   	ContainerName: jsii.String("containerName"),
//   	ImageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
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
//   	Essential: jsii.Boolean(false),
//   	HealthCheck: &ContainerHealthCheckProperty{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//
//   		// the properties below are optional
//   		Interval: jsii.Number(123),
//   		Retries: jsii.Number(123),
//   		StartPeriod: jsii.Number(123),
//   		Timeout: jsii.Number(123),
//   	},
//   	MemoryHardLimitMebibytes: jsii.Number(123),
//   	MountPoints: []interface{}{
//   		&ContainerMountPointProperty{
//   			InstancePath: jsii.String("instancePath"),
//
//   			// the properties below are optional
//   			AccessLevel: jsii.String("accessLevel"),
//   			ContainerPath: jsii.String("containerPath"),
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
//   	Vcpu: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html
//
type CfnContainerGroupDefinition_SupportContainerDefinitionProperty struct {
	// The container definition identifier.
	//
	// Container names are unique within a container group definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-containername
	//
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The URI to the image that Amazon GameLift deploys to a container fleet.
	//
	// For a more specific identifier, see `ResolvedImageDigest` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-imageuri
	//
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Indicates that the container relies on the status of other containers in the same container group during its startup and shutdown sequences.
	//
	// A container might have dependencies on multiple containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-dependson
	//
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// A set of environment variables that's passed to the container on startup.
	//
	// See the [ContainerDefinition::environment](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html#ECS-Type-ContainerDefinition-environment) parameter in the *Amazon Elastic Container Service API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-environmentoverride
	//
	EnvironmentOverride interface{} `field:"optional" json:"environmentOverride" yaml:"environmentOverride"`
	// Indicates whether the container is vital to the container group.
	//
	// If an essential container fails, the entire container group restarts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-essential
	//
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// A configuration for a non-terminal health check.
	//
	// A support container automatically restarts if it stops functioning or if it fails this health check.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The amount of memory that Amazon GameLift makes available to the container.
	//
	// If memory limits aren't set for an individual container, the container shares the container group's total memory allocation.
	//
	// *Related data type:* [ContainerGroupDefinition TotalMemoryLimitMebibytes](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ContainerGroupDefinition.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-memoryhardlimitmebibytes
	//
	MemoryHardLimitMebibytes *float64 `field:"optional" json:"memoryHardLimitMebibytes" yaml:"memoryHardLimitMebibytes"`
	// A mount point that binds a path inside the container to a file or directory on the host system and lets it access the file or directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-mountpoints
	//
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// A set of ports that allow access to the container from external users.
	//
	// Processes running in the container can bind to a one of these ports. Container ports aren't directly accessed by inbound traffic. Amazon GameLift maps these container ports to externally accessible connection ports, which are assigned as needed from the container fleet's `ConnectionPortRange` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-portconfiguration
	//
	PortConfiguration interface{} `field:"optional" json:"portConfiguration" yaml:"portConfiguration"`
	// A unique and immutable identifier for the container image.
	//
	// The digest is a SHA 256 hash of the container image manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-resolvedimagedigest
	//
	ResolvedImageDigest *string `field:"optional" json:"resolvedImageDigest" yaml:"resolvedImageDigest"`
	// The number of vCPU units that are reserved for the container.
	//
	// If no resources are reserved, the container shares the total vCPU limit for the container group.
	//
	// *Related data type:* [ContainerGroupDefinition TotalVcpuLimit](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ContainerGroupDefinition.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-vcpu
	//
	Vcpu *float64 `field:"optional" json:"vcpu" yaml:"vcpu"`
}


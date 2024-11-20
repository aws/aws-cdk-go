package awsgamelift


// Supports the function of the main container group.
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
	// A descriptive label for the container definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-containername
	//
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Specifies the image URI of this container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-imageuri
	//
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// A list of container dependencies that determines when this container starts up and shuts down.
	//
	// For container groups with multiple containers, dependencies let you define a startup/shutdown sequence across the containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-dependson
	//
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// The environment variables to pass to a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-environmentoverride
	//
	EnvironmentOverride interface{} `field:"optional" json:"environmentOverride" yaml:"environmentOverride"`
	// Specifies if the container is essential.
	//
	// If an essential container fails a health check, then all containers in the container group will be restarted. You must specify exactly 1 essential container in a container group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-essential
	//
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// Specifies how the process manager checks the health of containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The total memory limit of container groups following this definition in MiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-memoryhardlimitmebibytes
	//
	MemoryHardLimitMebibytes *float64 `field:"optional" json:"memoryHardLimitMebibytes" yaml:"memoryHardLimitMebibytes"`
	// A list of mount point configurations to be used in a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-mountpoints
	//
	MountPoints interface{} `field:"optional" json:"mountPoints" yaml:"mountPoints"`
	// Defines the ports on a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-portconfiguration
	//
	PortConfiguration interface{} `field:"optional" json:"portConfiguration" yaml:"portConfiguration"`
	// The digest of the container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-resolvedimagedigest
	//
	ResolvedImageDigest *string `field:"optional" json:"resolvedImageDigest" yaml:"resolvedImageDigest"`
	// The number of virtual CPUs to give to the support group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-supportcontainerdefinition.html#cfn-gamelift-containergroupdefinition-supportcontainerdefinition-vcpu
	//
	Vcpu *float64 `field:"optional" json:"vcpu" yaml:"vcpu"`
}


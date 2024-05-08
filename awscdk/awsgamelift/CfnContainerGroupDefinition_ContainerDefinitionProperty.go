package awsgamelift


// *This data type is used with the Amazon GameLift containers feature, which is currently in public preview.*.
//
// Describes a container in a container fleet, the resources available to the container, and the commands that are run when the container starts. Container properties can't be updated. To change a property, create a new container group definition. See also `ContainerDefinitionInput` .
//
// *Part of:* `ContainerGroupDefinition`
//
// *Returned by:* `DescribeContainerGroupDefinition` , `ListContainerGroupDefinitions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerDefinitionProperty := &ContainerDefinitionProperty{
//   	ContainerName: jsii.String("containerName"),
//   	ImageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Cpu: jsii.Number(123),
//   	DependsOn: []interface{}{
//   		&ContainerDependencyProperty{
//   			Condition: jsii.String("condition"),
//   			ContainerName: jsii.String("containerName"),
//   		},
//   	},
//   	EntryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	Environment: []interface{}{
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
//   	MemoryLimits: &MemoryLimitsProperty{
//   		HardLimit: jsii.Number(123),
//   		SoftLimit: jsii.Number(123),
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
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html
//
type CfnContainerGroupDefinition_ContainerDefinitionProperty struct {
	// The container definition identifier.
	//
	// Container names are unique within a container group definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-containername
	//
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The URI to the image that $short;
	//
	// copied and deployed to a container fleet. For a more specific identifier, see `ResolvedImageDigest` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-imageuri
	//
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// A command that's passed to the container on startup.
	//
	// Each argument for the command is an additional string in the array. See the [ContainerDefinition::command](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html#ECS-Type-ContainerDefinition-command) parameter in the *Amazon Elastic Container Service API reference.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of CPU units that are reserved for the container.
	//
	// Note: 1 vCPU unit equals 1024 CPU units. If no resources are reserved, the container shares the total CPU limit for the container group.
	//
	// *Related data type:* `ContainerGroupDefinition$TotalCpuLimit`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-cpu
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Indicates that the container relies on the status of other containers in the same container group during its startup and shutdown sequences.
	//
	// A container might have dependencies on multiple containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-dependson
	//
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// The entry point that's passed to the container on startup.
	//
	// If there are multiple arguments, each argument is an additional string in the array. See the [ContainerDefinition::entryPoint](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html#ECS-Type-ContainerDefinition-entryPoint) parameter in the *Amazon Elastic Container Service API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-entrypoint
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// A set of environment variables that's passed to the container on startup.
	//
	// See the [ContainerDefinition::environment](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html#ECS-Type-ContainerDefinition-environment) parameter in the *Amazon Elastic Container Service API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Indicates whether the container is vital to the container group.
	//
	// If an essential container fails, the entire container group is restarted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-essential
	//
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// A configuration for a non-terminal health check.
	//
	// A container, which automatically restarts if it stops functioning, also restarts if it fails this health check. If an essential container in the daemon group fails a health check, the entire container group is restarted. The essential container in the replica group doesn't use this health check mechanism, because the Amazon GameLift Agent automatically handles the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The amount of memory that Amazon GameLift makes available to the container.
	//
	// If memory limits aren't set for an individual container, the container shares the container group's total memory allocation.
	//
	// *Related data type:* `ContainerGroupDefinition$TotalMemoryLimit`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-memorylimits
	//
	MemoryLimits interface{} `field:"optional" json:"memoryLimits" yaml:"memoryLimits"`
	// Defines the ports that are available to assign to processes in the container.
	//
	// For example, a game server process requires a container port to allow game clients to connect to it. Container ports aren't directly accessed by inbound traffic. Amazon GameLift maps these container ports to externally accessible connection ports, which are assigned as needed from the container fleet's `ConnectionPortRange` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-portconfiguration
	//
	PortConfiguration interface{} `field:"optional" json:"portConfiguration" yaml:"portConfiguration"`
	// A unique and immutable identifier for the container image that is deployed to a container fleet.
	//
	// The digest is a SHA 256 hash of the container image manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-resolvedimagedigest
	//
	ResolvedImageDigest *string `field:"optional" json:"resolvedImageDigest" yaml:"resolvedImageDigest"`
	// The directory in the container where commands are run.
	//
	// See the [ContainerDefinition::workingDirectory](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html#ECS-Type-ContainerDefinition-workingDirectory) parameter in the *Amazon Elastic Container Service API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-workingdirectory
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}


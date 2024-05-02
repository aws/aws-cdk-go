package awsgamelift


// Details about a container that is used in a container fleet.
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
	// A descriptive label for the container definition.
	//
	// Container definition names must be unique with a container group definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-containername
	//
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Specifies the image URI of this container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-imageuri
	//
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The command that's passed to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The maximum number of CPU units reserved for this container.
	//
	// The value is expressed as an integer amount of CPU units. 1 vCPU is equal to 1024 CPU units
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-cpu
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// A list of container dependencies that determines when this container starts up and shuts down.
	//
	// For container groups with multiple containers, dependencies let you define a startup/shutdown sequence across the containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-dependson
	//
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// The entry point that's passed to the container so that it will run as an executable.
	//
	// If there are multiple arguments, each argument is a string in the array.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-entrypoint
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Specifies if the container is essential.
	//
	// If an essential container fails a health check, then all containers in the container group will be restarted. You must specify exactly 1 essential container in a container group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-essential
	//
	Essential interface{} `field:"optional" json:"essential" yaml:"essential"`
	// Specifies how the process manager checks the health of containers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-healthcheck
	//
	HealthCheck interface{} `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Specifies how much memory is available to the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-memorylimits
	//
	MemoryLimits interface{} `field:"optional" json:"memoryLimits" yaml:"memoryLimits"`
	// Defines the ports on a container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-portconfiguration
	//
	PortConfiguration interface{} `field:"optional" json:"portConfiguration" yaml:"portConfiguration"`
	// The digest of the container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-resolvedimagedigest
	//
	ResolvedImageDigest *string `field:"optional" json:"resolvedImageDigest" yaml:"resolvedImageDigest"`
	// The working directory to run commands inside the container in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-containerdefinition.html#cfn-gamelift-containergroupdefinition-containerdefinition-workingdirectory
	//
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}


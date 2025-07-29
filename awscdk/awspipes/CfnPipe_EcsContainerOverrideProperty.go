package awspipes


// The overrides that are sent to a container.
//
// An empty container override can be passed in. An example of an empty container override is `{"containerOverrides": [ ] }` . If a non-empty container override is specified, the `name` parameter must be included.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsContainerOverrideProperty := &EcsContainerOverrideProperty{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Cpu: jsii.Number(123),
//   	Environment: []interface{}{
//   		&EcsEnvironmentVariableProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	EnvironmentFiles: []interface{}{
//   		&EcsEnvironmentFileProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Memory: jsii.Number(123),
//   	MemoryReservation: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	ResourceRequirements: []interface{}{
//   		&EcsResourceRequirementProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html
//
type CfnPipe_EcsContainerOverrideProperty struct {
	// The command to send to the container that overrides the default command from the Docker image or the task definition.
	//
	// You must also specify a container name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-command
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of `cpu` units reserved for the container, instead of the default value from the task definition.
	//
	// You must also specify a container name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-cpu
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The environment variables to send to the container.
	//
	// You can add new environment variables, which are added to the container at launch, or you can override the existing environment variables from the Docker image or the task definition. You must also specify a container name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// A list of files containing the environment variables to pass to a container, instead of the value from the container definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-environmentfiles
	//
	EnvironmentFiles interface{} `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// The hard limit (in MiB) of memory to present to the container, instead of the default value from the task definition.
	//
	// If your container attempts to exceed the memory specified here, the container is killed. You must also specify a container name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-memory
	//
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The soft limit (in MiB) of memory to reserve for the container, instead of the default value from the task definition.
	//
	// You must also specify a container name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-memoryreservation
	//
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// The name of the container that receives the override.
	//
	// This parameter is required if any override is specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type and amount of a resource to assign to a container, instead of the default value from the task definition.
	//
	// The only supported resource is a GPU.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-resourcerequirements
	//
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
}


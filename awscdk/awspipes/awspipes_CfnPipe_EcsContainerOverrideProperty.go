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
//   ecsContainerOverrideProperty := &ecsContainerOverrideProperty{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	cpu: jsii.Number(123),
//   	environment: []interface{}{
//   		&ecsEnvironmentVariableProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	environmentFiles: []interface{}{
//   		&ecsEnvironmentFileProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	memory: jsii.Number(123),
//   	memoryReservation: jsii.Number(123),
//   	name: jsii.String("name"),
//   	resourceRequirements: []interface{}{
//   		&ecsResourceRequirementProperty{
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipe_EcsContainerOverrideProperty struct {
	// The command to send to the container that overrides the default command from the Docker image or the task definition.
	//
	// You must also specify a container name.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of `cpu` units reserved for the container, instead of the default value from the task definition.
	//
	// You must also specify a container name.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The environment variables to send to the container.
	//
	// You can add new environment variables, which are added to the container at launch, or you can override the existing environment variables from the Docker image or the task definition. You must also specify a container name.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// A list of files containing the environment variables to pass to a container, instead of the value from the container definition.
	EnvironmentFiles interface{} `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// The hard limit (in MiB) of memory to present to the container, instead of the default value from the task definition.
	//
	// If your container attempts to exceed the memory specified here, the container is killed. You must also specify a container name.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The soft limit (in MiB) of memory to reserve for the container, instead of the default value from the task definition.
	//
	// You must also specify a container name.
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
	// The name of the container that receives the override.
	//
	// This parameter is required if any override is specified.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type and amount of a resource to assign to a container, instead of the default value from the task definition.
	//
	// The only supported resource is a GPU.
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
}


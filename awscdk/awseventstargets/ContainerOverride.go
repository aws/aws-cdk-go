package awseventstargets


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerOverride := &ContainerOverride{
//   	ContainerName: jsii.String("containerName"),
//
//   	// the properties below are optional
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Cpu: jsii.Number(123),
//   	Environment: []taskEnvironmentVariable{
//   		&taskEnvironmentVariable{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MemoryLimit: jsii.Number(123),
//   	MemoryReservation: jsii.Number(123),
//   }
//
type ContainerOverride struct {
	// Name of the container inside the task definition.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Command to run inside the container.
	// Default: Default command.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of cpu units reserved for the container.
	// Default: The default value from the task definition.
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Variables to set in the container's environment.
	Environment *[]*TaskEnvironmentVariable `field:"optional" json:"environment" yaml:"environment"`
	// Hard memory limit on the container.
	// Default: The default value from the task definition.
	//
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// Soft memory limit on the container.
	// Default: The default value from the task definition.
	//
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
}


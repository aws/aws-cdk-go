package awseventstargets


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerOverride := &containerOverride{
//   	containerName: jsii.String("containerName"),
//
//   	// the properties below are optional
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	cpu: jsii.Number(123),
//   	environment: []taskEnvironmentVariable{
//   		&taskEnvironmentVariable{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	memoryLimit: jsii.Number(123),
//   	memoryReservation: jsii.Number(123),
//   }
//
type ContainerOverride struct {
	// Name of the container inside the task definition.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Command to run inside the container.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of cpu units reserved for the container.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Variables to set in the container's environment.
	Environment *[]*TaskEnvironmentVariable `field:"optional" json:"environment" yaml:"environment"`
	// Hard memory limit on the container.
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// Soft memory limit on the container.
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
}


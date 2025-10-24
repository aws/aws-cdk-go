package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// A list of container overrides that specify the name of a container and the overrides it should receive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerDefinition ContainerDefinition
//
//   containerOverride := &ContainerOverride{
//   	ContainerDefinition: containerDefinition,
//
//   	// the properties below are optional
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Cpu: jsii.Number(123),
//   	Environment: []TaskEnvironmentVariable{
//   		&TaskEnvironmentVariable{
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
	ContainerDefinition awsecs.ContainerDefinition `field:"required" json:"containerDefinition" yaml:"containerDefinition"`
	// Command to run inside the container.
	// Default: - Default command from the Docker image or the task definition.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of cpu units reserved for the container.
	// Default: - The default value from the task definition.
	//
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The environment variables to send to the container.
	//
	// You can add new environment variables, which are added to the container at launch,
	// or you can override the existing environment variables from the Docker image or the task definition.
	// Default: - The existing environment variables from the Docker image or the task definition.
	//
	Environment *[]*TaskEnvironmentVariable `field:"optional" json:"environment" yaml:"environment"`
	// The hard limit (in MiB) of memory to present to the container.
	// Default: - The default value from the task definition.
	//
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// The soft limit (in MiB) of memory to reserve for the container.
	// Default: - The default value from the task definition.
	//
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
}


package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
)

// A list of container overrides that specify the name of a container and the overrides it should receive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var containerDefinition containerDefinition
//
//   containerOverride := &containerOverride{
//   	containerDefinition: containerDefinition,
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
// Experimental.
type ContainerOverride struct {
	// Name of the container inside the task definition.
	// Experimental.
	ContainerDefinition awsecs.ContainerDefinition `field:"required" json:"containerDefinition" yaml:"containerDefinition"`
	// Command to run inside the container.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of cpu units reserved for the container.
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The environment variables to send to the container.
	//
	// You can add new environment variables, which are added to the container at launch,
	// or you can override the existing environment variables from the Docker image or the task definition.
	// Experimental.
	Environment *[]*TaskEnvironmentVariable `field:"optional" json:"environment" yaml:"environment"`
	// The hard limit (in MiB) of memory to present to the container.
	// Experimental.
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// The soft limit (in MiB) of memory to reserve for the container.
	// Experimental.
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
}


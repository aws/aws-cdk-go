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
//   var containerDefinition containerDefinition
//
//   containerOverride := &ContainerOverride{
//   	ContainerDefinition: containerDefinition,
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
	ContainerDefinition awsecs.ContainerDefinition `field:"required" json:"containerDefinition" yaml:"containerDefinition"`
	// Command to run inside the container.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The number of cpu units reserved for the container.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The environment variables to send to the container.
	//
	// You can add new environment variables, which are added to the container at launch,
	// or you can override the existing environment variables from the Docker image or the task definition.
	Environment *[]*TaskEnvironmentVariable `field:"optional" json:"environment" yaml:"environment"`
	// The hard limit (in MiB) of memory to present to the container.
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// The soft limit (in MiB) of memory to reserve for the container.
	MemoryReservation *float64 `field:"optional" json:"memoryReservation" yaml:"memoryReservation"`
}


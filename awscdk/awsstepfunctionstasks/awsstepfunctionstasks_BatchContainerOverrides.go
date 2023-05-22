package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// The overrides that should be sent to a container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instanceType instanceType
//   var size size
//
//   batchContainerOverrides := &BatchContainerOverrides{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	GpuCount: jsii.Number(123),
//   	InstanceType: instanceType,
//   	Memory: size,
//   	Vcpus: jsii.Number(123),
//   }
//
// Experimental.
type BatchContainerOverrides struct {
	// The command to send to the container that overrides the default command from the Docker image or the job definition.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The environment variables to send to the container.
	//
	// You can add new environment variables, which are added to the container
	// at launch, or you can override the existing environment variables from
	// the Docker image or the job definition.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The number of physical GPUs to reserve for the container.
	//
	// The number of GPUs reserved for all containers in a job
	// should not exceed the number of available GPUs on the compute
	// resource that the job is launched on.
	// Experimental.
	GpuCount *float64 `field:"optional" json:"gpuCount" yaml:"gpuCount"`
	// The instance type to use for a multi-node parallel job.
	//
	// This parameter is not valid for single-node container jobs.
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Memory reserved for the job.
	// Experimental.
	Memory awscdk.Size `field:"optional" json:"memory" yaml:"memory"`
	// The number of vCPUs to reserve for the container.
	//
	// This value overrides the value set in the job definition.
	// Experimental.
	Vcpus *float64 `field:"optional" json:"vcpus" yaml:"vcpus"`
}


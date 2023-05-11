// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A container orchestrated by ECS that uses EC2 resources.
// Experimental.
type IEcsEc2ContainerDefinition interface {
	IEcsContainerDefinition
	// Add a ulimit to this container.
	// Experimental.
	AddUlimit(ulimit *Ulimit)
	// The number of physical GPUs to reserve for the container.
	//
	// Make sure that the number of GPUs reserved for all containers in a job doesn't exceed
	// the number of available GPUs on the compute resource that the job is launched on.
	// Experimental.
	Gpu() *float64
	// When this parameter is true, the container is given elevated permissions on the host container instance (similar to the root user).
	// Experimental.
	Privileged() *bool
	// Limits to set for the user this docker container will run as.
	// Experimental.
	Ulimits() *[]*Ulimit
}

// The jsii proxy for IEcsEc2ContainerDefinition
type jsiiProxy_IEcsEc2ContainerDefinition struct {
	jsiiProxy_IEcsContainerDefinition
}

func (i *jsiiProxy_IEcsEc2ContainerDefinition) AddUlimit(ulimit *Ulimit) {
	if err := i.validateAddUlimitParameters(ulimit); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addUlimit",
		[]interface{}{ulimit},
	)
}

func (j *jsiiProxy_IEcsEc2ContainerDefinition) Gpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsEc2ContainerDefinition) Privileged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEcsEc2ContainerDefinition) Ulimits() *[]*Ulimit {
	var returns *[]*Ulimit
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}


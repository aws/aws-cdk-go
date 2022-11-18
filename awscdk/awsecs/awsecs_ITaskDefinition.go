package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// The interface for all task definitions.
// Experimental.
type ITaskDefinition interface {
	awscdk.IResource
	// What launch types this task definition should be compatible with.
	// Experimental.
	Compatibility() Compatibility
	// Execution role for this task definition.
	// Experimental.
	ExecutionRole() awsiam.IRole
	// Return true if the task definition can be run on an EC2 cluster.
	// Experimental.
	IsEc2Compatible() *bool
	// Return true if the task definition can be run on a ECS Anywhere cluster.
	// Experimental.
	IsExternalCompatible() *bool
	// Return true if the task definition can be run on a Fargate cluster.
	// Experimental.
	IsFargateCompatible() *bool
	// The networking mode to use for the containers in the task.
	// Experimental.
	NetworkMode() NetworkMode
	// ARN of this task definition.
	// Experimental.
	TaskDefinitionArn() *string
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Experimental.
	TaskRole() awsiam.IRole
}

// The jsii proxy for ITaskDefinition
type jsiiProxy_ITaskDefinition struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITaskDefinition) Compatibility() Compatibility {
	var returns Compatibility
	_jsii_.Get(
		j,
		"compatibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) IsEc2Compatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isEc2Compatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) IsExternalCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isExternalCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) IsFargateCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFargateCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) NetworkMode() NetworkMode {
	var returns NetworkMode
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinition) TaskRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"taskRole",
		&returns,
	)
	return returns
}


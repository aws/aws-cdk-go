package awscdkpipesalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/internal"
)

// Interface representing a created or an imported `Pipe`.
// Experimental.
type IPipe interface {
	awscdk.IResource
	// The ARN of the pipe.
	// Experimental.
	PipeArn() *string
	// The name of the pipe.
	// Experimental.
	PipeName() *string
	// The role used by the pipe.
	//
	// For imported pipes it assumes that the default role is used.
	// Experimental.
	PipeRole() awsiam.IRole
}

// The jsii proxy for IPipe
type jsiiProxy_IPipe struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IPipe) PipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipe) PipeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipe) PipeRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"pipeRole",
		&returns,
	)
	return returns
}


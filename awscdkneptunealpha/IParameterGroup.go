package awscdkneptunealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkneptunealpha/v2/internal"
)

// A parameter group.
// Experimental.
type IParameterGroup interface {
	awscdk.IResource
	// The name of this parameter group.
	// Experimental.
	ParameterGroupName() *string
}

// The jsii proxy for IParameterGroup
type jsiiProxy_IParameterGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IParameterGroup) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}


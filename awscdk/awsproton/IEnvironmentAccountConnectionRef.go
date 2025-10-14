package awsproton

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsproton/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentAccountConnection.
// Experimental.
type IEnvironmentAccountConnectionRef interface {
	constructs.IConstruct
	// A reference to a EnvironmentAccountConnection resource.
	// Experimental.
	EnvironmentAccountConnectionRef() *EnvironmentAccountConnectionReference
}

// The jsii proxy for IEnvironmentAccountConnectionRef
type jsiiProxy_IEnvironmentAccountConnectionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEnvironmentAccountConnectionRef) EnvironmentAccountConnectionRef() *EnvironmentAccountConnectionReference {
	var returns *EnvironmentAccountConnectionReference
	_jsii_.Get(
		j,
		"environmentAccountConnectionRef",
		&returns,
	)
	return returns
}


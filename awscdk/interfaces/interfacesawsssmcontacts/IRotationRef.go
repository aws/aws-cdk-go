package interfacesawsssmcontacts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsssmcontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Rotation.
// Experimental.
type IRotationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Rotation resource.
	// Experimental.
	RotationRef() *RotationReference
}

// The jsii proxy for IRotationRef
type jsiiProxy_IRotationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRotationRef) RotationRef() *RotationReference {
	var returns *RotationReference
	_jsii_.Get(
		j,
		"rotationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRotationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRotationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


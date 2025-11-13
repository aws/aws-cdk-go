package interfacesawsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SdiSource.
// Experimental.
type ISdiSourceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SdiSource resource.
	// Experimental.
	SdiSourceRef() *SdiSourceReference
}

// The jsii proxy for ISdiSourceRef
type jsiiProxy_ISdiSourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISdiSourceRef) SdiSourceRef() *SdiSourceReference {
	var returns *SdiSourceReference
	_jsii_.Get(
		j,
		"sdiSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISdiSourceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISdiSourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Standard.
// Experimental.
type IStandardRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Standard resource.
	// Experimental.
	StandardRef() *StandardReference
}

// The jsii proxy for IStandardRef
type jsiiProxy_IStandardRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStandardRef) StandardRef() *StandardReference {
	var returns *StandardReference
	_jsii_.Get(
		j,
		"standardRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStandardRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStandardRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


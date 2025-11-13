package interfacesawsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Studio.
// Experimental.
type IStudioRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Studio resource.
	// Experimental.
	StudioRef() *StudioReference
}

// The jsii proxy for IStudioRef
type jsiiProxy_IStudioRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStudioRef) StudioRef() *StudioReference {
	var returns *StudioReference
	_jsii_.Get(
		j,
		"studioRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStudioRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStudioRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Space.
// Experimental.
type ISpaceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Space resource.
	// Experimental.
	SpaceRef() *SpaceReference
}

// The jsii proxy for ISpaceRef
type jsiiProxy_ISpaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISpaceRef) SpaceRef() *SpaceReference {
	var returns *SpaceReference
	_jsii_.Get(
		j,
		"spaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpaceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


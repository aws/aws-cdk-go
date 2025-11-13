package interfacesawscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TypeActivation.
// Experimental.
type ITypeActivationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TypeActivation resource.
	// Experimental.
	TypeActivationRef() *TypeActivationReference
}

// The jsii proxy for ITypeActivationRef
type jsiiProxy_ITypeActivationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITypeActivationRef) TypeActivationRef() *TypeActivationReference {
	var returns *TypeActivationReference
	_jsii_.Get(
		j,
		"typeActivationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITypeActivationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITypeActivationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


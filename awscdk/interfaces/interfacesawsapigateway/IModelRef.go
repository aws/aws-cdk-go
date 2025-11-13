package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Model.
// Experimental.
type IModelRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Model resource.
	// Experimental.
	ModelRef() *ModelReference
}

// The jsii proxy for IModelRef
type jsiiProxy_IModelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IModelRef) ModelRef() *ModelReference {
	var returns *ModelReference
	_jsii_.Get(
		j,
		"modelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


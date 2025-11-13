package interfacesawsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FormType.
// Experimental.
type IFormTypeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FormType resource.
	// Experimental.
	FormTypeRef() *FormTypeReference
}

// The jsii proxy for IFormTypeRef
type jsiiProxy_IFormTypeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IFormTypeRef) FormTypeRef() *FormTypeReference {
	var returns *FormTypeReference
	_jsii_.Get(
		j,
		"formTypeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormTypeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormTypeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PredefinedAttribute.
// Experimental.
type IPredefinedAttributeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PredefinedAttribute resource.
	// Experimental.
	PredefinedAttributeRef() *PredefinedAttributeReference
}

// The jsii proxy for IPredefinedAttributeRef
type jsiiProxy_IPredefinedAttributeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPredefinedAttributeRef) PredefinedAttributeRef() *PredefinedAttributeReference {
	var returns *PredefinedAttributeReference
	_jsii_.Get(
		j,
		"predefinedAttributeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPredefinedAttributeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPredefinedAttributeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


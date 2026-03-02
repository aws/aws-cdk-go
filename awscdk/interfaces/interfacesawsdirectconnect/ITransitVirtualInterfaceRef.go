package interfacesawsdirectconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdirectconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitVirtualInterface.
// Experimental.
type ITransitVirtualInterfaceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitVirtualInterface resource.
	// Experimental.
	TransitVirtualInterfaceRef() *TransitVirtualInterfaceReference
}

// The jsii proxy for ITransitVirtualInterfaceRef
type jsiiProxy_ITransitVirtualInterfaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITransitVirtualInterfaceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITransitVirtualInterfaceRef) TransitVirtualInterfaceRef() *TransitVirtualInterfaceReference {
	var returns *TransitVirtualInterfaceReference
	_jsii_.Get(
		j,
		"transitVirtualInterfaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitVirtualInterfaceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitVirtualInterfaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


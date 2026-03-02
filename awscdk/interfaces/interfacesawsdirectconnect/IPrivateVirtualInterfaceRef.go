package interfacesawsdirectconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdirectconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrivateVirtualInterface.
// Experimental.
type IPrivateVirtualInterfaceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PrivateVirtualInterface resource.
	// Experimental.
	PrivateVirtualInterfaceRef() *PrivateVirtualInterfaceReference
}

// The jsii proxy for IPrivateVirtualInterfaceRef
type jsiiProxy_IPrivateVirtualInterfaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPrivateVirtualInterfaceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPrivateVirtualInterfaceRef) PrivateVirtualInterfaceRef() *PrivateVirtualInterfaceReference {
	var returns *PrivateVirtualInterfaceReference
	_jsii_.Get(
		j,
		"privateVirtualInterfaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrivateVirtualInterfaceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrivateVirtualInterfaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


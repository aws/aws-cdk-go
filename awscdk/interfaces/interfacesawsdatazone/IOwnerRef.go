package interfacesawsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Owner.
// Experimental.
type IOwnerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Owner resource.
	// Experimental.
	OwnerRef() *OwnerReference
}

// The jsii proxy for IOwnerRef
type jsiiProxy_IOwnerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOwnerRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOwnerRef) OwnerRef() *OwnerReference {
	var returns *OwnerReference
	_jsii_.Get(
		j,
		"ownerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOwnerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOwnerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPv4Pool.
// Experimental.
type IIPv4PoolRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IPv4Pool resource.
	// Experimental.
	IPv4PoolRef() *IPv4PoolReference
}

// The jsii proxy for IIPv4PoolRef
type jsiiProxy_IIPv4PoolRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIPv4PoolRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIPv4PoolRef) IPv4PoolRef() *IPv4PoolReference {
	var returns *IPv4PoolReference
	_jsii_.Get(
		j,
		"iPv4PoolRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIPv4PoolRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIPv4PoolRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverConfig.
// Experimental.
type IResolverConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ResolverConfig resource.
	// Experimental.
	ResolverConfigRef() *ResolverConfigReference
}

// The jsii proxy for IResolverConfigRef
type jsiiProxy_IResolverConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IResolverConfigRef) ResolverConfigRef() *ResolverConfigReference {
	var returns *ResolverConfigReference
	_jsii_.Get(
		j,
		"resolverConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolverConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolverConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverEndpoint.
// Experimental.
type IResolverEndpointRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResolverEndpoint resource.
	// Experimental.
	ResolverEndpointRef() *ResolverEndpointReference
}

// The jsii proxy for IResolverEndpointRef
type jsiiProxy_IResolverEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResolverEndpointRef) ResolverEndpointRef() *ResolverEndpointReference {
	var returns *ResolverEndpointReference
	_jsii_.Get(
		j,
		"resolverEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolverEndpointRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResolverEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


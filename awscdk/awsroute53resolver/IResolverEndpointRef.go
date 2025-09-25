package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverEndpoint.
// Experimental.
type IResolverEndpointRef interface {
	constructs.IConstruct
	// A reference to a ResolverEndpoint resource.
	// Experimental.
	ResolverEndpointRef() *ResolverEndpointReference
}

// The jsii proxy for IResolverEndpointRef
type jsiiProxy_IResolverEndpointRef struct {
	internal.Type__constructsIConstruct
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


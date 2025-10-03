package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverEndpoint.
// Experimental.
type IResolverEndpointRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResolverEndpointRef
type jsiiProxy_IResolverEndpointRef struct {
	internal.Type__constructsIConstruct
}


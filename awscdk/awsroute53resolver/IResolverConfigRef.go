package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverConfig.
// Experimental.
type IResolverConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResolverConfigRef
type jsiiProxy_IResolverConfigRef struct {
	internal.Type__constructsIConstruct
}


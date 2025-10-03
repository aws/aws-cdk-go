package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Resolver.
// Experimental.
type IResolverRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResolverRef
type jsiiProxy_IResolverRef struct {
	internal.Type__constructsIConstruct
}


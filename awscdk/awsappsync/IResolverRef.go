package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Resolver.
// Experimental.
type IResolverRef interface {
	constructs.IConstruct
	// A reference to a Resolver resource.
	// Experimental.
	ResolverRef() *ResolverReference
}

// The jsii proxy for IResolverRef
type jsiiProxy_IResolverRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResolverRef) ResolverRef() *ResolverReference {
	var returns *ResolverReference
	_jsii_.Get(
		j,
		"resolverRef",
		&returns,
	)
	return returns
}


package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResolverConfig.
// Experimental.
type IResolverConfigRef interface {
	constructs.IConstruct
	// A reference to a ResolverConfig resource.
	// Experimental.
	ResolverConfigRef() *ResolverConfigReference
}

// The jsii proxy for IResolverConfigRef
type jsiiProxy_IResolverConfigRef struct {
	internal.Type__constructsIConstruct
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


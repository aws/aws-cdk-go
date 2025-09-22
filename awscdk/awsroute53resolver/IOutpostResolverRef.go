package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OutpostResolver.
// Experimental.
type IOutpostResolverRef interface {
	constructs.IConstruct
	// A reference to a OutpostResolver resource.
	// Experimental.
	OutpostResolverRef() *OutpostResolverReference
}

// The jsii proxy for IOutpostResolverRef
type jsiiProxy_IOutpostResolverRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IOutpostResolverRef) OutpostResolverRef() *OutpostResolverReference {
	var returns *OutpostResolverReference
	_jsii_.Get(
		j,
		"outpostResolverRef",
		&returns,
	)
	return returns
}


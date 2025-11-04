package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OutpostResolver.
// Experimental.
type IOutpostResolverRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a OutpostResolver resource.
	// Experimental.
	OutpostResolverRef() *OutpostResolverReference
}

// The jsii proxy for IOutpostResolverRef
type jsiiProxy_IOutpostResolverRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IOutpostResolverRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOutpostResolverRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


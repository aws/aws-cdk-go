package awsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsathena/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NamedQuery.
// Experimental.
type INamedQueryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a NamedQuery resource.
	// Experimental.
	NamedQueryRef() *NamedQueryReference
}

// The jsii proxy for INamedQueryRef
type jsiiProxy_INamedQueryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_INamedQueryRef) NamedQueryRef() *NamedQueryReference {
	var returns *NamedQueryReference
	_jsii_.Get(
		j,
		"namedQueryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamedQueryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamedQueryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


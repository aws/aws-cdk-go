package awsresourceexplorer2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsresourceexplorer2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Index.
// Experimental.
type IIndexRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Index resource.
	// Experimental.
	IndexRef() *IndexReference
}

// The jsii proxy for IIndexRef
type jsiiProxy_IIndexRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IIndexRef) IndexRef() *IndexReference {
	var returns *IndexReference
	_jsii_.Get(
		j,
		"indexRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIndexRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIndexRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


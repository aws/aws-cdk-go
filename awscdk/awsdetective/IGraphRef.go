package awsdetective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdetective/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Graph.
// Experimental.
type IGraphRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Graph resource.
	// Experimental.
	GraphRef() *GraphReference
}

// The jsii proxy for IGraphRef
type jsiiProxy_IGraphRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IGraphRef) GraphRef() *GraphReference {
	var returns *GraphReference
	_jsii_.Get(
		j,
		"graphRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


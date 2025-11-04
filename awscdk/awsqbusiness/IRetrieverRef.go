package awsqbusiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Retriever.
// Experimental.
type IRetrieverRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Retriever resource.
	// Experimental.
	RetrieverRef() *RetrieverReference
}

// The jsii proxy for IRetrieverRef
type jsiiProxy_IRetrieverRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRetrieverRef) RetrieverRef() *RetrieverReference {
	var returns *RetrieverReference
	_jsii_.Get(
		j,
		"retrieverRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRetrieverRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRetrieverRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


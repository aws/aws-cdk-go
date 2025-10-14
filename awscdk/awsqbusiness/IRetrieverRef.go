package awsqbusiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Retriever.
// Experimental.
type IRetrieverRef interface {
	constructs.IConstruct
	// A reference to a Retriever resource.
	// Experimental.
	RetrieverRef() *RetrieverReference
}

// The jsii proxy for IRetrieverRef
type jsiiProxy_IRetrieverRef struct {
	internal.Type__constructsIConstruct
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


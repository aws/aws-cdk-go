package awsgameliftstreams

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgameliftstreams/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamGroup.
// Experimental.
type IStreamGroupRef interface {
	constructs.IConstruct
	// A reference to a StreamGroup resource.
	// Experimental.
	StreamGroupRef() *StreamGroupReference
}

// The jsii proxy for IStreamGroupRef
type jsiiProxy_IStreamGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStreamGroupRef) StreamGroupRef() *StreamGroupReference {
	var returns *StreamGroupReference
	_jsii_.Get(
		j,
		"streamGroupRef",
		&returns,
	)
	return returns
}


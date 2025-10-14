package awskinesis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamConsumer.
// Experimental.
type IStreamConsumerRef interface {
	constructs.IConstruct
	// A reference to a StreamConsumer resource.
	// Experimental.
	StreamConsumerRef() *StreamConsumerReference
}

// The jsii proxy for IStreamConsumerRef
type jsiiProxy_IStreamConsumerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStreamConsumerRef) StreamConsumerRef() *StreamConsumerReference {
	var returns *StreamConsumerReference
	_jsii_.Get(
		j,
		"streamConsumerRef",
		&returns,
	)
	return returns
}


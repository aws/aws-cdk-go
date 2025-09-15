package awsrekognition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrekognition/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamProcessor.
// Experimental.
type IStreamProcessorRef interface {
	constructs.IConstruct
	// A reference to a StreamProcessor resource.
	// Experimental.
	StreamProcessorRef() *StreamProcessorReference
}

// The jsii proxy for IStreamProcessorRef
type jsiiProxy_IStreamProcessorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStreamProcessorRef) StreamProcessorRef() *StreamProcessorReference {
	var returns *StreamProcessorReference
	_jsii_.Get(
		j,
		"streamProcessorRef",
		&returns,
	)
	return returns
}


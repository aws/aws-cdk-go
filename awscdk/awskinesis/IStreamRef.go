package awskinesis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Stream.
// Experimental.
type IStreamRef interface {
	constructs.IConstruct
	// A reference to a Stream resource.
	// Experimental.
	StreamRef() *StreamReference
}

// The jsii proxy for IStreamRef
type jsiiProxy_IStreamRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStreamRef) StreamRef() *StreamReference {
	var returns *StreamReference
	_jsii_.Get(
		j,
		"streamRef",
		&returns,
	)
	return returns
}


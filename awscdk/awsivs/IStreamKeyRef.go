package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamKey.
// Experimental.
type IStreamKeyRef interface {
	constructs.IConstruct
	// A reference to a StreamKey resource.
	// Experimental.
	StreamKeyRef() *StreamKeyReference
}

// The jsii proxy for IStreamKeyRef
type jsiiProxy_IStreamKeyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStreamKeyRef) StreamKeyRef() *StreamKeyReference {
	var returns *StreamKeyReference
	_jsii_.Get(
		j,
		"streamKeyRef",
		&returns,
	)
	return returns
}


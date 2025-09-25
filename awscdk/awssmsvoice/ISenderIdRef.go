package awssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SenderId.
// Experimental.
type ISenderIdRef interface {
	constructs.IConstruct
	// A reference to a SenderId resource.
	// Experimental.
	SenderIdRef() *SenderIdReference
}

// The jsii proxy for ISenderIdRef
type jsiiProxy_ISenderIdRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISenderIdRef) SenderIdRef() *SenderIdReference {
	var returns *SenderIdReference
	_jsii_.Get(
		j,
		"senderIdRef",
		&returns,
	)
	return returns
}


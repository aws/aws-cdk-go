package awsiotevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Input.
// Experimental.
type IInputRef interface {
	constructs.IConstruct
	// A reference to a Input resource.
	// Experimental.
	InputRef() *InputReference
}

// The jsii proxy for IInputRef
type jsiiProxy_IInputRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInputRef) InputRef() *InputReference {
	var returns *InputReference
	_jsii_.Get(
		j,
		"inputRef",
		&returns,
	)
	return returns
}


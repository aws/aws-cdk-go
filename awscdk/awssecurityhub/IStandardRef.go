package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Standard.
// Experimental.
type IStandardRef interface {
	constructs.IConstruct
	// A reference to a Standard resource.
	// Experimental.
	StandardRef() *StandardReference
}

// The jsii proxy for IStandardRef
type jsiiProxy_IStandardRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStandardRef) StandardRef() *StandardReference {
	var returns *StandardReference
	_jsii_.Get(
		j,
		"standardRef",
		&returns,
	)
	return returns
}


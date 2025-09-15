package awssam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Application.
// Experimental.
type IApplicationRef interface {
	constructs.IConstruct
	// A reference to a Application resource.
	// Experimental.
	ApplicationRef() *ApplicationReference
}

// The jsii proxy for IApplicationRef
type jsiiProxy_IApplicationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationRef) ApplicationRef() *ApplicationReference {
	var returns *ApplicationReference
	_jsii_.Get(
		j,
		"applicationRef",
		&returns,
	)
	return returns
}


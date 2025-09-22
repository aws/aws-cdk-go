package awschatbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awschatbot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomAction.
// Experimental.
type ICustomActionRef interface {
	constructs.IConstruct
	// A reference to a CustomAction resource.
	// Experimental.
	CustomActionRef() *CustomActionReference
}

// The jsii proxy for ICustomActionRef
type jsiiProxy_ICustomActionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICustomActionRef) CustomActionRef() *CustomActionReference {
	var returns *CustomActionReference
	_jsii_.Get(
		j,
		"customActionRef",
		&returns,
	)
	return returns
}


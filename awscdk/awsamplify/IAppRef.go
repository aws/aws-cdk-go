package awsamplify

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplify/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a App.
// Experimental.
type IAppRef interface {
	constructs.IConstruct
	// A reference to a App resource.
	// Experimental.
	AppRef() *AppReference
}

// The jsii proxy for IAppRef
type jsiiProxy_IAppRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAppRef) AppRef() *AppReference {
	var returns *AppReference
	_jsii_.Get(
		j,
		"appRef",
		&returns,
	)
	return returns
}


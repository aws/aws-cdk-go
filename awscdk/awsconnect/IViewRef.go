package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a View.
// Experimental.
type IViewRef interface {
	constructs.IConstruct
	// A reference to a View resource.
	// Experimental.
	ViewRef() *ViewReference
}

// The jsii proxy for IViewRef
type jsiiProxy_IViewRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IViewRef) ViewRef() *ViewReference {
	var returns *ViewReference
	_jsii_.Get(
		j,
		"viewRef",
		&returns,
	)
	return returns
}


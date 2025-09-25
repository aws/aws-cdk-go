package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ViewVersion.
// Experimental.
type IViewVersionRef interface {
	constructs.IConstruct
	// A reference to a ViewVersion resource.
	// Experimental.
	ViewVersionRef() *ViewVersionReference
}

// The jsii proxy for IViewVersionRef
type jsiiProxy_IViewVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IViewVersionRef) ViewVersionRef() *ViewVersionReference {
	var returns *ViewVersionReference
	_jsii_.Get(
		j,
		"viewVersionRef",
		&returns,
	)
	return returns
}


package awsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Component.
// Experimental.
type IComponentRef interface {
	constructs.IConstruct
	// A reference to a Component resource.
	// Experimental.
	ComponentRef() *ComponentReference
}

// The jsii proxy for IComponentRef
type jsiiProxy_IComponentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IComponentRef) ComponentRef() *ComponentReference {
	var returns *ComponentReference
	_jsii_.Get(
		j,
		"componentRef",
		&returns,
	)
	return returns
}


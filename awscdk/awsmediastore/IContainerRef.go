package awsmediastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediastore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Container.
// Experimental.
type IContainerRef interface {
	constructs.IConstruct
	// A reference to a Container resource.
	// Experimental.
	ContainerRef() *ContainerReference
}

// The jsii proxy for IContainerRef
type jsiiProxy_IContainerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IContainerRef) ContainerRef() *ContainerReference {
	var returns *ContainerReference
	_jsii_.Get(
		j,
		"containerRef",
		&returns,
	)
	return returns
}


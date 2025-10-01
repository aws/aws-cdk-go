package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Space.
// Experimental.
type ISpaceRef interface {
	constructs.IConstruct
	// A reference to a Space resource.
	// Experimental.
	SpaceRef() *SpaceReference
}

// The jsii proxy for ISpaceRef
type jsiiProxy_ISpaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISpaceRef) SpaceRef() *SpaceReference {
	var returns *SpaceReference
	_jsii_.Get(
		j,
		"spaceRef",
		&returns,
	)
	return returns
}


package awsssmcontacts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmcontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Rotation.
// Experimental.
type IRotationRef interface {
	constructs.IConstruct
	// A reference to a Rotation resource.
	// Experimental.
	RotationRef() *RotationReference
}

// The jsii proxy for IRotationRef
type jsiiProxy_IRotationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRotationRef) RotationRef() *RotationReference {
	var returns *RotationReference
	_jsii_.Get(
		j,
		"rotationRef",
		&returns,
	)
	return returns
}


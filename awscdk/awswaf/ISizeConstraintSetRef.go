package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswaf/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SizeConstraintSet.
// Experimental.
type ISizeConstraintSetRef interface {
	constructs.IConstruct
	// A reference to a SizeConstraintSet resource.
	// Experimental.
	SizeConstraintSetRef() *SizeConstraintSetReference
}

// The jsii proxy for ISizeConstraintSetRef
type jsiiProxy_ISizeConstraintSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISizeConstraintSetRef) SizeConstraintSetRef() *SizeConstraintSetReference {
	var returns *SizeConstraintSetReference
	_jsii_.Get(
		j,
		"sizeConstraintSetRef",
		&returns,
	)
	return returns
}


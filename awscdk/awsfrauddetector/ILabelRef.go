package awsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Label.
// Experimental.
type ILabelRef interface {
	constructs.IConstruct
	// A reference to a Label resource.
	// Experimental.
	LabelRef() *LabelReference
}

// The jsii proxy for ILabelRef
type jsiiProxy_ILabelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILabelRef) LabelRef() *LabelReference {
	var returns *LabelReference
	_jsii_.Get(
		j,
		"labelRef",
		&returns,
	)
	return returns
}


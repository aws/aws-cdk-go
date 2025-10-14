package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationOutput.
// Experimental.
type IApplicationOutputRef interface {
	constructs.IConstruct
	// A reference to a ApplicationOutput resource.
	// Experimental.
	ApplicationOutputRef() *ApplicationOutputReference
}

// The jsii proxy for IApplicationOutputRef
type jsiiProxy_IApplicationOutputRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationOutputRef) ApplicationOutputRef() *ApplicationOutputReference {
	var returns *ApplicationOutputReference
	_jsii_.Get(
		j,
		"applicationOutputRef",
		&returns,
	)
	return returns
}


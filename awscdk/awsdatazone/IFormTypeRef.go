package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FormType.
// Experimental.
type IFormTypeRef interface {
	constructs.IConstruct
	// A reference to a FormType resource.
	// Experimental.
	FormTypeRef() *FormTypeReference
}

// The jsii proxy for IFormTypeRef
type jsiiProxy_IFormTypeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFormTypeRef) FormTypeRef() *FormTypeReference {
	var returns *FormTypeReference
	_jsii_.Get(
		j,
		"formTypeRef",
		&returns,
	)
	return returns
}


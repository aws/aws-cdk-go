package awsamplifyuibuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplifyuibuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Form.
// Experimental.
type IFormRef interface {
	constructs.IConstruct
	// A reference to a Form resource.
	// Experimental.
	FormRef() *FormReference
}

// The jsii proxy for IFormRef
type jsiiProxy_IFormRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFormRef) FormRef() *FormReference {
	var returns *FormReference
	_jsii_.Get(
		j,
		"formRef",
		&returns,
	)
	return returns
}


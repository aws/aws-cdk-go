package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswaf/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a XssMatchSet.
// Experimental.
type IXssMatchSetRef interface {
	constructs.IConstruct
	// A reference to a XssMatchSet resource.
	// Experimental.
	XssMatchSetRef() *XssMatchSetReference
}

// The jsii proxy for IXssMatchSetRef
type jsiiProxy_IXssMatchSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IXssMatchSetRef) XssMatchSetRef() *XssMatchSetReference {
	var returns *XssMatchSetReference
	_jsii_.Get(
		j,
		"xssMatchSetRef",
		&returns,
	)
	return returns
}


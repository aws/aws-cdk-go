package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Faq.
// Experimental.
type IFaqRef interface {
	constructs.IConstruct
	// A reference to a Faq resource.
	// Experimental.
	FaqRef() *FaqReference
}

// The jsii proxy for IFaqRef
type jsiiProxy_IFaqRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFaqRef) FaqRef() *FaqReference {
	var returns *FaqReference
	_jsii_.Get(
		j,
		"faqRef",
		&returns,
	)
	return returns
}


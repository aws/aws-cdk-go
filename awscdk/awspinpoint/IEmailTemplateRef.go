package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailTemplate.
// Experimental.
type IEmailTemplateRef interface {
	constructs.IConstruct
	// A reference to a EmailTemplate resource.
	// Experimental.
	EmailTemplateRef() *EmailTemplateReference
}

// The jsii proxy for IEmailTemplateRef
type jsiiProxy_IEmailTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEmailTemplateRef) EmailTemplateRef() *EmailTemplateReference {
	var returns *EmailTemplateReference
	_jsii_.Get(
		j,
		"emailTemplateRef",
		&returns,
	)
	return returns
}


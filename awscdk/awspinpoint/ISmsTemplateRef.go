package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SmsTemplate.
// Experimental.
type ISmsTemplateRef interface {
	constructs.IConstruct
	// A reference to a SmsTemplate resource.
	// Experimental.
	SmsTemplateRef() *SmsTemplateReference
}

// The jsii proxy for ISmsTemplateRef
type jsiiProxy_ISmsTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISmsTemplateRef) SmsTemplateRef() *SmsTemplateReference {
	var returns *SmsTemplateReference
	_jsii_.Get(
		j,
		"smsTemplateRef",
		&returns,
	)
	return returns
}


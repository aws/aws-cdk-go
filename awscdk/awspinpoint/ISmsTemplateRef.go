package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SmsTemplate.
// Experimental.
type ISmsTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISmsTemplateRef
type jsiiProxy_ISmsTemplateRef struct {
	internal.Type__constructsIConstruct
}


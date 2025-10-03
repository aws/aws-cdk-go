package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailTemplate.
// Experimental.
type IEmailTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEmailTemplateRef
type jsiiProxy_IEmailTemplateRef struct {
	internal.Type__constructsIConstruct
}


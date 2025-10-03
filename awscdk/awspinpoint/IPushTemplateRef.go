package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PushTemplate.
// Experimental.
type IPushTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPushTemplateRef
type jsiiProxy_IPushTemplateRef struct {
	internal.Type__constructsIConstruct
}


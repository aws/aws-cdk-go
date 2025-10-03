package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InAppTemplate.
// Experimental.
type IInAppTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInAppTemplateRef
type jsiiProxy_IInAppTemplateRef struct {
	internal.Type__constructsIConstruct
}


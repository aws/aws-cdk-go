package awsmediaconvert

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconvert/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a JobTemplate.
// Experimental.
type IJobTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IJobTemplateRef
type jsiiProxy_IJobTemplateRef struct {
	internal.Type__constructsIConstruct
}


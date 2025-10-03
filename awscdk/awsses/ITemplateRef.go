package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Template.
// Experimental.
type ITemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITemplateRef
type jsiiProxy_ITemplateRef struct {
	internal.Type__constructsIConstruct
}


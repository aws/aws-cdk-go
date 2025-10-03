package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QuickResponse.
// Experimental.
type IQuickResponseRef interface {
	constructs.IConstruct
}

// The jsii proxy for IQuickResponseRef
type jsiiProxy_IQuickResponseRef struct {
	internal.Type__constructsIConstruct
}


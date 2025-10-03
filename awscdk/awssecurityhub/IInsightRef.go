package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Insight.
// Experimental.
type IInsightRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInsightRef
type jsiiProxy_IInsightRef struct {
	internal.Type__constructsIConstruct
}


package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FindingAggregator.
// Experimental.
type IFindingAggregatorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFindingAggregatorRef
type jsiiProxy_IFindingAggregatorRef struct {
	internal.Type__constructsIConstruct
}


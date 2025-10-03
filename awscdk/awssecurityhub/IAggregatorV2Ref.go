package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AggregatorV2.
// Experimental.
type IAggregatorV2Ref interface {
	constructs.IConstruct
}

// The jsii proxy for IAggregatorV2Ref
type jsiiProxy_IAggregatorV2Ref struct {
	internal.Type__constructsIConstruct
}


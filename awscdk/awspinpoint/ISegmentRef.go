package awspinpoint

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Segment.
// Experimental.
type ISegmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISegmentRef
type jsiiProxy_ISegmentRef struct {
	internal.Type__constructsIConstruct
}


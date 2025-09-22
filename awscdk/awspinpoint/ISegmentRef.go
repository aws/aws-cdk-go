package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Segment.
// Experimental.
type ISegmentRef interface {
	constructs.IConstruct
	// A reference to a Segment resource.
	// Experimental.
	SegmentRef() *SegmentReference
}

// The jsii proxy for ISegmentRef
type jsiiProxy_ISegmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISegmentRef) SegmentRef() *SegmentReference {
	var returns *SegmentReference
	_jsii_.Get(
		j,
		"segmentRef",
		&returns,
	)
	return returns
}


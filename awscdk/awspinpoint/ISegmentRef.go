package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Segment.
// Experimental.
type ISegmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Segment resource.
	// Experimental.
	SegmentRef() *SegmentReference
}

// The jsii proxy for ISegmentRef
type jsiiProxy_ISegmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ISegmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISegmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


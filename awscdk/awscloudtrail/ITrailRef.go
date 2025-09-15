package awscloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Trail.
// Experimental.
type ITrailRef interface {
	constructs.IConstruct
	// A reference to a Trail resource.
	// Experimental.
	TrailRef() *TrailReference
}

// The jsii proxy for ITrailRef
type jsiiProxy_ITrailRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrailRef) TrailRef() *TrailReference {
	var returns *TrailReference
	_jsii_.Get(
		j,
		"trailRef",
		&returns,
	)
	return returns
}


package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficMirrorTarget.
// Experimental.
type ITrafficMirrorTargetRef interface {
	constructs.IConstruct
	// A reference to a TrafficMirrorTarget resource.
	// Experimental.
	TrafficMirrorTargetRef() *TrafficMirrorTargetReference
}

// The jsii proxy for ITrafficMirrorTargetRef
type jsiiProxy_ITrafficMirrorTargetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrafficMirrorTargetRef) TrafficMirrorTargetRef() *TrafficMirrorTargetReference {
	var returns *TrafficMirrorTargetReference
	_jsii_.Get(
		j,
		"trafficMirrorTargetRef",
		&returns,
	)
	return returns
}


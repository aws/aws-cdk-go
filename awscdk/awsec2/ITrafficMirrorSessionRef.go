package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficMirrorSession.
// Experimental.
type ITrafficMirrorSessionRef interface {
	constructs.IConstruct
	// A reference to a TrafficMirrorSession resource.
	// Experimental.
	TrafficMirrorSessionRef() *TrafficMirrorSessionReference
}

// The jsii proxy for ITrafficMirrorSessionRef
type jsiiProxy_ITrafficMirrorSessionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrafficMirrorSessionRef) TrafficMirrorSessionRef() *TrafficMirrorSessionReference {
	var returns *TrafficMirrorSessionReference
	_jsii_.Get(
		j,
		"trafficMirrorSessionRef",
		&returns,
	)
	return returns
}


package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficMirrorFilter.
// Experimental.
type ITrafficMirrorFilterRef interface {
	constructs.IConstruct
	// A reference to a TrafficMirrorFilter resource.
	// Experimental.
	TrafficMirrorFilterRef() *TrafficMirrorFilterReference
}

// The jsii proxy for ITrafficMirrorFilterRef
type jsiiProxy_ITrafficMirrorFilterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrafficMirrorFilterRef) TrafficMirrorFilterRef() *TrafficMirrorFilterReference {
	var returns *TrafficMirrorFilterReference
	_jsii_.Get(
		j,
		"trafficMirrorFilterRef",
		&returns,
	)
	return returns
}


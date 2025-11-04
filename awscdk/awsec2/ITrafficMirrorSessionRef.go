package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficMirrorSession.
// Experimental.
type ITrafficMirrorSessionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TrafficMirrorSession resource.
	// Experimental.
	TrafficMirrorSessionRef() *TrafficMirrorSessionReference
}

// The jsii proxy for ITrafficMirrorSessionRef
type jsiiProxy_ITrafficMirrorSessionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ITrafficMirrorSessionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrafficMirrorSessionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


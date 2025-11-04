package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficMirrorFilterRule.
// Experimental.
type ITrafficMirrorFilterRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TrafficMirrorFilterRule resource.
	// Experimental.
	TrafficMirrorFilterRuleRef() *TrafficMirrorFilterRuleReference
}

// The jsii proxy for ITrafficMirrorFilterRuleRef
type jsiiProxy_ITrafficMirrorFilterRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITrafficMirrorFilterRuleRef) TrafficMirrorFilterRuleRef() *TrafficMirrorFilterRuleReference {
	var returns *TrafficMirrorFilterRuleReference
	_jsii_.Get(
		j,
		"trafficMirrorFilterRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrafficMirrorFilterRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrafficMirrorFilterRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


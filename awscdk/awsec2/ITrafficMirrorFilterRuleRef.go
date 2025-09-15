package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficMirrorFilterRule.
// Experimental.
type ITrafficMirrorFilterRuleRef interface {
	constructs.IConstruct
	// A reference to a TrafficMirrorFilterRule resource.
	// Experimental.
	TrafficMirrorFilterRuleRef() *TrafficMirrorFilterRuleReference
}

// The jsii proxy for ITrafficMirrorFilterRuleRef
type jsiiProxy_ITrafficMirrorFilterRuleRef struct {
	internal.Type__constructsIConstruct
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


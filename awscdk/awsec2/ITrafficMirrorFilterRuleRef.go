package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficMirrorFilterRule.
// Experimental.
type ITrafficMirrorFilterRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITrafficMirrorFilterRuleRef
type jsiiProxy_ITrafficMirrorFilterRuleRef struct {
	internal.Type__constructsIConstruct
}


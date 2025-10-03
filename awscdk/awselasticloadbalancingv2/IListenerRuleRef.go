package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ListenerRule.
// Experimental.
type IListenerRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IListenerRuleRef
type jsiiProxy_IListenerRuleRef struct {
	internal.Type__constructsIConstruct
}


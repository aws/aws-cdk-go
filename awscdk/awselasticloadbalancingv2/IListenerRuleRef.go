package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ListenerRule.
// Experimental.
type IListenerRuleRef interface {
	constructs.IConstruct
	// A reference to a ListenerRule resource.
	// Experimental.
	ListenerRuleRef() *ListenerRuleReference
}

// The jsii proxy for IListenerRuleRef
type jsiiProxy_IListenerRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IListenerRuleRef) ListenerRuleRef() *ListenerRuleReference {
	var returns *ListenerRuleReference
	_jsii_.Get(
		j,
		"listenerRuleRef",
		&returns,
	)
	return returns
}


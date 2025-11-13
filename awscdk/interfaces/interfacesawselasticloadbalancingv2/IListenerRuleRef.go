package interfacesawselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ListenerRule.
// Experimental.
type IListenerRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ListenerRule resource.
	// Experimental.
	ListenerRuleRef() *ListenerRuleReference
}

// The jsii proxy for IListenerRuleRef
type jsiiProxy_IListenerRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IListenerRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IListenerRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


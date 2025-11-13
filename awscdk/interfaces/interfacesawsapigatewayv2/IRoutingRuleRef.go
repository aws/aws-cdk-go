package interfacesawsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RoutingRule.
// Experimental.
type IRoutingRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RoutingRule resource.
	// Experimental.
	RoutingRuleRef() *RoutingRuleReference
}

// The jsii proxy for IRoutingRuleRef
type jsiiProxy_IRoutingRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRoutingRuleRef) RoutingRuleRef() *RoutingRuleReference {
	var returns *RoutingRuleReference
	_jsii_.Get(
		j,
		"routingRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoutingRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoutingRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


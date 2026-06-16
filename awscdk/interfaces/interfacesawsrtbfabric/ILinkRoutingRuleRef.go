package interfacesawsrtbfabric

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrtbfabric/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LinkRoutingRule.
// Experimental.
type ILinkRoutingRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LinkRoutingRule resource.
	// Experimental.
	LinkRoutingRuleRef() *LinkRoutingRuleReference
}

// The jsii proxy for ILinkRoutingRuleRef
type jsiiProxy_ILinkRoutingRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILinkRoutingRuleRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ILinkRoutingRuleRef) LinkRoutingRuleRef() *LinkRoutingRuleReference {
	var returns *LinkRoutingRuleReference
	_jsii_.Get(
		j,
		"linkRoutingRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILinkRoutingRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILinkRoutingRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallRuleGroup.
// Experimental.
type IFirewallRuleGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FirewallRuleGroup resource.
	// Experimental.
	FirewallRuleGroupRef() *FirewallRuleGroupReference
}

// The jsii proxy for IFirewallRuleGroupRef
type jsiiProxy_IFirewallRuleGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IFirewallRuleGroupRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFirewallRuleGroupRef) FirewallRuleGroupRef() *FirewallRuleGroupReference {
	var returns *FirewallRuleGroupReference
	_jsii_.Get(
		j,
		"firewallRuleGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFirewallRuleGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFirewallRuleGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


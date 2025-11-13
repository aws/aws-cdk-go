package interfacesawsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallRuleGroupAssociation.
// Experimental.
type IFirewallRuleGroupAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FirewallRuleGroupAssociation resource.
	// Experimental.
	FirewallRuleGroupAssociationRef() *FirewallRuleGroupAssociationReference
}

// The jsii proxy for IFirewallRuleGroupAssociationRef
type jsiiProxy_IFirewallRuleGroupAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IFirewallRuleGroupAssociationRef) FirewallRuleGroupAssociationRef() *FirewallRuleGroupAssociationReference {
	var returns *FirewallRuleGroupAssociationReference
	_jsii_.Get(
		j,
		"firewallRuleGroupAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFirewallRuleGroupAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFirewallRuleGroupAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallRuleGroup.
// Experimental.
type IFirewallRuleGroupRef interface {
	constructs.IConstruct
	// A reference to a FirewallRuleGroup resource.
	// Experimental.
	FirewallRuleGroupRef() *FirewallRuleGroupReference
}

// The jsii proxy for IFirewallRuleGroupRef
type jsiiProxy_IFirewallRuleGroupRef struct {
	internal.Type__constructsIConstruct
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


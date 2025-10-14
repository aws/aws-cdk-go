package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallRuleGroupAssociation.
// Experimental.
type IFirewallRuleGroupAssociationRef interface {
	constructs.IConstruct
	// A reference to a FirewallRuleGroupAssociation resource.
	// Experimental.
	FirewallRuleGroupAssociationRef() *FirewallRuleGroupAssociationReference
}

// The jsii proxy for IFirewallRuleGroupAssociationRef
type jsiiProxy_IFirewallRuleGroupAssociationRef struct {
	internal.Type__constructsIConstruct
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


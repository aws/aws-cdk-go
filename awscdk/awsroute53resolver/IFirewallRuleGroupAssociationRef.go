package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallRuleGroupAssociation.
// Experimental.
type IFirewallRuleGroupAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFirewallRuleGroupAssociationRef
type jsiiProxy_IFirewallRuleGroupAssociationRef struct {
	internal.Type__constructsIConstruct
}


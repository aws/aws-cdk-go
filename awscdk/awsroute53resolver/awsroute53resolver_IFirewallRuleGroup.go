package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53resolver/internal"
)

// A Firewall Rule Group.
// Experimental.
type IFirewallRuleGroup interface {
	awscdk.IResource
	// The ID of the rule group.
	// Experimental.
	FirewallRuleGroupId() *string
}

// The jsii proxy for IFirewallRuleGroup
type jsiiProxy_IFirewallRuleGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IFirewallRuleGroup) FirewallRuleGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupId",
		&returns,
	)
	return returns
}


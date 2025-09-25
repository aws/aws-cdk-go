package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallPolicy.
// Experimental.
type IFirewallPolicyRef interface {
	constructs.IConstruct
	// A reference to a FirewallPolicy resource.
	// Experimental.
	FirewallPolicyRef() *FirewallPolicyReference
}

// The jsii proxy for IFirewallPolicyRef
type jsiiProxy_IFirewallPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFirewallPolicyRef) FirewallPolicyRef() *FirewallPolicyReference {
	var returns *FirewallPolicyReference
	_jsii_.Get(
		j,
		"firewallPolicyRef",
		&returns,
	)
	return returns
}


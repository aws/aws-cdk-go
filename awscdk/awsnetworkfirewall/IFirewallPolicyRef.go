package awsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallPolicy.
// Experimental.
type IFirewallPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFirewallPolicyRef
type jsiiProxy_IFirewallPolicyRef struct {
	internal.Type__constructsIConstruct
}


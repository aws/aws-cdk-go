package awsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Firewall.
// Experimental.
type IFirewallRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFirewallRef
type jsiiProxy_IFirewallRef struct {
	internal.Type__constructsIConstruct
}


package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Firewall.
// Experimental.
type IFirewallRef interface {
	constructs.IConstruct
	// A reference to a Firewall resource.
	// Experimental.
	FirewallRef() *FirewallReference
}

// The jsii proxy for IFirewallRef
type jsiiProxy_IFirewallRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFirewallRef) FirewallRef() *FirewallReference {
	var returns *FirewallReference
	_jsii_.Get(
		j,
		"firewallRef",
		&returns,
	)
	return returns
}


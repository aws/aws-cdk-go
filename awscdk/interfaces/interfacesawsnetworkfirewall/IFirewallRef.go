package interfacesawsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Firewall.
// Experimental.
type IFirewallRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Firewall resource.
	// Experimental.
	FirewallRef() *FirewallReference
}

// The jsii proxy for IFirewallRef
type jsiiProxy_IFirewallRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IFirewallRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFirewallRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


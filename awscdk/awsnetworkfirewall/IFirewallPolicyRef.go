package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallPolicy.
// Experimental.
type IFirewallPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a FirewallPolicy resource.
	// Experimental.
	FirewallPolicyRef() *FirewallPolicyReference
}

// The jsii proxy for IFirewallPolicyRef
type jsiiProxy_IFirewallPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IFirewallPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFirewallPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


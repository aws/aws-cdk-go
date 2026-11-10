package interfacesawsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallConfig.
// Experimental.
type IFirewallConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FirewallConfig resource.
	// Experimental.
	FirewallConfigRef() *FirewallConfigReference
}

// The jsii proxy for IFirewallConfigRef
type jsiiProxy_IFirewallConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IFirewallConfigRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IFirewallConfigRef) FirewallConfigRef() *FirewallConfigReference {
	var returns *FirewallConfigReference
	_jsii_.Get(
		j,
		"firewallConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFirewallConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFirewallConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


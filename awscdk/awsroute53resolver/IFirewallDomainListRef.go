package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallDomainList.
// Experimental.
type IFirewallDomainListRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a FirewallDomainList resource.
	// Experimental.
	FirewallDomainListRef() *FirewallDomainListReference
}

// The jsii proxy for IFirewallDomainListRef
type jsiiProxy_IFirewallDomainListRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFirewallDomainListRef) FirewallDomainListRef() *FirewallDomainListReference {
	var returns *FirewallDomainListReference
	_jsii_.Get(
		j,
		"firewallDomainListRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFirewallDomainListRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFirewallDomainListRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


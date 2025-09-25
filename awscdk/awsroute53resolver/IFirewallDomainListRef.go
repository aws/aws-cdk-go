package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53resolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FirewallDomainList.
// Experimental.
type IFirewallDomainListRef interface {
	constructs.IConstruct
	// A reference to a FirewallDomainList resource.
	// Experimental.
	FirewallDomainListRef() *FirewallDomainListReference
}

// The jsii proxy for IFirewallDomainListRef
type jsiiProxy_IFirewallDomainListRef struct {
	internal.Type__constructsIConstruct
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


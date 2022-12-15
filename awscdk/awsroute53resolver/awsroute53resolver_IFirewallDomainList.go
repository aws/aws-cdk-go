package awsroute53resolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53resolver/internal"
)

// A Firewall Domain List.
// Experimental.
type IFirewallDomainList interface {
	awscdk.IResource
	// The ID of the domain list.
	// Experimental.
	FirewallDomainListId() *string
}

// The jsii proxy for IFirewallDomainList
type jsiiProxy_IFirewallDomainList struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IFirewallDomainList) FirewallDomainListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListId",
		&returns,
	)
	return returns
}


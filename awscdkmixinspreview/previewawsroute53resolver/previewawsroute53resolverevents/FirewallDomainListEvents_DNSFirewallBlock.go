package previewawsroute53resolverevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.route53resolver@DNSFirewallBlock event types for FirewallDomainList.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dNSFirewallBlock := #error#.NewDNSFirewallBlock()
//
// Experimental.
type FirewallDomainListEvents_DNSFirewallBlock interface {
}

// The jsii proxy struct for FirewallDomainListEvents_DNSFirewallBlock
type jsiiProxy_FirewallDomainListEvents_DNSFirewallBlock struct {
	_ byte // padding
}

// Experimental.
func NewFirewallDomainListEvents_DNSFirewallBlock() FirewallDomainListEvents_DNSFirewallBlock {
	_init_.Initialize()

	j := jsiiProxy_FirewallDomainListEvents_DNSFirewallBlock{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallBlock",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallDomainListEvents_DNSFirewallBlock_Override(f FirewallDomainListEvents_DNSFirewallBlock) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallBlock",
		nil, // no parameters
		f,
	)
}


package previewawsroute53resolverevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.route53resolver@DNSFirewallAlert event types for FirewallDomainList.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dNSFirewallAlert := #error#.NewDNSFirewallAlert()
//
// Experimental.
type FirewallDomainListEvents_DNSFirewallAlert interface {
}

// The jsii proxy struct for FirewallDomainListEvents_DNSFirewallAlert
type jsiiProxy_FirewallDomainListEvents_DNSFirewallAlert struct {
	_ byte // padding
}

// Experimental.
func NewFirewallDomainListEvents_DNSFirewallAlert() FirewallDomainListEvents_DNSFirewallAlert {
	_init_.Initialize()

	j := jsiiProxy_FirewallDomainListEvents_DNSFirewallAlert{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallAlert",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallDomainListEvents_DNSFirewallAlert_Override(f FirewallDomainListEvents_DNSFirewallAlert) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallAlert",
		nil, // no parameters
		f,
	)
}


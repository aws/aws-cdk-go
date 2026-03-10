package previewawsroute53resolverevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.route53resolver@DNSFirewallBlock.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dNSFirewallBlock := awscdkmixinspreview.Events.NewDNSFirewallBlock()
//
// Experimental.
type DNSFirewallBlock interface {
}

// The jsii proxy struct for DNSFirewallBlock
type jsiiProxy_DNSFirewallBlock struct {
	_ byte // padding
}

// Experimental.
func NewDNSFirewallBlock() DNSFirewallBlock {
	_init_.Initialize()

	j := jsiiProxy_DNSFirewallBlock{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallBlock",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDNSFirewallBlock_Override(d DNSFirewallBlock) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallBlock",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DNS Firewall Block.
// Experimental.
func DNSFirewallBlock_DnsFirewallBlockPattern(options *DNSFirewallBlock_DNSFirewallBlockProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDNSFirewallBlock_DnsFirewallBlockPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallBlock",
		"dnsFirewallBlockPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


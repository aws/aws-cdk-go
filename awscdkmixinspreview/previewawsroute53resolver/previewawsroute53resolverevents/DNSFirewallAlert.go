package previewawsroute53resolverevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.route53resolver@DNSFirewallAlert.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dNSFirewallAlert := awscdkmixinspreview.Events.NewDNSFirewallAlert()
//
// Experimental.
type DNSFirewallAlert interface {
}

// The jsii proxy struct for DNSFirewallAlert
type jsiiProxy_DNSFirewallAlert struct {
	_ byte // padding
}

// Experimental.
func NewDNSFirewallAlert() DNSFirewallAlert {
	_init_.Initialize()

	j := jsiiProxy_DNSFirewallAlert{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallAlert",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDNSFirewallAlert_Override(d DNSFirewallAlert) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallAlert",
		nil, // no parameters
		d,
	)
}

// EventBridge event pattern for DNS Firewall Alert.
// Experimental.
func DNSFirewallAlert_DnsFirewallAlertPattern(options *DNSFirewallAlert_DNSFirewallAlertProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateDNSFirewallAlert_DnsFirewallAlertPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallAlert",
		"dnsFirewallAlertPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


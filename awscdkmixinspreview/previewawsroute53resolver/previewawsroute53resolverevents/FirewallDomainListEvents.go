package previewawsroute53resolverevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53resolver"
)

// EventBridge event patterns for FirewallDomainList.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var firewallDomainListRef IFirewallDomainListRef
//
//   firewallDomainListEvents := awscdkmixinspreview.Events.FirewallDomainListEvents_FromFirewallDomainList(firewallDomainListRef)
//
// Experimental.
type FirewallDomainListEvents interface {
	// EventBridge event pattern for FirewallDomainList DNS Firewall Alert.
	// Experimental.
	DNSFirewallAlertPattern(options *FirewallDomainListEvents_DNSFirewallAlert_DNSFirewallAlertProps) *awsevents.EventPattern
	// EventBridge event pattern for FirewallDomainList DNS Firewall Block.
	// Experimental.
	DNSFirewallBlockPattern(options *FirewallDomainListEvents_DNSFirewallBlock_DNSFirewallBlockProps) *awsevents.EventPattern
}

// The jsii proxy struct for FirewallDomainListEvents
type jsiiProxy_FirewallDomainListEvents struct {
	_ byte // padding
}

// Create FirewallDomainListEvents from a FirewallDomainList reference.
// Experimental.
func FirewallDomainListEvents_FromFirewallDomainList(firewallDomainListRef interfacesawsroute53resolver.IFirewallDomainListRef) FirewallDomainListEvents {
	_init_.Initialize()

	if err := validateFirewallDomainListEvents_FromFirewallDomainListParameters(firewallDomainListRef); err != nil {
		panic(err)
	}
	var returns FirewallDomainListEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents",
		"fromFirewallDomainList",
		[]interface{}{firewallDomainListRef},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainListEvents) DNSFirewallAlertPattern(options *FirewallDomainListEvents_DNSFirewallAlert_DNSFirewallAlertProps) *awsevents.EventPattern {
	if err := f.validateDNSFirewallAlertPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		f,
		"dNSFirewallAlertPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainListEvents) DNSFirewallBlockPattern(options *FirewallDomainListEvents_DNSFirewallBlock_DNSFirewallBlockProps) *awsevents.EventPattern {
	if err := f.validateDNSFirewallBlockPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		f,
		"dNSFirewallBlockPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


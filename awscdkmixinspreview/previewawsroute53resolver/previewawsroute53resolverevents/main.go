package previewawsroute53resolverevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallAlert",
		reflect.TypeOf((*DNSFirewallAlert)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DNSFirewallAlert{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallAlert.DNSFirewallAlertProps",
		reflect.TypeOf((*DNSFirewallAlert_DNSFirewallAlertProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallAlert.DnsFirewallAlertItem",
		reflect.TypeOf((*DNSFirewallAlert_DnsFirewallAlertItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallAlert.ResolverEndpointDetails",
		reflect.TypeOf((*DNSFirewallAlert_ResolverEndpointDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallAlert.ResolverNetworkInterfaceDetails",
		reflect.TypeOf((*DNSFirewallAlert_ResolverNetworkInterfaceDetails)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallBlock",
		reflect.TypeOf((*DNSFirewallBlock)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DNSFirewallBlock{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallBlock.DNSFirewallBlockProps",
		reflect.TypeOf((*DNSFirewallBlock_DNSFirewallBlockProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallBlock.DnsFirewallBlockItem",
		reflect.TypeOf((*DNSFirewallBlock_DnsFirewallBlockItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallBlock.ResolverEndpointDetails",
		reflect.TypeOf((*DNSFirewallBlock_ResolverEndpointDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.DNSFirewallBlock.ResolverNetworkInterfaceDetails",
		reflect.TypeOf((*DNSFirewallBlock_ResolverNetworkInterfaceDetails)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents",
		reflect.TypeOf((*FirewallDomainListEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "dnsFirewallAlertPattern", GoMethod: "DnsFirewallAlertPattern"},
			_jsii_.MemberMethod{JsiiMethod: "dnsFirewallBlockPattern", GoMethod: "DnsFirewallBlockPattern"},
		},
		func() interface{} {
			return &jsiiProxy_FirewallDomainListEvents{}
		},
	)
}

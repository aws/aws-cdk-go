package previewawsroute53resolverevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents",
		reflect.TypeOf((*FirewallDomainListEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "dNSFirewallAlertPattern", GoMethod: "DNSFirewallAlertPattern"},
			_jsii_.MemberMethod{JsiiMethod: "dNSFirewallBlockPattern", GoMethod: "DNSFirewallBlockPattern"},
		},
		func() interface{} {
			return &jsiiProxy_FirewallDomainListEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallAlert",
		reflect.TypeOf((*FirewallDomainListEvents_DNSFirewallAlert)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FirewallDomainListEvents_DNSFirewallAlert{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallAlert.DNSFirewallAlertProps",
		reflect.TypeOf((*FirewallDomainListEvents_DNSFirewallAlert_DNSFirewallAlertProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallAlert.DnsFirewallAlertItem",
		reflect.TypeOf((*FirewallDomainListEvents_DNSFirewallAlert_DnsFirewallAlertItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallAlert.ResolverEndpointDetails",
		reflect.TypeOf((*FirewallDomainListEvents_DNSFirewallAlert_ResolverEndpointDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallAlert.ResolverNetworkInterfaceDetails",
		reflect.TypeOf((*FirewallDomainListEvents_DNSFirewallAlert_ResolverNetworkInterfaceDetails)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallBlock",
		reflect.TypeOf((*FirewallDomainListEvents_DNSFirewallBlock)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_FirewallDomainListEvents_DNSFirewallBlock{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallBlock.DNSFirewallBlockProps",
		reflect.TypeOf((*FirewallDomainListEvents_DNSFirewallBlock_DNSFirewallBlockProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallBlock.DnsFirewallBlockItem",
		reflect.TypeOf((*FirewallDomainListEvents_DNSFirewallBlock_DnsFirewallBlockItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallBlock.ResolverEndpointDetails",
		reflect.TypeOf((*FirewallDomainListEvents_DNSFirewallBlock_ResolverEndpointDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53resolver.events.FirewallDomainListEvents.DNSFirewallBlock.ResolverNetworkInterfaceDetails",
		reflect.TypeOf((*FirewallDomainListEvents_DNSFirewallBlock_ResolverNetworkInterfaceDetails)(nil)).Elem(),
	)
}

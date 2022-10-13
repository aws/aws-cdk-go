package awsroute53resolver

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// The way that you want DNS Firewall to block the request.
//
// Example:
//   var myBlockList firewallDomainList
//   var ruleGroup firewallRuleGroup
//
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(10),
//   	firewallDomainList: myBlockList,
//   	// block and reply with NXDOMAIN
//   	action: route53resolver.firewallRuleAction.block(route53resolver.dnsBlockResponse.nxDomain()),
//   })
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(20),
//   	firewallDomainList: myBlockList,
//   	// block and override DNS response with a custom domain
//   	action: route53resolver.*firewallRuleAction.block(route53resolver.*dnsBlockResponse.override(jsii.String("amazon.com"))),
//   })
//
// Experimental.
type DnsBlockResponse interface {
	// The DNS record's type.
	// Experimental.
	BlockOverrideDnsType() *string
	// The custom DNS record to send back in response to the query.
	// Experimental.
	BlockOverrideDomain() *string
	// The recommended amount of time for the DNS resolver or web browser to cache the provided override record.
	// Experimental.
	BlockOverrideTtl() awscdk.Duration
	// The way that you want DNS Firewall to block the request.
	// Experimental.
	BlockResponse() *string
}

// The jsii proxy struct for DnsBlockResponse
type jsiiProxy_DnsBlockResponse struct {
	_ byte // padding
}

func (j *jsiiProxy_DnsBlockResponse) BlockOverrideDnsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDnsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsBlockResponse) BlockOverrideDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockOverrideDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsBlockResponse) BlockOverrideTtl() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"blockOverrideTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsBlockResponse) BlockResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blockResponse",
		&returns,
	)
	return returns
}


// Experimental.
func NewDnsBlockResponse_Override(d DnsBlockResponse) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.DnsBlockResponse",
		nil, // no parameters
		d,
	)
}

// Respond indicating that the query was successful, but no response is available for it.
// Experimental.
func DnsBlockResponse_NoData() DnsBlockResponse {
	_init_.Initialize()

	var returns DnsBlockResponse

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.DnsBlockResponse",
		"noData",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Respond indicating that the domain name that's in the query doesn't exist.
// Experimental.
func DnsBlockResponse_NxDomain() DnsBlockResponse {
	_init_.Initialize()

	var returns DnsBlockResponse

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.DnsBlockResponse",
		"nxDomain",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Provides a custom override response to the query.
// Experimental.
func DnsBlockResponse_Override(domain *string, ttl awscdk.Duration) DnsBlockResponse {
	_init_.Initialize()

	if err := validateDnsBlockResponse_OverrideParameters(domain); err != nil {
		panic(err)
	}
	var returns DnsBlockResponse

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.DnsBlockResponse",
		"override",
		[]interface{}{domain, ttl},
		&returns,
	)

	return returns
}


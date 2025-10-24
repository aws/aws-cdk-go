package awscdkroute53resolveralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The way that you want DNS Firewall to block the request.
//
// Example:
//   var myBlockList FirewallDomainList
//   var ruleGroup FirewallRuleGroup
//
//
//   ruleGroup.AddRule(&FirewallRule{
//   	Priority: jsii.Number(10),
//   	FirewallDomainList: myBlockList,
//   	// block and reply with NXDOMAIN
//   	Action: route53resolver.FirewallRuleAction_Block(route53resolver.DnsBlockResponse_NxDomain()),
//   })
//
//   ruleGroup.AddRule(&FirewallRule{
//   	Priority: jsii.Number(20),
//   	FirewallDomainList: myBlockList,
//   	// block and override DNS response with a custom domain
//   	Action: route53resolver.FirewallRuleAction_*Block(route53resolver.DnsBlockResponse_Override(jsii.String("amazon.com"))),
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
		"@aws-cdk/aws-route53resolver-alpha.DnsBlockResponse",
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
		"@aws-cdk/aws-route53resolver-alpha.DnsBlockResponse",
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
		"@aws-cdk/aws-route53resolver-alpha.DnsBlockResponse",
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
		"@aws-cdk/aws-route53resolver-alpha.DnsBlockResponse",
		"override",
		[]interface{}{domain, ttl},
		&returns,
	)

	return returns
}


package awsroute53resolver

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
)

// A list of domains.
//
// Example:
//   blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &firewallDomainListProps{
//   	domains: route53resolver.firewallDomains.fromList([]*string{
//   		jsii.String("bad-domain.com"),
//   		jsii.String("bot-domain.net"),
//   	}),
//   })
//
//   s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromS3Url(jsii.String("s3://bucket/prefix/object")),
//   })
//
//   assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromAsset(jsii.String("/path/to/domains.txt")),
//   })
//
// Experimental.
type FirewallDomains interface {
	// Binds the domains to a domain list.
	// Experimental.
	Bind(scope constructs.Construct) *DomainsConfig
}

// The jsii proxy struct for FirewallDomains
type jsiiProxy_FirewallDomains struct {
	_ byte // padding
}

// Experimental.
func NewFirewallDomains_Override(f FirewallDomains) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_route53resolver.FirewallDomains",
		nil, // no parameters
		f,
	)
}

// Firewall domains created from a local disk path to a text file.
//
// The file must be a text file (`.txt` extension) and must contain a single
// domain per line. It will be uploaded to S3.
// Experimental.
func FirewallDomains_FromAsset(assetPath *string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomains",
		"fromAsset",
		[]interface{}{assetPath},
		&returns,
	)

	return returns
}

// Firewall domains created from a list of domains.
// Experimental.
func FirewallDomains_FromList(list *[]*string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomains",
		"fromList",
		[]interface{}{list},
		&returns,
	)

	return returns
}

// Firewall domains created from a file stored in Amazon S3.
//
// The file must be a text file and must contain a single domain per line.
// The content type of the S3 object must be `plain/text`.
// Experimental.
func FirewallDomains_FromS3(bucket awss3.IBucket, key *string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomains",
		"fromS3",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// Firewall domains created from the URL of a file stored in Amazon S3.
//
// The file must be a text file and must contain a single domain per line.
// The content type of the S3 object must be `plain/text`.
// Experimental.
func FirewallDomains_FromS3Url(url *string) FirewallDomains {
	_init_.Initialize()

	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"monocdk.aws_route53resolver.FirewallDomains",
		"fromS3Url",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomains) Bind(scope constructs.Construct) *DomainsConfig {
	var returns *DomainsConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


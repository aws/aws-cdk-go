package awscdkroute53resolveralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// A list of domains.
//
// Example:
//   blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &FirewallDomainListProps{
//   	Domains: route53resolver.FirewallDomains_FromList([]*string{
//   		jsii.String("bad-domain.com"),
//   		jsii.String("bot-domain.net"),
//   	}),
//   })
//
//   s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &FirewallDomainListProps{
//   	Domains: route53resolver.FirewallDomains_FromS3Url(jsii.String("s3://bucket/prefix/object")),
//   })
//
//   assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &FirewallDomainListProps{
//   	Domains: route53resolver.FirewallDomains_FromAsset(jsii.String("/path/to/domains.txt")),
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
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
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

	if err := validateFirewallDomains_FromAssetParameters(assetPath); err != nil {
		panic(err)
	}
	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
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

	if err := validateFirewallDomains_FromListParameters(list); err != nil {
		panic(err)
	}
	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
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

	if err := validateFirewallDomains_FromS3Parameters(bucket, key); err != nil {
		panic(err)
	}
	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
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

	if err := validateFirewallDomains_FromS3UrlParameters(url); err != nil {
		panic(err)
	}
	var returns FirewallDomains

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomains",
		"fromS3Url",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomains) Bind(scope constructs.Construct) *DomainsConfig {
	if err := f.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *DomainsConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


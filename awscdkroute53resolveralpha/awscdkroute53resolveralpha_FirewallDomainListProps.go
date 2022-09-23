// The CDK Construct Library for AWS::Route53Resolver
package awscdkroute53resolveralpha


// Properties for a Firewall Domain List.
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
type FirewallDomainListProps struct {
	// A list of domains.
	// Experimental.
	Domains FirewallDomains `field:"required" json:"domains" yaml:"domains"`
	// A name for the domain list.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}


package awsroute53resolver


// Properties for a Firewall Domain List.
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
type FirewallDomainListProps struct {
	// A list of domains.
	// Experimental.
	Domains FirewallDomains `field:"required" json:"domains" yaml:"domains"`
	// A name for the domain list.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}


package awsapigateway


// Example:
//   var api restApi
//
//
//   domainName := apigateway.DomainName_FromDomainNameAttributes(this, jsii.String("DomainName"), &DomainNameAttributes{
//   	DomainName: jsii.String("domainName"),
//   	DomainNameAliasHostedZoneId: jsii.String("domainNameAliasHostedZoneId"),
//   	DomainNameAliasTarget: jsii.String("domainNameAliasTarget"),
//   })
//
//   apigateway.NewBasePathMapping(this, jsii.String("BasePathMapping"), &BasePathMappingProps{
//   	DomainName: domainName,
//   	RestApi: api,
//   })
//
type DomainNameAttributes struct {
	// The domain name (e.g. `example.com`).
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The Route53 hosted zone ID to use in order to connect a record set to this domain through an alias.
	DomainNameAliasHostedZoneId *string `field:"required" json:"domainNameAliasHostedZoneId" yaml:"domainNameAliasHostedZoneId"`
	// The Route53 alias target to use in order to connect a record set to this domain through an alias.
	DomainNameAliasTarget *string `field:"required" json:"domainNameAliasTarget" yaml:"domainNameAliasTarget"`
}


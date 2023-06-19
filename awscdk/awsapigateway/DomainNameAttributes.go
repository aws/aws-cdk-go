package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainNameAttributes := &DomainNameAttributes{
//   	DomainName: jsii.String("domainName"),
//   	DomainNameAliasHostedZoneId: jsii.String("domainNameAliasHostedZoneId"),
//   	DomainNameAliasTarget: jsii.String("domainNameAliasTarget"),
//   }
//
// Experimental.
type DomainNameAttributes struct {
	// The domain name (e.g. `example.com`).
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The Route53 hosted zone ID to use in order to connect a record set to this domain through an alias.
	// Experimental.
	DomainNameAliasHostedZoneId *string `field:"required" json:"domainNameAliasHostedZoneId" yaml:"domainNameAliasHostedZoneId"`
	// The Route53 alias target to use in order to connect a record set to this domain through an alias.
	// Experimental.
	DomainNameAliasTarget *string `field:"required" json:"domainNameAliasTarget" yaml:"domainNameAliasTarget"`
}


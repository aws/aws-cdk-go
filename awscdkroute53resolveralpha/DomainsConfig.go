// The CDK Construct Library for AWS::Route53Resolver
package awscdkroute53resolveralpha


// Domains configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import route53resolver_alpha "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha"
//
//   domainsConfig := &DomainsConfig{
//   	DomainFileUrl: jsii.String("domainFileUrl"),
//   	Domains: []*string{
//   		jsii.String("domains"),
//   	},
//   }
//
// Experimental.
type DomainsConfig struct {
	// The fully qualified URL or URI of the file stored in Amazon S3 that contains the list of domains to import.
	//
	// The file must be a text file and must contain
	// a single domain per line. The content type of the S3 object must be `plain/text`.
	// Experimental.
	DomainFileUrl *string `field:"optional" json:"domainFileUrl" yaml:"domainFileUrl"`
	// A list of domains.
	// Experimental.
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
}


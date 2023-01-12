package awsopensearchservice


// Reference to an Amazon OpenSearch Service domain.
//
// Example:
//   domainArn := awscdk.Fn.importValue(jsii.String("another-cf-stack-export-domain-arn"))
//   domainEndpoint := awscdk.Fn.importValue(jsii.String("another-cf-stack-export-domain-endpoint"))
//   domain := awscdk.Domain.fromDomainAttributes(this, jsii.String("ImportedDomain"), &domainAttributes{
//   	domainArn: jsii.String(domainArn),
//   	domainEndpoint: jsii.String(domainEndpoint),
//   })
//
type DomainAttributes struct {
	// The ARN of the Amazon OpenSearch Service domain.
	DomainArn *string `field:"required" json:"domainArn" yaml:"domainArn"`
	// The domain endpoint of the Amazon OpenSearch Service domain.
	DomainEndpoint *string `field:"required" json:"domainEndpoint" yaml:"domainEndpoint"`
}


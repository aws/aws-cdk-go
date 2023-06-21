package awsopensearchservice


// Reference to an Amazon OpenSearch Service domain.
//
// Example:
//   domainArn := awscdk.Fn_ImportValue(jsii.String("another-cf-stack-export-domain-arn"))
//   domainEndpoint := awscdk.Fn_ImportValue(jsii.String("another-cf-stack-export-domain-endpoint"))
//   domain := awscdk.Domain_FromDomainAttributes(this, jsii.String("ImportedDomain"), &DomainAttributes{
//   	DomainArn: jsii.String(DomainArn),
//   	DomainEndpoint: jsii.String(DomainEndpoint),
//   })
//
type DomainAttributes struct {
	// The ARN of the Amazon OpenSearch Service domain.
	DomainArn *string `field:"required" json:"domainArn" yaml:"domainArn"`
	// The domain endpoint of the Amazon OpenSearch Service domain.
	DomainEndpoint *string `field:"required" json:"domainEndpoint" yaml:"domainEndpoint"`
}


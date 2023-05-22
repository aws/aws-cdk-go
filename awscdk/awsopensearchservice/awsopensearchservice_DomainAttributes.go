package awsopensearchservice


// Reference to an Amazon OpenSearch Service domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainAttributes := &DomainAttributes{
//   	DomainArn: jsii.String("domainArn"),
//   	DomainEndpoint: jsii.String("domainEndpoint"),
//   }
//
// Experimental.
type DomainAttributes struct {
	// The ARN of the Amazon OpenSearch Service domain.
	// Experimental.
	DomainArn *string `field:"required" json:"domainArn" yaml:"domainArn"`
	// The domain endpoint of the Amazon OpenSearch Service domain.
	// Experimental.
	DomainEndpoint *string `field:"required" json:"domainEndpoint" yaml:"domainEndpoint"`
}


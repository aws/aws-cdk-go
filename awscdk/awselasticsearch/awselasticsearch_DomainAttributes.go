package awselasticsearch


// Reference to an Elasticsearch domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainAttributes := &domainAttributes{
//   	domainArn: jsii.String("domainArn"),
//   	domainEndpoint: jsii.String("domainEndpoint"),
//   }
//
// Deprecated: use opensearchservice module instead.
type DomainAttributes struct {
	// The ARN of the Elasticsearch domain.
	// Deprecated: use opensearchservice module instead.
	DomainArn *string `field:"required" json:"domainArn" yaml:"domainArn"`
	// The domain endpoint of the Elasticsearch domain.
	// Deprecated: use opensearchservice module instead.
	DomainEndpoint *string `field:"required" json:"domainEndpoint" yaml:"domainEndpoint"`
}


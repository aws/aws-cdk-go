package awsappsync


// Properties for defining a `CfnDomainNameApiAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameApiAssociationProps := &cfnDomainNameApiAssociationProps{
//   	apiId: jsii.String("apiId"),
//   	domainName: jsii.String("domainName"),
//   }
//
type CfnDomainNameApiAssociationProps struct {
	// The API ID.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The domain name.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}


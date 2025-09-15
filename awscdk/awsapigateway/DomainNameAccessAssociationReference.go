package awsapigateway


// A reference to a DomainNameAccessAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainNameAccessAssociationReference := &DomainNameAccessAssociationReference{
//   	DomainNameAccessAssociationArn: jsii.String("domainNameAccessAssociationArn"),
//   }
//
type DomainNameAccessAssociationReference struct {
	// The DomainNameAccessAssociationArn of the DomainNameAccessAssociation resource.
	DomainNameAccessAssociationArn *string `field:"required" json:"domainNameAccessAssociationArn" yaml:"domainNameAccessAssociationArn"`
}


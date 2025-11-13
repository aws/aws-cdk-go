package interfacesawsappsync


// A reference to a DomainNameApiAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainNameApiAssociationReference := &DomainNameApiAssociationReference{
//   	ApiAssociationIdentifier: jsii.String("apiAssociationIdentifier"),
//   }
//
type DomainNameApiAssociationReference struct {
	// The ApiAssociationIdentifier of the DomainNameApiAssociation resource.
	ApiAssociationIdentifier *string `field:"required" json:"apiAssociationIdentifier" yaml:"apiAssociationIdentifier"`
}


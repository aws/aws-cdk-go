package interfacesawsappsync


// A reference to a SourceApiAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceApiAssociationReference := &SourceApiAssociationReference{
//   	AssociationArn: jsii.String("associationArn"),
//   }
//
type SourceApiAssociationReference struct {
	// The AssociationArn of the SourceApiAssociation resource.
	AssociationArn *string `field:"required" json:"associationArn" yaml:"associationArn"`
}


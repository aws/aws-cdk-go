package awsfsx


// A reference to a DataRepositoryAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataRepositoryAssociationReference := &DataRepositoryAssociationReference{
//   	AssociationId: jsii.String("associationId"),
//   }
//
type DataRepositoryAssociationReference struct {
	// The AssociationId of the DataRepositoryAssociation resource.
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
}


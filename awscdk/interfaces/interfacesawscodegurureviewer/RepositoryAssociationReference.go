package interfacesawscodegurureviewer


// A reference to a RepositoryAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryAssociationReference := &RepositoryAssociationReference{
//   	AssociationArn: jsii.String("associationArn"),
//   }
//
type RepositoryAssociationReference struct {
	// The AssociationArn of the RepositoryAssociation resource.
	AssociationArn *string `field:"required" json:"associationArn" yaml:"associationArn"`
}


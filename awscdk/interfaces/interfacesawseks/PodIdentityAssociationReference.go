package interfacesawseks


// A reference to a PodIdentityAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   podIdentityAssociationReference := &PodIdentityAssociationReference{
//   	AssociationArn: jsii.String("associationArn"),
//   }
//
type PodIdentityAssociationReference struct {
	// The AssociationArn of the PodIdentityAssociation resource.
	AssociationArn *string `field:"required" json:"associationArn" yaml:"associationArn"`
}


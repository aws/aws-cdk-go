package interfacesawsram


// A reference to a PrincipalAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   principalAssociationReference := &PrincipalAssociationReference{
//   	Principal: jsii.String("principal"),
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//   }
//
type PrincipalAssociationReference struct {
	// The Principal of the PrincipalAssociation resource.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The ResourceShareArn of the PrincipalAssociation resource.
	ResourceShareArn *string `field:"required" json:"resourceShareArn" yaml:"resourceShareArn"`
}


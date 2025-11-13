package interfacesawsroute53profiles


// A reference to a ProfileResourceAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profileResourceAssociationReference := &ProfileResourceAssociationReference{
//   	ProfileResourceAssociationId: jsii.String("profileResourceAssociationId"),
//   }
//
type ProfileResourceAssociationReference struct {
	// The Id of the ProfileResourceAssociation resource.
	ProfileResourceAssociationId *string `field:"required" json:"profileResourceAssociationId" yaml:"profileResourceAssociationId"`
}


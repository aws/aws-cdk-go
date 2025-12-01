package interfacesawsroute53profiles


// A reference to a ProfileAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profileAssociationReference := &ProfileAssociationReference{
//   	ProfileAssociationId: jsii.String("profileAssociationId"),
//   	ResourceId: jsii.String("resourceId"),
//   }
//
type ProfileAssociationReference struct {
	// The Id of the ProfileAssociation resource.
	ProfileAssociationId *string `field:"required" json:"profileAssociationId" yaml:"profileAssociationId"`
	// The ResourceId of the ProfileAssociation resource.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}


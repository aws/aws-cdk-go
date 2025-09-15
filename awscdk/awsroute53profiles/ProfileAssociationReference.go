package awsroute53profiles


// A reference to a ProfileAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profileAssociationReference := &ProfileAssociationReference{
//   	ProfileAssociationId: jsii.String("profileAssociationId"),
//   }
//
type ProfileAssociationReference struct {
	// The Id of the ProfileAssociation resource.
	ProfileAssociationId *string `field:"required" json:"profileAssociationId" yaml:"profileAssociationId"`
}


package interfacesawswisdom


// A reference to a ContentAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentAssociationReference := &ContentAssociationReference{
//   	ContentAssociationArn: jsii.String("contentAssociationArn"),
//   }
//
type ContentAssociationReference struct {
	// The ContentAssociationArn of the ContentAssociation resource.
	ContentAssociationArn *string `field:"required" json:"contentAssociationArn" yaml:"contentAssociationArn"`
}


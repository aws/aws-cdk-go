package awscognito


// A reference to a IdentityPoolRoleAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityPoolRoleAttachmentReference := &IdentityPoolRoleAttachmentReference{
//   	IdentityPoolRoleAttachmentId: jsii.String("identityPoolRoleAttachmentId"),
//   }
//
type IdentityPoolRoleAttachmentReference struct {
	// The Id of the IdentityPoolRoleAttachment resource.
	IdentityPoolRoleAttachmentId *string `field:"required" json:"identityPoolRoleAttachmentId" yaml:"identityPoolRoleAttachmentId"`
}


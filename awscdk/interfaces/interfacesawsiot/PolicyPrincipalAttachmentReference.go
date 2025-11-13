package interfacesawsiot


// A reference to a PolicyPrincipalAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyPrincipalAttachmentReference := &PolicyPrincipalAttachmentReference{
//   	PolicyPrincipalAttachmentId: jsii.String("policyPrincipalAttachmentId"),
//   }
//
type PolicyPrincipalAttachmentReference struct {
	// The Id of the PolicyPrincipalAttachment resource.
	PolicyPrincipalAttachmentId *string `field:"required" json:"policyPrincipalAttachmentId" yaml:"policyPrincipalAttachmentId"`
}


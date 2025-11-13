package interfacesawssecretsmanager


// A reference to a SecretTargetAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretTargetAttachmentReference := &SecretTargetAttachmentReference{
//   	SecretTargetAttachmentId: jsii.String("secretTargetAttachmentId"),
//   }
//
type SecretTargetAttachmentReference struct {
	// The Id of the SecretTargetAttachment resource.
	SecretTargetAttachmentId *string `field:"required" json:"secretTargetAttachmentId" yaml:"secretTargetAttachmentId"`
}


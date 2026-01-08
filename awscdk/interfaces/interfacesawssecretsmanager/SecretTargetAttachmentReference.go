package interfacesawssecretsmanager


// A reference to a SecretTargetAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretTargetAttachmentReference := &SecretTargetAttachmentReference{
//   	SecretId: jsii.String("secretId"),
//   }
//
type SecretTargetAttachmentReference struct {
	// The SecretId of the SecretTargetAttachment resource.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
}


package awssecretsmanager


// Construction properties for an AttachedSecret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//   var secretAttachmentTarget iSecretAttachmentTarget
//
//   secretTargetAttachmentProps := &secretTargetAttachmentProps{
//   	secret: secret,
//   	target: secretAttachmentTarget,
//   }
//
// Experimental.
type SecretTargetAttachmentProps struct {
	// The target to attach the secret to.
	// Experimental.
	Target ISecretAttachmentTarget `field:"required" json:"target" yaml:"target"`
	// The secret to attach to the target.
	// Experimental.
	Secret ISecret `field:"required" json:"secret" yaml:"secret"`
}


package awssecretsmanager


// Options to add a secret attachment to a secret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretAttachmentTarget iSecretAttachmentTarget
//
//   attachedSecretOptions := &AttachedSecretOptions{
//   	Target: secretAttachmentTarget,
//   }
//
// Experimental.
type AttachedSecretOptions struct {
	// The target to attach the secret to.
	// Experimental.
	Target ISecretAttachmentTarget `field:"required" json:"target" yaml:"target"`
}


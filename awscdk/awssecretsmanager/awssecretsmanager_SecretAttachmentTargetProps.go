package awssecretsmanager


// Attachment target specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretAttachmentTargetProps := &secretAttachmentTargetProps{
//   	targetId: jsii.String("targetId"),
//   	targetType: awscdk.Aws_secretsmanager.attachmentTargetType_RDS_DB_INSTANCE,
//   }
//
type SecretAttachmentTargetProps struct {
	// The id of the target to attach the secret to.
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// The type of the target to attach the secret to.
	TargetType AttachmentTargetType `field:"required" json:"targetType" yaml:"targetType"`
}


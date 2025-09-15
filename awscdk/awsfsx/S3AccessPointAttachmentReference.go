package awsfsx


// A reference to a S3AccessPointAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3AccessPointAttachmentReference := &S3AccessPointAttachmentReference{
//   	S3AccessPointAttachmentName: jsii.String("s3AccessPointAttachmentName"),
//   }
//
type S3AccessPointAttachmentReference struct {
	// The Name of the S3AccessPointAttachment resource.
	S3AccessPointAttachmentName *string `field:"required" json:"s3AccessPointAttachmentName" yaml:"s3AccessPointAttachmentName"`
}


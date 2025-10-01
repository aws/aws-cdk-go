package awsglobalaccelerator


// A reference to a CrossAccountAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossAccountAttachmentReference := &CrossAccountAttachmentReference{
//   	AttachmentArn: jsii.String("attachmentArn"),
//   }
//
type CrossAccountAttachmentReference struct {
	// The AttachmentArn of the CrossAccountAttachment resource.
	AttachmentArn *string `field:"required" json:"attachmentArn" yaml:"attachmentArn"`
}


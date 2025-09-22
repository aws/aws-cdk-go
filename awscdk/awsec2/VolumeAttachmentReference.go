package awsec2


// A reference to a VolumeAttachment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeAttachmentReference := &VolumeAttachmentReference{
//   	InstanceId: jsii.String("instanceId"),
//   	VolumeId: jsii.String("volumeId"),
//   }
//
type VolumeAttachmentReference struct {
	// The InstanceId of the VolumeAttachment resource.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The VolumeId of the VolumeAttachment resource.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
}


package awsec2


// Properties for defining a `CfnVolumeAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVolumeAttachmentProps := &cfnVolumeAttachmentProps{
//   	device: jsii.String("device"),
//   	instanceId: jsii.String("instanceId"),
//   	volumeId: jsii.String("volumeId"),
//   }
//
type CfnVolumeAttachmentProps struct {
	// The device name (for example, `/dev/sdh` or `xvdh` ).
	Device *string `field:"required" json:"device" yaml:"device"`
	// The ID of the instance to which the volume attaches.
	//
	// This value can be a reference to an [`AWS::EC2::Instance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource, or it can be the physical ID of an existing EC2 instance.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The ID of the Amazon EBS volume.
	//
	// The volume and instance must be within the same Availability Zone. This value can be a reference to an [`AWS::EC2::Volume`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html) resource, or it can be the volume ID of an existing Amazon EBS volume.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
}


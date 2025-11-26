package previewawsec2mixins


// Properties for CfnVolumeAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVolumeAttachmentMixinProps := &CfnVolumeAttachmentMixinProps{
//   	Device: jsii.String("device"),
//   	InstanceId: jsii.String("instanceId"),
//   	VolumeId: jsii.String("volumeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volumeattachment.html
//
type CfnVolumeAttachmentMixinProps struct {
	// The device name (for example, `/dev/sdh` or `xvdh` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volumeattachment.html#cfn-ec2-volumeattachment-device
	//
	Device *string `field:"optional" json:"device" yaml:"device"`
	// The ID of the instance to which the volume attaches.
	//
	// This value can be a reference to an [`AWS::EC2::Instance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource, or it can be the physical ID of an existing EC2 instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volumeattachment.html#cfn-ec2-volumeattachment-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ID of the Amazon EBS volume.
	//
	// The volume and instance must be within the same Availability Zone. This value can be a reference to an [`AWS::EC2::Volume`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html) resource, or it can be the volume ID of an existing Amazon EBS volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volumeattachment.html#cfn-ec2-volumeattachment-volumeid
	//
	VolumeId *string `field:"optional" json:"volumeId" yaml:"volumeId"`
}


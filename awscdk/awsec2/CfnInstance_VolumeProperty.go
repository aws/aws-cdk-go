package awsec2


// Specifies a volume to attach to an instance.
//
// `Volume` is an embedded property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeProperty := &VolumeProperty{
//   	Device: jsii.String("device"),
//   	VolumeId: jsii.String("volumeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-volume.html
//
type CfnInstance_VolumeProperty struct {
	// The device name (for example, `/dev/sdh` or `xvdh` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-volume.html#cfn-ec2-instance-volume-device
	//
	Device *string `field:"required" json:"device" yaml:"device"`
	// The ID of the EBS volume.
	//
	// The volume and instance must be within the same Availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-volume.html#cfn-ec2-instance-volume-volumeid
	//
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
}


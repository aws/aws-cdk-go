package awsec2


// Specifies a block device mapping for an instance.
//
// You must specify exactly one of the following properties: `VirtualName` , `Ebs` , or `NoDevice` .
//
// `BlockDeviceMapping` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
//
// > After the instance is running, you can modify only the `DeleteOnTermination` parameter for the attached volumes without interrupting the instance. Modifying any other parameter results in instance [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockDeviceMappingProperty := &BlockDeviceMappingProperty{
//   	DeviceName: jsii.String("deviceName"),
//
//   	// the properties below are optional
//   	Ebs: &EbsProperty{
//   		DeleteOnTermination: jsii.Boolean(false),
//   		Encrypted: jsii.Boolean(false),
//   		Iops: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		SnapshotId: jsii.String("snapshotId"),
//   		VolumeSize: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	NoDevice: &NoDeviceProperty{
//   	},
//   	VirtualName: jsii.String("virtualName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html
//
type CfnInstance_BlockDeviceMappingProperty struct {
	// The device name (for example, `/dev/sdh` or `xvdh` ).
	//
	// > After the instance is running, this parameter is used to specify the device name of the block device mapping to update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html#cfn-ec2-instance-blockdevicemapping-devicename
	//
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Parameters used to automatically set up EBS volumes when the instance is launched.
	//
	// > After the instance is running, you can modify only the `DeleteOnTermination` parameter for the attached volumes without interrupting the instance. Modifying any other parameter results in instance [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html#cfn-ec2-instance-blockdevicemapping-ebs
	//
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// To omit the device from the block device mapping, specify an empty string.
	//
	// > After the instance is running, modifying this parameter results in instance [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html#cfn-ec2-instance-blockdevicemapping-nodevice
	//
	NoDevice interface{} `field:"optional" json:"noDevice" yaml:"noDevice"`
	// The virtual device name ( `ephemeral` N).
	//
	// The name must be in the form `ephemeral` *X* where *X* is a number starting from zero (0). For example, an instance type with 2 available instance store volumes can specify mappings for `ephemeral0` and `ephemeral1` . The number of available instance store volumes depends on the instance type. After you connect to the instance, you must mount the volume.
	//
	// NVMe instance store volumes are automatically enumerated and assigned a device name. Including them in your block device mapping has no effect.
	//
	// *Constraints* : For M3 instances, you must specify instance store volumes in the block device mapping for the instance. When you launch an M3 instance, we ignore any instance store volumes specified in the block device mapping for the AMI.
	//
	// > After the instance is running, modifying this parameter results in instance [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html#cfn-ec2-instance-blockdevicemapping-virtualname
	//
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}


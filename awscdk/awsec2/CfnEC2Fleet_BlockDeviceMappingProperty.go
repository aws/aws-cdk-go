package awsec2


// Describes a block device mapping, which defines the EBS volumes and instance store volumes to attach to an instance at launch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockDeviceMappingProperty := &BlockDeviceMappingProperty{
//   	DeviceName: jsii.String("deviceName"),
//   	Ebs: &EbsBlockDeviceProperty{
//   		DeleteOnTermination: jsii.Boolean(false),
//   		Encrypted: jsii.Boolean(false),
//   		Iops: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		SnapshotId: jsii.String("snapshotId"),
//   		VolumeSize: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	NoDevice: jsii.String("noDevice"),
//   	VirtualName: jsii.String("virtualName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-blockdevicemapping.html
//
type CfnEC2Fleet_BlockDeviceMappingProperty struct {
	// The device name (for example, `/dev/sdh` or `xvdh` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-blockdevicemapping.html#cfn-ec2-ec2fleet-blockdevicemapping-devicename
	//
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Parameters used to automatically set up EBS volumes when the instance is launched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-blockdevicemapping.html#cfn-ec2-ec2fleet-blockdevicemapping-ebs
	//
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// To omit the device from the block device mapping, specify an empty string.
	//
	// When this property is specified, the device is removed from the block device mapping regardless of the assigned value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-blockdevicemapping.html#cfn-ec2-ec2fleet-blockdevicemapping-nodevice
	//
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// The virtual device name ( `ephemeral` N).
	//
	// Instance store volumes are numbered starting from 0. An instance type with 2 available instance store volumes can specify mappings for `ephemeral0` and `ephemeral1` . The number of available instance store volumes depends on the instance type. After you connect to the instance, you must mount the volume.
	//
	// NVMe instance store volumes are automatically enumerated and assigned a device name. Including them in your block device mapping has no effect.
	//
	// Constraints: For M3 instances, you must specify instance store volumes in the block device mapping for the instance. When you launch an M3 instance, we ignore any instance store volumes specified in the block device mapping for the AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-blockdevicemapping.html#cfn-ec2-ec2fleet-blockdevicemapping-virtualname
	//
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}


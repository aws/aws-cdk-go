package awsec2


// Specifies a block device mapping.
//
// You can specify `Ebs` or `VirtualName` , but not both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockDeviceMappingProperty := &blockDeviceMappingProperty{
//   	deviceName: jsii.String("deviceName"),
//
//   	// the properties below are optional
//   	ebs: &ebsBlockDeviceProperty{
//   		deleteOnTermination: jsii.Boolean(false),
//   		encrypted: jsii.Boolean(false),
//   		iops: jsii.Number(123),
//   		snapshotId: jsii.String("snapshotId"),
//   		volumeSize: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//   	},
//   	noDevice: jsii.String("noDevice"),
//   	virtualName: jsii.String("virtualName"),
//   }
//
type CfnSpotFleet_BlockDeviceMappingProperty struct {
	// The device name (for example, `/dev/sdh` or `xvdh` ).
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Parameters used to automatically set up EBS volumes when the instance is launched.
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// To omit the device from the block device mapping, specify an empty string.
	//
	// When this property is specified, the device is removed from the block device mapping regardless of the assigned value.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// The virtual device name ( `ephemeral` N).
	//
	// Instance store volumes are numbered starting from 0. An instance type with 2 available instance store volumes can specify mappings for `ephemeral0` and `ephemeral1` . The number of available instance store volumes depends on the instance type. After you connect to the instance, you must mount the volume.
	//
	// NVMe instance store volumes are automatically enumerated and assigned a device name. Including them in your block device mapping has no effect.
	//
	// Constraints: For M3 instances, you must specify instance store volumes in the block device mapping for the instance. When you launch an M3 instance, we ignore any instance store volumes specified in the block device mapping for the AMI.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}


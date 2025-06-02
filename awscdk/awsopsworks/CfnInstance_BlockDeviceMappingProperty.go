package awsopsworks


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockDeviceMappingProperty := &BlockDeviceMappingProperty{
//   	DeviceName: jsii.String("deviceName"),
//   	Ebs: &EbsBlockDeviceProperty{
//   		DeleteOnTermination: jsii.Boolean(false),
//   		Iops: jsii.Number(123),
//   		SnapshotId: jsii.String("snapshotId"),
//   		VolumeSize: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	NoDevice: jsii.String("noDevice"),
//   	VirtualName: jsii.String("virtualName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html
//
type CfnInstance_BlockDeviceMappingProperty struct {
	// The device name that is exposed to the instance, such as `/dev/sdh` .
	//
	// For the root device, you can use the explicit device name or you can set this parameter to `ROOT_DEVICE` and AWS OpsWorks Stacks will provide the correct device name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-devicename
	//
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// An `EBSBlockDevice` that defines how to configure an Amazon EBS volume when the instance is launched.
	//
	// You can specify either the `VirtualName` or `Ebs` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-ebs
	//
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// Suppresses the specified device included in the AMI's block device mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-nodevice
	//
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// The virtual device name.
	//
	// For more information, see [BlockDeviceMapping](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_BlockDeviceMapping.html) . You can specify either the `VirtualName` or `Ebs` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html#cfn-opsworks-instance-blockdevicemapping-virtualname
	//
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}


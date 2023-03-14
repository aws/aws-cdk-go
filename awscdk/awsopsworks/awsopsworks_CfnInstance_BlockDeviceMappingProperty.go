package awsopsworks


// Describes a block device mapping.
//
// This data type maps directly to the Amazon EC2 [BlockDeviceMapping](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_BlockDeviceMapping.html) data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockDeviceMappingProperty := &blockDeviceMappingProperty{
//   	deviceName: jsii.String("deviceName"),
//   	ebs: &ebsBlockDeviceProperty{
//   		deleteOnTermination: jsii.Boolean(false),
//   		iops: jsii.Number(123),
//   		snapshotId: jsii.String("snapshotId"),
//   		volumeSize: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//   	},
//   	noDevice: jsii.String("noDevice"),
//   	virtualName: jsii.String("virtualName"),
//   }
//
type CfnInstance_BlockDeviceMappingProperty struct {
	// The device name that is exposed to the instance, such as `/dev/sdh` .
	//
	// For the root device, you can use the explicit device name or you can set this parameter to `ROOT_DEVICE` and AWS OpsWorks Stacks will provide the correct device name.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// An `EBSBlockDevice` that defines how to configure an Amazon EBS volume when the instance is launched.
	//
	// You can specify either the `VirtualName` or `Ebs` , but not both.
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// Suppresses the specified device included in the AMI's block device mapping.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// The virtual device name.
	//
	// For more information, see [BlockDeviceMapping](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_BlockDeviceMapping.html) . You can specify either the `VirtualName` or `Ebs` , but not both.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}


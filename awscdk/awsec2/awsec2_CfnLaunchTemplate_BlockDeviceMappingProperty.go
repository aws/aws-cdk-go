package awsec2


// Information about a block device mapping for an Amazon EC2 launch template.
//
// `BlockDeviceMapping` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockDeviceMappingProperty := &blockDeviceMappingProperty{
//   	deviceName: jsii.String("deviceName"),
//   	ebs: &ebsProperty{
//   		deleteOnTermination: jsii.Boolean(false),
//   		encrypted: jsii.Boolean(false),
//   		iops: jsii.Number(123),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   		snapshotId: jsii.String("snapshotId"),
//   		throughput: jsii.Number(123),
//   		volumeSize: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//   	},
//   	noDevice: jsii.String("noDevice"),
//   	virtualName: jsii.String("virtualName"),
//   }
//
type CfnLaunchTemplate_BlockDeviceMappingProperty struct {
	// The device name (for example, /dev/sdh or xvdh).
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Parameters used to automatically set up EBS volumes when the instance is launched.
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// To omit the device from the block device mapping, specify an empty string.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// The virtual device name (ephemeralN).
	//
	// Instance store volumes are numbered starting from 0. An instance type with 2 available instance store volumes can specify mappings for ephemeral0 and ephemeral1. The number of available instance store volumes depends on the instance type. After you connect to the instance, you must mount the volume.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}


package awsautoscaling


// `BlockDeviceMapping` specifies a block device mapping for the `BlockDeviceMappings` property of the [AWS::AutoScaling::LaunchConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html) resource.
//
// Each instance that is launched has an associated root device volume, either an Amazon EBS volume or an instance store volume. You can use block device mappings to specify additional EBS volumes or instance store volumes to attach to an instance when it is launched.
//
// For more information, see [Example block device mapping](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html#block-device-mapping-ex) in the *Amazon EC2 User Guide for Linux Instances* .
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
//   	ebs: &blockDeviceProperty{
//   		deleteOnTermination: jsii.Boolean(false),
//   		encrypted: jsii.Boolean(false),
//   		iops: jsii.Number(123),
//   		snapshotId: jsii.String("snapshotId"),
//   		throughput: jsii.Number(123),
//   		volumeSize: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//   	},
//   	noDevice: jsii.Boolean(false),
//   	virtualName: jsii.String("virtualName"),
//   }
//
type CfnLaunchConfiguration_BlockDeviceMappingProperty struct {
	// The device name assigned to the volume (for example, `/dev/sdh` or `xvdh` ).
	//
	// For more information, see [Device naming on Linux instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// > To define a block device mapping, set the device name and exactly one of the following properties: `Ebs` , `NoDevice` , or `VirtualName` .
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Information to attach an EBS volume to an instance at launch.
	Ebs interface{} `field:"optional" json:"ebs" yaml:"ebs"`
	// Setting this value to `true` prevents a volume that is included in the block device mapping of the AMI from being mapped to the specified device name at launch.
	//
	// If `NoDevice` is `true` for the root device, instances might fail the EC2 health check. In that case, Amazon EC2 Auto Scaling launches replacement instances.
	NoDevice interface{} `field:"optional" json:"noDevice" yaml:"noDevice"`
	// The name of the instance store volume (virtual device) to attach to an instance at launch.
	//
	// The name must be in the form ephemeral *X* where *X* is a number starting from zero (0), for example, `ephemeral0` .
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}


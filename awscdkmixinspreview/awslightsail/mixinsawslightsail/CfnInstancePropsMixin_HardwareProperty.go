package mixinsawslightsail


// `Hardware` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the hardware properties for the instance, such as the vCPU count, attached disks, and amount of RAM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hardwareProperty := &HardwareProperty{
//   	CpuCount: jsii.Number(123),
//   	Disks: []interface{}{
//   		&DiskProperty{
//   			AttachedTo: jsii.String("attachedTo"),
//   			AttachmentState: jsii.String("attachmentState"),
//   			DiskName: jsii.String("diskName"),
//   			Iops: jsii.Number(123),
//   			IsSystemDisk: jsii.Boolean(false),
//   			Path: jsii.String("path"),
//   			SizeInGb: jsii.String("sizeInGb"),
//   		},
//   	},
//   	RamSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-hardware.html
//
type CfnInstancePropsMixin_HardwareProperty struct {
	// The number of vCPUs the instance has.
	//
	// > The `CpuCount` property is read-only and should not be specified in a create instance or update instance request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-hardware.html#cfn-lightsail-instance-hardware-cpucount
	//
	CpuCount *float64 `field:"optional" json:"cpuCount" yaml:"cpuCount"`
	// The disks attached to the instance.
	//
	// The instance restarts when performing an attach disk or detach disk request. This resets the public IP address of your instance if a static IP isn't attached to it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-hardware.html#cfn-lightsail-instance-hardware-disks
	//
	Disks interface{} `field:"optional" json:"disks" yaml:"disks"`
	// The amount of RAM in GB on the instance (for example, `1.0` ).
	//
	// > The `RamSizeInGb` property is read-only and should not be specified in a create instance or update instance request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-hardware.html#cfn-lightsail-instance-hardware-ramsizeingb
	//
	RamSizeInGb *float64 `field:"optional" json:"ramSizeInGb" yaml:"ramSizeInGb"`
}


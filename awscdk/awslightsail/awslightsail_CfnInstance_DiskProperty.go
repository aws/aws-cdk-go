package awslightsail


// `Disk` is a property of the [Hardware](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-hardware.html) property. It describes a disk attached to an instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   diskProperty := &diskProperty{
//   	diskName: jsii.String("diskName"),
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	attachedTo: jsii.String("attachedTo"),
//   	attachmentState: jsii.String("attachmentState"),
//   	iops: jsii.Number(123),
//   	isSystemDisk: jsii.Boolean(false),
//   	sizeInGb: jsii.String("sizeInGb"),
//   }
//
type CfnInstance_DiskProperty struct {
	// The unique name of the disk.
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// The disk path.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The resources to which the disk is attached.
	AttachedTo *string `field:"optional" json:"attachedTo" yaml:"attachedTo"`
	// (Deprecated) The attachment state of the disk.
	//
	// > In releases prior to November 14, 2017, this parameter returned `attached` for system disks in the API response. It is now deprecated, but still included in the response. Use `isAttached` instead.
	AttachmentState *string `field:"optional" json:"attachmentState" yaml:"attachmentState"`
	// The input/output operations per second (IOPS) of the disk.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// A Boolean value indicating whether this disk is a system disk (has an operating system loaded on it).
	IsSystemDisk interface{} `field:"optional" json:"isSystemDisk" yaml:"isSystemDisk"`
	// The size of the disk in GB.
	SizeInGb *string `field:"optional" json:"sizeInGb" yaml:"sizeInGb"`
}


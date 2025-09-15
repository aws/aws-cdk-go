package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Block device options for an EBS volume.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//
//   	// ...
//
//   	BlockDevices: []blockDevice{
//   		&blockDevice{
//   			DeviceName: jsii.String("/dev/sda1"),
//   			Volume: ec2.BlockDeviceVolume_Ebs(jsii.Number(100), &EbsDeviceOptions{
//   				VolumeType: ec2.EbsDeviceVolumeType_GP3,
//   				Throughput: jsii.Number(250),
//   			}),
//   		},
//   	},
//   })
//
type EbsDeviceOptions struct {
	// Indicates whether to delete the volume when the instance is terminated.
	// Default: - true for Amazon EC2 Auto Scaling, false otherwise (e.g. EBS)
	//
	DeleteOnTermination *bool `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for `volumeType`: `EbsDeviceVolumeType.IO1`
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	// Default: - none, required for `EbsDeviceVolumeType.IO1`
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The throughput to provision for a `gp3` volume.
	//
	// Valid Range: Minimum value of 125. Maximum value of 1000.
	//
	// `gp3` volumes deliver a consistent baseline throughput performance of 125 MiB/s.
	// You can provision additional throughput for an additional cost at a ratio of 0.25 MiB/s per provisioned IOPS.
	// See: https://docs.aws.amazon.com/ebs/latest/userguide/general-purpose.html#gp3-performance
	//
	// Default: - 125 MiB/s.
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	// Default: `EbsDeviceVolumeType.GENERAL_PURPOSE_SSD` or `EbsDeviceVolumeType.GENERAL_PURPOSE_SSD_GP3` if
	// `@aws-cdk/aws-ec2:ebsDefaultGp3Volume` is enabled.
	//
	VolumeType EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
	// Specifies whether the EBS volume is encrypted.
	//
	// Encrypted EBS volumes can only be attached to instances that support Amazon EBS encryption.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances
	//
	// Default: false.
	//
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The ARN of the AWS Key Management Service (AWS KMS) CMK used for encryption.
	//
	// You have to ensure that the KMS CMK has the correct permissions to be used by the service launching the ec2 instances.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#ebs-encryption-requirements
	//
	// Default: - If encrypted is true, the default aws/ebs KMS key will be used.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}


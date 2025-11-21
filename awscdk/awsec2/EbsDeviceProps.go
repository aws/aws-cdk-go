package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties of an EBS block device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key Key
//
//   ebsDeviceProps := &EbsDeviceProps{
//   	DeleteOnTermination: jsii.Boolean(false),
//   	Encrypted: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	KmsKey: key,
//   	SnapshotId: jsii.String("snapshotId"),
//   	Throughput: jsii.Number(123),
//   	VolumeSize: jsii.Number(123),
//   	VolumeType: awscdk.Aws_ec2.EbsDeviceVolumeType_STANDARD,
//   }
//
type EbsDeviceProps struct {
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
	// Valid Range: Minimum value of 125. Maximum value of 2000.
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
	// The volume size, in Gibibytes (GiB).
	//
	// If you specify volumeSize, it must be equal or greater than the size of the snapshot.
	// Default: - The snapshot size.
	//
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
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
	// The snapshot ID of the volume to use.
	// Default: - No snapshot will be used.
	//
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
}


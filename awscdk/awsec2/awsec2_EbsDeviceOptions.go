package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Block device options for an EBS volume.
//
// Example:
//   host := ec2.NewBastionHostLinux(this, jsii.String("BastionHost"), &bastionHostLinuxProps{
//   	vpc: vpc,
//   	blockDevices: []blockDevice{
//   		&blockDevice{
//   			deviceName: jsii.String("EBSBastionHost"),
//   			volume: ec2.blockDeviceVolume.ebs(jsii.Number(10), &ebsDeviceOptions{
//   				encrypted: jsii.Boolean(true),
//   			}),
//   		},
//   	},
//   })
//
type EbsDeviceOptions struct {
	// Indicates whether to delete the volume when the instance is terminated.
	DeleteOnTermination *bool `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for {@link volumeType}: {@link EbsDeviceVolumeType.IO1}
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	VolumeType EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
	// Specifies whether the EBS volume is encrypted.
	//
	// Encrypted EBS volumes can only be attached to instances that support Amazon EBS encryption.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances
	//
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The ARN of the AWS Key Management Service (AWS KMS) CMK used for encryption.
	//
	// You have to ensure that the KMS CMK has the correct permissions to be used by the service launching the ec2 instances.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#ebs-encryption-requirements
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}


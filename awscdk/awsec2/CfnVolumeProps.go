package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVolume`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVolumeProps := &CfnVolumeProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//
//   	// the properties below are optional
//   	AutoEnableIo: jsii.Boolean(false),
//   	Encrypted: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MultiAttachEnabled: jsii.Boolean(false),
//   	OutpostArn: jsii.String("outpostArn"),
//   	Size: jsii.Number(123),
//   	SnapshotId: jsii.String("snapshotId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Throughput: jsii.Number(123),
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html
//
type CfnVolumeProps struct {
	// The ID of the Availability Zone in which to create the volume.
	//
	// For example, `us-east-1a` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-availabilityzone
	//
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Indicates whether the volume is auto-enabled for I/O operations.
	//
	// By default, Amazon EBS disables I/O to the volume from attached EC2 instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-autoenableio
	//
	AutoEnableIo interface{} `field:"optional" json:"autoEnableIo" yaml:"autoEnableIo"`
	// Indicates whether the volume should be encrypted.
	//
	// The effect of setting the encryption state to `true` depends on the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Encryption by default](https://docs.aws.amazon.com/ebs/latest/userguide/work-with-ebs-encr.html#encryption-by-default) in the *Amazon EBS User Guide* .
	//
	// Encrypted Amazon EBS volumes must be attached to instances that support Amazon EBS encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption-requirements.html#ebs-encryption_supported_instances) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-encrypted
	//
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS).
	//
	// For `gp3` , `io1` , and `io2` volumes, this represents the number of IOPS that are provisioned for the volume. For `gp2` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.
	//
	// The following are the supported values for each volume type:
	//
	// - `gp3` : 3,000 - 16,000 IOPS
	// - `io1` : 100 - 64,000 IOPS
	// - `io2` : 100 - 256,000 IOPS
	//
	// For `io2` volumes, you can achieve up to 256,000 IOPS on [instances built on the Nitro System](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances) . On other instances, you can achieve performance up to 32,000 IOPS.
	//
	// This parameter is required for `io1` and `io2` volumes. The default for `gp3` volumes is 3,000 IOPS. This parameter is not supported for `gp2` , `st1` , `sc1` , or `standard` volumes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The identifier of the AWS KMS key to use for Amazon EBS encryption.
	//
	// If `KmsKeyId` is specified, the encrypted state must be `true` .
	//
	// If you omit this property and your account is enabled for encryption by default, or *Encrypted* is set to `true` , then the volume is encrypted using the default key specified for your account. If your account does not have a default key, then the volume is encrypted using the AWS managed key .
	//
	// Alternatively, if you want to specify a different key, you can specify one of the following:
	//
	// - Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.
	// - Key alias. Specify the alias for the key, prefixed with `alias/` . For example, for a key with the alias `my_cmk` , use `alias/my_cmk` . Or to specify the AWS managed key , use `alias/aws/ebs` .
	// - Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	// - Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Indicates whether Amazon EBS Multi-Attach is enabled.
	//
	// AWS CloudFormation does not currently support updating a single-attach volume to be multi-attach enabled, updating a multi-attach enabled volume to be single-attach, or updating the size or number of I/O operations per second (IOPS) of a multi-attach enabled volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-multiattachenabled
	//
	MultiAttachEnabled interface{} `field:"optional" json:"multiAttachEnabled" yaml:"multiAttachEnabled"`
	// The Amazon Resource Name (ARN) of the Outpost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-outpostarn
	//
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// The size of the volume, in GiBs.
	//
	// You must specify either a snapshot ID or a volume size. If you specify a snapshot, the default is the snapshot size. You can specify a volume size that is equal to or larger than the snapshot size.
	//
	// The following are the supported volumes sizes for each volume type:
	//
	// - `gp2` and `gp3` : 1 - 16,384 GiB
	// - `io1` : 4 - 16,384 GiB
	// - `io2` : 4 - 65,536 GiB
	// - `st1` and `sc1` : 125 - 16,384 GiB
	// - `standard` : 1 - 1024 GiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-size
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The snapshot from which to create the volume.
	//
	// You must specify either a snapshot ID or a volume size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-snapshotid
	//
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The tags to apply to the volume during creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The throughput to provision for a volume, with a maximum of 1,000 MiB/s.
	//
	// This parameter is valid only for `gp3` volumes. The default value is 125.
	//
	// Valid Range: Minimum value of 125. Maximum value of 1000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-throughput
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The volume type. This parameter can be one of the following values:.
	//
	// - General Purpose SSD: `gp2` | `gp3`
	// - Provisioned IOPS SSD: `io1` | `io2`
	// - Throughput Optimized HDD: `st1`
	// - Cold HDD: `sc1`
	// - Magnetic: `standard`
	//
	// For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-volume-types.html) .
	//
	// Default: `gp2`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-volume.html#cfn-ec2-volume-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}


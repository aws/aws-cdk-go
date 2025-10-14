package awsec2


// Describes a block device for an EBS volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsBlockDeviceProperty := &EbsBlockDeviceProperty{
//   	DeleteOnTermination: jsii.Boolean(false),
//   	Encrypted: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SnapshotId: jsii.String("snapshotId"),
//   	VolumeSize: jsii.Number(123),
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ebsblockdevice.html
//
type CfnEC2Fleet_EbsBlockDeviceProperty struct {
	// Indicates whether the EBS volume is deleted on instance termination.
	//
	// For more information, see [Preserving Amazon EBS volumes on instance termination](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#preserving-volumes-on-termination) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ebsblockdevice.html#cfn-ec2-ec2fleet-ebsblockdevice-deleteontermination
	//
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Indicates whether the encryption state of an EBS volume is changed while being restored from a backing snapshot.
	//
	// The effect of setting the encryption state to `true` depends on the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Amazon EBS encryption](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption.html#encryption-parameters) in the *Amazon EBS User Guide* .
	//
	// In no case can you remove encryption from an encrypted volume.
	//
	// Encrypted volumes can only be attached to instances that support Amazon EBS encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption-requirements.html#ebs-encryption_supported_instances) .
	//
	// - If you are creating a block device mapping for a *new (empty) volume* , you can include this parameter, and specify either `true` for an encrypted volume, or `false` for an unencrypted volume. If you omit this parameter, it defaults to `false` (unencrypted).
	// - If you are creating a block device mapping from an *existing encrypted or unencrypted snapshot* , you must omit this parameter. If you include this parameter, the request will fail, regardless of the value that you specify.
	// - If you are creating a block device mapping from an *existing unencrypted volume* , you can include this parameter, but you must specify `false` . If you specify `true` , the request will fail. In this case, we recommend that you omit the parameter.
	// - If you are creating a block device mapping from an *existing encrypted volume* , you can include this parameter, and specify either `true` or `false` . However, if you specify `false` , the parameter is ignored and the block device mapping is always encrypted. In this case, we recommend that you omit the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ebsblockdevice.html#cfn-ec2-ec2fleet-ebsblockdevice-encrypted
	//
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS).
	//
	// For `gp3` , `io1` , and `io2` volumes, this represents the number of IOPS that are provisioned for the volume. For `gp2` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.
	//
	// The following are the supported values for each volume type:
	//
	// - `gp3` : 3,000 - 80,000 IOPS
	// - `io1` : 100 - 64,000 IOPS
	// - `io2` : 100 - 256,000 IOPS
	//
	// For `io2` volumes, you can achieve up to 256,000 IOPS on [instances built on the Nitro System](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances) . On other instances, you can achieve performance up to 32,000 IOPS.
	//
	// This parameter is required for `io1` and `io2` volumes. The default for `gp3` volumes is 3,000 IOPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ebsblockdevice.html#cfn-ec2-ec2fleet-ebsblockdevice-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Identifier (key ID, key alias, key ARN, or alias ARN) of the customer managed KMS key to use for EBS encryption.
	//
	// This parameter is only supported on `BlockDeviceMapping` objects called by [RunInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html) , [RequestSpotFleet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RequestSpotFleet.html) , and [RequestSpotInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RequestSpotInstances.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ebsblockdevice.html#cfn-ec2-ec2fleet-ebsblockdevice-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The ID of the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ebsblockdevice.html#cfn-ec2-ec2fleet-ebsblockdevice-snapshotid
	//
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The size of the volume, in GiBs.
	//
	// You must specify either a snapshot ID or a volume size. If you specify a snapshot, the default is the snapshot size. You can specify a volume size that is equal to or larger than the snapshot size.
	//
	// The following are the supported sizes for each volume type:
	//
	// - `gp2` : 1 - 16,384 GiB
	// - `gp3` : 1 - 65,536 GiB
	// - `io1` : 4 - 16,384 GiB
	// - `io2` : 4 - 65,536 GiB
	// - `st1` and `sc1` : 125 - 16,384 GiB
	// - `standard` : 1 - 1024 GiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ebsblockdevice.html#cfn-ec2-ec2fleet-ebsblockdevice-volumesize
	//
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The volume type.
	//
	// For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-volume-types.html) in the *Amazon EBS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-ebsblockdevice.html#cfn-ec2-ec2fleet-ebsblockdevice-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}


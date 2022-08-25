package awsautoscaling


// `BlockDevice` is a property of the `EBS` property of the [AWS::AutoScaling::LaunchConfiguration BlockDeviceMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html) property type that describes an Amazon EBS volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockDeviceProperty := &blockDeviceProperty{
//   	deleteOnTermination: jsii.Boolean(false),
//   	encrypted: jsii.Boolean(false),
//   	iops: jsii.Number(123),
//   	snapshotId: jsii.String("snapshotId"),
//   	throughput: jsii.Number(123),
//   	volumeSize: jsii.Number(123),
//   	volumeType: jsii.String("volumeType"),
//   }
//
type CfnLaunchConfiguration_BlockDeviceProperty struct {
	// Indicates whether the volume is deleted on instance termination.
	//
	// For Amazon EC2 Auto Scaling, the default value is `true` .
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Specifies whether the volume should be encrypted.
	//
	// Encrypted EBS volumes can only be attached to instances that support Amazon EBS encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances) . If your AMI uses encrypted volumes, you can also only launch it on supported instance types.
	//
	// > If you are creating a volume from a snapshot, you cannot create an unencrypted volume from an encrypted snapshot. Also, you cannot specify a KMS key ID when using a launch configuration.
	// >
	// > If you enable encryption by default, the EBS volumes that you create are always encrypted, either using the AWS managed KMS key or a customer-managed KMS key, regardless of whether the snapshot was encrypted.
	// >
	// > For more information, see [Use AWS KMS keys to encrypt Amazon EBS volumes](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-data-protection.html#encryption) in the *Amazon EC2 Auto Scaling User Guide* .
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of input/output (I/O) operations per second (IOPS) to provision for the volume.
	//
	// For `gp3` and `io1` volumes, this represents the number of IOPS that are provisioned for the volume. For `gp2` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.
	//
	// The following are the supported values for each volume type:
	//
	// - `gp3` : 3,000-16,000 IOPS
	// - `io1` : 100-64,000 IOPS
	//
	// For `io1` volumes, we guarantee 64,000 IOPS only for [Instances built on the Nitro System](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances) . Other instance families guarantee performance up to 32,000 IOPS.
	//
	// `Iops` is supported when the volume type is `gp3` or `io1` and required only when the volume type is `io1` . (Not used with `standard` , `gp2` , `st1` , or `sc1` volumes.)
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The snapshot ID of the volume to use.
	//
	// You must specify either a `VolumeSize` or a `SnapshotId` .
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The throughput (MiBps) to provision for a `gp3` volume.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The volume size, in GiBs. The following are the supported volumes sizes for each volume type:.
	//
	// - `gp2` and `gp3` : 1-16,384
	// - `io1` : 4-16,384
	// - `st1` and `sc1` : 125-16,384
	// - `standard` : 1-1,024
	//
	// You must specify either a `SnapshotId` or a `VolumeSize` . If you specify both `SnapshotId` and `VolumeSize` , the volume size must be equal or greater than the size of the snapshot.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The volume type.
	//
	// For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// Valid values: `standard` | `io1` | `gp2` | `st1` | `sc1` | `gp3`.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}


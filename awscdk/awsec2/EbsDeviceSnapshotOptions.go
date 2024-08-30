package awsec2


// Block device options for an EBS volume created from a snapshot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsDeviceSnapshotOptions := &EbsDeviceSnapshotOptions{
//   	DeleteOnTermination: jsii.Boolean(false),
//   	Iops: jsii.Number(123),
//   	Throughput: jsii.Number(123),
//   	VolumeSize: jsii.Number(123),
//   	VolumeType: awscdk.Aws_ec2.EbsDeviceVolumeType_STANDARD,
//   }
//
type EbsDeviceSnapshotOptions struct {
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
	// The volume size, in Gibibytes (GiB).
	//
	// If you specify volumeSize, it must be equal or greater than the size of the snapshot.
	// Default: - The snapshot size.
	//
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}


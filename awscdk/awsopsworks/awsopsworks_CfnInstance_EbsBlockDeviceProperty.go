package awsopsworks


// Describes an Amazon EBS volume.
//
// This data type maps directly to the Amazon EC2 [EbsBlockDevice](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EbsBlockDevice.html) data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsBlockDeviceProperty := &ebsBlockDeviceProperty{
//   	deleteOnTermination: jsii.Boolean(false),
//   	iops: jsii.Number(123),
//   	snapshotId: jsii.String("snapshotId"),
//   	volumeSize: jsii.Number(123),
//   	volumeType: jsii.String("volumeType"),
//   }
//
type CfnInstance_EbsBlockDeviceProperty struct {
	// Whether the volume is deleted on instance termination.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	//
	// For more information, see [EbsBlockDevice](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EbsBlockDevice.html) .
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The snapshot ID.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The volume size, in GiB.
	//
	// For more information, see [EbsBlockDevice](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EbsBlockDevice.html) .
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The volume type.
	//
	// `gp2` for General Purpose (SSD) volumes, `io1` for Provisioned IOPS (SSD) volumes, `st1` for Throughput Optimized hard disk drives (HDD), `sc1` for Cold HDD,and `standard` for Magnetic volumes.
	//
	// If you specify the `io1` volume type, you must also specify a value for the `Iops` attribute. The maximum ratio of provisioned IOPS to requested volume size (in GiB) is 50:1. AWS uses the default volume size (in GiB) specified in the AMI attributes to set IOPS to 50 x (volume size).
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}


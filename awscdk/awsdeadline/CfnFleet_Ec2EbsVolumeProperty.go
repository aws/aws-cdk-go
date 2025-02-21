package awsdeadline


// Specifies the EBS volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ec2EbsVolumeProperty := &Ec2EbsVolumeProperty{
//   	Iops: jsii.Number(123),
//   	SizeGiB: jsii.Number(123),
//   	ThroughputMiB: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-ec2ebsvolume.html
//
type CfnFleet_Ec2EbsVolumeProperty struct {
	// The IOPS per volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-ec2ebsvolume.html#cfn-deadline-fleet-ec2ebsvolume-iops
	//
	// Default: - 3000.
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The EBS volume size in GiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-ec2ebsvolume.html#cfn-deadline-fleet-ec2ebsvolume-sizegib
	//
	// Default: - 250.
	//
	SizeGiB *float64 `field:"optional" json:"sizeGiB" yaml:"sizeGiB"`
	// The throughput per volume in MiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-ec2ebsvolume.html#cfn-deadline-fleet-ec2ebsvolume-throughputmib
	//
	// Default: - 125.
	//
	ThroughputMiB *float64 `field:"optional" json:"throughputMiB" yaml:"throughputMiB"`
}


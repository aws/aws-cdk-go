package awsemr


// `VolumeSpecification` is a subproperty of the `EbsBlockDeviceConfig` property type.
//
// `VolumeSecification` determines the volume type, IOPS, and size (GiB) for EBS volumes attached to EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeSpecificationProperty := &volumeSpecificationProperty{
//   	sizeInGb: jsii.Number(123),
//   	volumeType: jsii.String("volumeType"),
//
//   	// the properties below are optional
//   	iops: jsii.Number(123),
//   }
//
type CfnInstanceGroupConfig_VolumeSpecificationProperty struct {
	// The volume size, in gibibytes (GiB).
	//
	// This can be a number from 1 - 1024. If the volume type is EBS-optimized, the minimum value is 10.
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// The volume type.
	//
	// Volume types supported are gp2, io1, and standard.
	VolumeType *string `field:"required" json:"volumeType" yaml:"volumeType"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
}


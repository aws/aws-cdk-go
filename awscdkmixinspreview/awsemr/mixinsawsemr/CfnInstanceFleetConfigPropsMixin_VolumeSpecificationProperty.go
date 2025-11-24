package mixinsawsemr


// `VolumeSpecification` is a subproperty of the `EbsBlockDeviceConfig` property type.
//
// `VolumeSecification` determines the volume type, IOPS, and size (GiB) for EBS volumes attached to EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   volumeSpecificationProperty := &VolumeSpecificationProperty{
//   	Iops: jsii.Number(123),
//   	SizeInGb: jsii.Number(123),
//   	Throughput: jsii.Number(123),
//   	VolumeType: jsii.String("volumeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-volumespecification.html
//
type CfnInstanceFleetConfigPropsMixin_VolumeSpecificationProperty struct {
	// The number of I/O operations per second (IOPS) that the volume supports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-volumespecification.html#cfn-emr-instancefleetconfig-volumespecification-iops
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The volume size, in gibibytes (GiB).
	//
	// This can be a number from 1 - 1024. If the volume type is EBS-optimized, the minimum value is 10.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-volumespecification.html#cfn-emr-instancefleetconfig-volumespecification-sizeingb
	//
	SizeInGb *float64 `field:"optional" json:"sizeInGb" yaml:"sizeInGb"`
	// The throughput, in mebibyte per second (MiB/s).
	//
	// This optional parameter can be a number from 125 - 1000 and is valid only for gp3 volumes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-volumespecification.html#cfn-emr-instancefleetconfig-volumespecification-throughput
	//
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The volume type.
	//
	// Volume types supported are gp3, gp2, io1, st1, sc1, and standard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-volumespecification.html#cfn-emr-instancefleetconfig-volumespecification-volumetype
	//
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}


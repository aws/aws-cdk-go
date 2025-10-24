package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// EBS volume specifications such as volume type, IOPS, and size (GiB) that will be requested for the EBS volume attached to an EC2 instance in the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var size Size
//
//   volumeSpecificationProperty := &VolumeSpecificationProperty{
//   	VolumeSize: size,
//   	VolumeType: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.EbsBlockDeviceVolumeType_GP3,
//
//   	// the properties below are optional
//   	Iops: jsii.Number(123),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_VolumeSpecification.html
//
type EmrCreateCluster_VolumeSpecificationProperty struct {
	// The volume size.
	//
	// If the volume type is EBS-optimized, the minimum value is 10GiB.
	// Maximum size is 1TiB.
	VolumeSize awscdk.Size `field:"required" json:"volumeSize" yaml:"volumeSize"`
	// The volume type.
	//
	// Volume types supported are gp2, io1, standard.
	VolumeType EmrCreateCluster_EbsBlockDeviceVolumeType `field:"required" json:"volumeType" yaml:"volumeType"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	// Default: - EMR selected default.
	//
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
}


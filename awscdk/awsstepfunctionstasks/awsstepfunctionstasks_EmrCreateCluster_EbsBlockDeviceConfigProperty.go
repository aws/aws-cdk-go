package awsstepfunctionstasks


// Configuration of requested EBS block device associated with the instance group with count of volumes that will be associated to every instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var size size
//
//   ebsBlockDeviceConfigProperty := &ebsBlockDeviceConfigProperty{
//   	volumeSpecification: &volumeSpecificationProperty{
//   		volumeSize: size,
//   		volumeType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.ebsBlockDeviceVolumeType_GP2,
//
//   		// the properties below are optional
//   		iops: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	volumesPerInstance: jsii.Number(123),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_EbsBlockDeviceConfig.html
//
type EmrCreateCluster_EbsBlockDeviceConfigProperty struct {
	// EBS volume specifications such as volume type, IOPS, and size (GiB) that will be requested for the EBS volume attached to an EC2 instance in the cluster.
	VolumeSpecification *EmrCreateCluster_VolumeSpecificationProperty `field:"required" json:"volumeSpecification" yaml:"volumeSpecification"`
	// Number of EBS volumes with a specific volume configuration that will be associated with every instance in the instance group.
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}


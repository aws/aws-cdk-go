package awsemr


// `EbsBlockDeviceConfig` is a subproperty of the `EbsConfiguration` property type.
//
// `EbsBlockDeviceConfig` defines the number and type of EBS volumes to associate with all EC2 instances in an EMR cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsBlockDeviceConfigProperty := &EbsBlockDeviceConfigProperty{
//   	VolumeSpecification: &VolumeSpecificationProperty{
//   		SizeInGb: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//
//   		// the properties below are optional
//   		Iops: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	VolumesPerInstance: jsii.Number(123),
//   }
//
type CfnCluster_EbsBlockDeviceConfigProperty struct {
	// EBS volume specifications such as volume type, IOPS, size (GiB) and throughput (MiB/s) that are requested for the EBS volume attached to an EC2 instance in the cluster.
	VolumeSpecification interface{} `field:"required" json:"volumeSpecification" yaml:"volumeSpecification"`
	// Number of EBS volumes with a specific volume configuration that are associated with every instance in the instance group.
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}


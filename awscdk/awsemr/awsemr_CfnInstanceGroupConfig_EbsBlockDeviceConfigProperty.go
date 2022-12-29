package awsemr


// Configuration of requested EBS block device associated with the instance group with count of volumes that will be associated to every instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsBlockDeviceConfigProperty := &ebsBlockDeviceConfigProperty{
//   	volumeSpecification: &volumeSpecificationProperty{
//   		sizeInGb: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//
//   		// the properties below are optional
//   		iops: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	volumesPerInstance: jsii.Number(123),
//   }
//
type CfnInstanceGroupConfig_EbsBlockDeviceConfigProperty struct {
	// EBS volume specifications such as volume type, IOPS, and size (GiB) that will be requested for the EBS volume attached to an EC2 instance in the cluster.
	VolumeSpecification interface{} `field:"required" json:"volumeSpecification" yaml:"volumeSpecification"`
	// Number of EBS volumes with a specific volume configuration that will be associated with every instance in the instance group.
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}


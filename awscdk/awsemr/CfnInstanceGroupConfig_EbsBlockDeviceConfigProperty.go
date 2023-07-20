package awsemr


// Configuration of requested EBS block device associated with the instance group with count of volumes that are associated to every instance.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-ebsblockdeviceconfig.html
//
type CfnInstanceGroupConfig_EbsBlockDeviceConfigProperty struct {
	// EBS volume specifications such as volume type, IOPS, size (GiB) and throughput (MiB/s) that are requested for the EBS volume attached to an Amazon EC2 instance in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-ebsblockdeviceconfig.html#cfn-emr-instancegroupconfig-ebsblockdeviceconfig-volumespecification
	//
	VolumeSpecification interface{} `field:"required" json:"volumeSpecification" yaml:"volumeSpecification"`
	// Number of EBS volumes with a specific volume configuration that are associated with every instance in the instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancegroupconfig-ebsblockdeviceconfig.html#cfn-emr-instancegroupconfig-ebsblockdeviceconfig-volumesperinstance
	//
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}


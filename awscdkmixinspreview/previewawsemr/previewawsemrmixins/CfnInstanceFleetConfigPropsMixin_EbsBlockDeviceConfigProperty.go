package previewawsemrmixins


// `EbsBlockDeviceConfig` is a subproperty of the `EbsConfiguration` property type.
//
// `EbsBlockDeviceConfig` defines the number and type of EBS volumes to associate with all EC2 instances in an EMR cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ebsBlockDeviceConfigProperty := &EbsBlockDeviceConfigProperty{
//   	VolumeSpecification: &VolumeSpecificationProperty{
//   		Iops: jsii.Number(123),
//   		SizeInGb: jsii.Number(123),
//   		Throughput: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	VolumesPerInstance: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-ebsblockdeviceconfig.html
//
type CfnInstanceFleetConfigPropsMixin_EbsBlockDeviceConfigProperty struct {
	// EBS volume specifications such as volume type, IOPS, size (GiB) and throughput (MiB/s) that are requested for the EBS volume attached to an Amazon EC2 instance in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-ebsblockdeviceconfig.html#cfn-emr-instancefleetconfig-ebsblockdeviceconfig-volumespecification
	//
	VolumeSpecification interface{} `field:"optional" json:"volumeSpecification" yaml:"volumeSpecification"`
	// Number of EBS volumes with a specific volume configuration that are associated with every instance in the instance group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-ebsblockdeviceconfig.html#cfn-emr-instancefleetconfig-ebsblockdeviceconfig-volumesperinstance
	//
	VolumesPerInstance *float64 `field:"optional" json:"volumesPerInstance" yaml:"volumesPerInstance"`
}


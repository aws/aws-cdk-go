package awsemr


// `EbsConfiguration` is a subproperty of `InstanceFleetConfig` or `InstanceGroupConfig` .
//
// `EbsConfiguration` determines the EBS volumes to attach to EMR cluster instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsConfigurationProperty := &EbsConfigurationProperty{
//   	EbsBlockDeviceConfigs: []interface{}{
//   		&EbsBlockDeviceConfigProperty{
//   			VolumeSpecification: &VolumeSpecificationProperty{
//   				SizeInGb: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//
//   				// the properties below are optional
//   				Iops: jsii.Number(123),
//   				Throughput: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			VolumesPerInstance: jsii.Number(123),
//   		},
//   	},
//   	EbsOptimized: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ebsconfiguration.html
//
type CfnCluster_EbsConfigurationProperty struct {
	// An array of Amazon EBS volume specifications attached to a cluster instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ebsconfiguration.html#cfn-emr-cluster-ebsconfiguration-ebsblockdeviceconfigs
	//
	EbsBlockDeviceConfigs interface{} `field:"optional" json:"ebsBlockDeviceConfigs" yaml:"ebsBlockDeviceConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized.
	//
	// The default is false. You should explicitly set this value to true to enable the Amazon EBS-optimized setting for an EC2 instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-ebsconfiguration.html#cfn-emr-cluster-ebsconfiguration-ebsoptimized
	//
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
}


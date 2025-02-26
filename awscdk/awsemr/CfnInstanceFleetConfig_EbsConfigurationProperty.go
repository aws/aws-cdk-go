package awsemr


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-ebsconfiguration.html
//
type CfnInstanceFleetConfig_EbsConfigurationProperty struct {
	// An array of Amazon EBS volume specifications attached to a cluster instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-ebsconfiguration.html#cfn-emr-instancefleetconfig-ebsconfiguration-ebsblockdeviceconfigs
	//
	EbsBlockDeviceConfigs interface{} `field:"optional" json:"ebsBlockDeviceConfigs" yaml:"ebsBlockDeviceConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-ebsconfiguration.html#cfn-emr-instancefleetconfig-ebsconfiguration-ebsoptimized
	//
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
}


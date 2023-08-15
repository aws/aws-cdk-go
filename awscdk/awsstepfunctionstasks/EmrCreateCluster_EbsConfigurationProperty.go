package awsstepfunctionstasks


// The Amazon EBS configuration of a cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var size size
//
//   ebsConfigurationProperty := &EbsConfigurationProperty{
//   	EbsBlockDeviceConfigs: []ebsBlockDeviceConfigProperty{
//   		&ebsBlockDeviceConfigProperty{
//   			VolumeSpecification: &VolumeSpecificationProperty{
//   				VolumeSize: size,
//   				VolumeType: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.EbsBlockDeviceVolumeType_GP2,
//
//   				// the properties below are optional
//   				Iops: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			VolumesPerInstance: jsii.Number(123),
//   		},
//   	},
//   	EbsOptimized: jsii.Boolean(false),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_EbsConfiguration.html
//
type EmrCreateCluster_EbsConfigurationProperty struct {
	// An array of Amazon EBS volume specifications attached to a cluster instance.
	// Default: - None.
	//
	EbsBlockDeviceConfigs *[]*EmrCreateCluster_EbsBlockDeviceConfigProperty `field:"optional" json:"ebsBlockDeviceConfigs" yaml:"ebsBlockDeviceConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized.
	// Default: - EMR selected default.
	//
	EbsOptimized *bool `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
}


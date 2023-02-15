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
//   ebsConfigurationProperty := &ebsConfigurationProperty{
//   	ebsBlockDeviceConfigs: []ebsBlockDeviceConfigProperty{
//   		&ebsBlockDeviceConfigProperty{
//   			volumeSpecification: &volumeSpecificationProperty{
//   				volumeSize: size,
//   				volumeType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.ebsBlockDeviceVolumeType_GP2,
//
//   				// the properties below are optional
//   				iops: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			volumesPerInstance: jsii.Number(123),
//   		},
//   	},
//   	ebsOptimized: jsii.Boolean(false),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_EbsConfiguration.html
//
type EmrCreateCluster_EbsConfigurationProperty struct {
	// An array of Amazon EBS volume specifications attached to a cluster instance.
	EbsBlockDeviceConfigs *[]*EmrCreateCluster_EbsBlockDeviceConfigProperty `field:"optional" json:"ebsBlockDeviceConfigs" yaml:"ebsBlockDeviceConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized.
	EbsOptimized *bool `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
}


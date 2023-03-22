package awsemr


// The Amazon EBS configuration of a cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsConfigurationProperty := &ebsConfigurationProperty{
//   	ebsBlockDeviceConfigs: []interface{}{
//   		&ebsBlockDeviceConfigProperty{
//   			volumeSpecification: &volumeSpecificationProperty{
//   				sizeInGb: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
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
type CfnInstanceGroupConfig_EbsConfigurationProperty struct {
	// An array of Amazon EBS volume specifications attached to a cluster instance.
	EbsBlockDeviceConfigs interface{} `field:"optional" json:"ebsBlockDeviceConfigs" yaml:"ebsBlockDeviceConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
}


package awsstepfunctionstasks


// An instance type configuration for each instance type in an instance fleet, which determines the EC2 instances Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//   var size size
//
//   instanceTypeConfigProperty := &instanceTypeConfigProperty{
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	bidPrice: jsii.String("bidPrice"),
//   	bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   	configurations: []*configurationProperty{
//   		&configurationProperty{
//   			classification: jsii.String("classification"),
//   			configurations: []*configurationProperty{
//   				configurationProperty_,
//   			},
//   			properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	ebsConfiguration: &ebsConfigurationProperty{
//   		ebsBlockDeviceConfigs: []ebsBlockDeviceConfigProperty{
//   			&ebsBlockDeviceConfigProperty{
//   				volumeSpecification: &volumeSpecificationProperty{
//   					volumeSize: size,
//   					volumeType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.ebsBlockDeviceVolumeType_GP2,
//
//   					// the properties below are optional
//   					iops: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				volumesPerInstance: jsii.Number(123),
//   			},
//   		},
//   		ebsOptimized: jsii.Boolean(false),
//   	},
//   	weightedCapacity: jsii.Number(123),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceTypeConfig.html
//
type EmrCreateCluster_InstanceTypeConfigProperty struct {
	// An EC2 instance type.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The bid price for each EC2 Spot instance type as defined by InstanceType.
	//
	// Expressed in USD.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// The bid price, as a percentage of On-Demand price.
	BidPriceAsPercentageOfOnDemandPrice *float64 `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// A configuration classification that applies when provisioning cluster instances, which can include configurations for applications and software that run on the cluster.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `field:"optional" json:"configurations" yaml:"configurations"`
	// The configuration of Amazon Elastic Block Storage (EBS) attached to each instance as defined by InstanceType.
	EbsConfiguration *EmrCreateCluster_EbsConfigurationProperty `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// The number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in the InstanceFleetConfig.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


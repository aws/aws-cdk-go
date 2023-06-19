package awsstepfunctionstasks


// An instance type configuration for each instance type in an instance fleet, which determines the EC2 instances Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//   var size size
//
//   instanceTypeConfigProperty := &InstanceTypeConfigProperty{
//   	InstanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	BidPrice: jsii.String("bidPrice"),
//   	BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   	Configurations: []*configurationProperty{
//   		&configurationProperty{
//   			Classification: jsii.String("classification"),
//   			Configurations: []*configurationProperty{
//   				configurationProperty_,
//   			},
//   			Properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	EbsConfiguration: &EbsConfigurationProperty{
//   		EbsBlockDeviceConfigs: []ebsBlockDeviceConfigProperty{
//   			&ebsBlockDeviceConfigProperty{
//   				VolumeSpecification: &VolumeSpecificationProperty{
//   					VolumeSize: size,
//   					VolumeType: awscdk.Aws_stepfunctions_tasks.EmrCreateCluster.EbsBlockDeviceVolumeType_GP2,
//
//   					// the properties below are optional
//   					Iops: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				VolumesPerInstance: jsii.Number(123),
//   			},
//   		},
//   		EbsOptimized: jsii.Boolean(false),
//   	},
//   	WeightedCapacity: jsii.Number(123),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceTypeConfig.html
//
// Experimental.
type EmrCreateCluster_InstanceTypeConfigProperty struct {
	// An EC2 instance type.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The bid price for each EC2 Spot instance type as defined by InstanceType.
	//
	// Expressed in USD.
	// Experimental.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// The bid price, as a percentage of On-Demand price.
	// Experimental.
	BidPriceAsPercentageOfOnDemandPrice *float64 `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// A configuration classification that applies when provisioning cluster instances, which can include configurations for applications and software that run on the cluster.
	// Experimental.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `field:"optional" json:"configurations" yaml:"configurations"`
	// The configuration of Amazon Elastic Block Storage (EBS) attached to each instance as defined by InstanceType.
	// Experimental.
	EbsConfiguration *EmrCreateCluster_EbsConfigurationProperty `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// The number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in the InstanceFleetConfig.
	// Experimental.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


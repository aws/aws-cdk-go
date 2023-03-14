package awsemr


// `InstanceType` config is a subproperty of `InstanceFleetConfig` .
//
// An instance type configuration specifies each instance type in an instance fleet. The configuration determines the EC2 instances Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   instanceTypeConfigProperty := &instanceTypeConfigProperty{
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	bidPrice: jsii.String("bidPrice"),
//   	bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   	configurations: []interface{}{
//   		&configurationProperty{
//   			classification: jsii.String("classification"),
//   			configurationProperties: map[string]*string{
//   				"configurationPropertiesKey": jsii.String("configurationProperties"),
//   			},
//   			configurations: []interface{}{
//   				configurationProperty_,
//   			},
//   		},
//   	},
//   	customAmiId: jsii.String("customAmiId"),
//   	ebsConfiguration: &ebsConfigurationProperty{
//   		ebsBlockDeviceConfigs: []interface{}{
//   			&ebsBlockDeviceConfigProperty{
//   				volumeSpecification: &volumeSpecificationProperty{
//   					sizeInGb: jsii.Number(123),
//   					volumeType: jsii.String("volumeType"),
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
type CfnInstanceFleetConfig_InstanceTypeConfigProperty struct {
	// An EC2 instance type, such as `m3.xlarge` .
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The bid price for each EC2 Spot Instance type as defined by `InstanceType` .
	//
	// Expressed in USD. If neither `BidPrice` nor `BidPriceAsPercentageOfOnDemandPrice` is provided, `BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// The bid price, as a percentage of On-Demand price, for each EC2 Spot Instance as defined by `InstanceType` .
	//
	// Expressed as a number (for example, 20 specifies 20%). If neither `BidPrice` nor `BidPriceAsPercentageOfOnDemandPrice` is provided, `BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.
	BidPriceAsPercentageOfOnDemandPrice *float64 `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// > Amazon EMR releases 4.x or later.
	//
	// An optional configuration specification to be used when provisioning cluster instances, which can include configurations for applications and software bundled with Amazon EMR. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html) .
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// The custom AMI ID to use for the instance type.
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// The configuration of Amazon Elastic Block Store (Amazon EBS) attached to each instance as defined by `InstanceType` .
	EbsConfiguration interface{} `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// The number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in `InstanceFleetConfig` .
	//
	// This value is 1 for a master instance fleet, and must be 1 or greater for core and task instance fleets. Defaults to 1 if not specified.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


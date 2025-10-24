package awsemr


// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// `InstanceTypeConfig` is a sub-property of `InstanceFleetConfig` . `InstanceTypeConfig` determines the EC2 instances that Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ ConfigurationProperty
//
//   instanceTypeConfigProperty := &InstanceTypeConfigProperty{
//   	InstanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	BidPrice: jsii.String("bidPrice"),
//   	BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   	Configurations: []interface{}{
//   		&ConfigurationProperty{
//   			Classification: jsii.String("classification"),
//   			ConfigurationProperties: map[string]*string{
//   				"configurationPropertiesKey": jsii.String("configurationProperties"),
//   			},
//   			Configurations: []interface{}{
//   				configurationProperty_,
//   			},
//   		},
//   	},
//   	CustomAmiId: jsii.String("customAmiId"),
//   	EbsConfiguration: &EbsConfigurationProperty{
//   		EbsBlockDeviceConfigs: []interface{}{
//   			&EbsBlockDeviceConfigProperty{
//   				VolumeSpecification: &VolumeSpecificationProperty{
//   					SizeInGb: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//
//   					// the properties below are optional
//   					Iops: jsii.Number(123),
//   					Throughput: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				VolumesPerInstance: jsii.Number(123),
//   			},
//   		},
//   		EbsOptimized: jsii.Boolean(false),
//   	},
//   	Priority: jsii.Number(123),
//   	WeightedCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancetypeconfig.html
//
type CfnCluster_InstanceTypeConfigProperty struct {
	// An Amazon EC2 instance type, such as `m3.xlarge` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancetypeconfig.html#cfn-emr-cluster-instancetypeconfig-instancetype
	//
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The bid price for each Amazon EC2 Spot Instance type as defined by `InstanceType` .
	//
	// Expressed in USD. If neither `BidPrice` nor `BidPriceAsPercentageOfOnDemandPrice` is provided, `BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancetypeconfig.html#cfn-emr-cluster-instancetypeconfig-bidprice
	//
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// The bid price, as a percentage of On-Demand price, for each Amazon EC2 Spot Instance as defined by `InstanceType` .
	//
	// Expressed as a number (for example, 20 specifies 20%). If neither `BidPrice` nor `BidPriceAsPercentageOfOnDemandPrice` is provided, `BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancetypeconfig.html#cfn-emr-cluster-instancetypeconfig-bidpriceaspercentageofondemandprice
	//
	BidPriceAsPercentageOfOnDemandPrice *float64 `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// A configuration classification that applies when provisioning cluster instances, which can include configurations for applications and software that run on the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancetypeconfig.html#cfn-emr-cluster-instancetypeconfig-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// The custom AMI ID to use for the instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancetypeconfig.html#cfn-emr-cluster-instancetypeconfig-customamiid
	//
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// The configuration of Amazon Elastic Block Store (Amazon EBS) attached to each instance as defined by `InstanceType` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancetypeconfig.html#cfn-emr-cluster-instancetypeconfig-ebsconfiguration
	//
	EbsConfiguration interface{} `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// The priority at which Amazon EMR launches the Amazon EC2 instances with this instance type.
	//
	// Priority starts at 0, which is the highest priority. Amazon EMR considers the highest priority first.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancetypeconfig.html#cfn-emr-cluster-instancetypeconfig-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in `InstanceFleetConfig` .
	//
	// This value is 1 for a master instance fleet, and must be 1 or greater for core and task instance fleets. Defaults to 1 if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancetypeconfig.html#cfn-emr-cluster-instancetypeconfig-weightedcapacity
	//
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


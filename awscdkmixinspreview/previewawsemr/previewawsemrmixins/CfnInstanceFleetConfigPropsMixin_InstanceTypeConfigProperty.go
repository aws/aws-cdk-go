package previewawsemrmixins


// `InstanceType` config is a subproperty of `InstanceFleetConfig` .
//
// An instance type configuration specifies each instance type in an instance fleet. The configuration determines the EC2 instances Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var configurationProperty_ ConfigurationProperty
//
//   instanceTypeConfigProperty := &InstanceTypeConfigProperty{
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
//   					Iops: jsii.Number(123),
//   					SizeInGb: jsii.Number(123),
//   					Throughput: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//   				},
//   				VolumesPerInstance: jsii.Number(123),
//   			},
//   		},
//   		EbsOptimized: jsii.Boolean(false),
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	Priority: jsii.Number(123),
//   	WeightedCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancetypeconfig.html
//
type CfnInstanceFleetConfigPropsMixin_InstanceTypeConfigProperty struct {
	// The bid price for each Amazon EC2 Spot Instance type as defined by `InstanceType` .
	//
	// Expressed in USD. If neither `BidPrice` nor `BidPriceAsPercentageOfOnDemandPrice` is provided, `BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancetypeconfig.html#cfn-emr-instancefleetconfig-instancetypeconfig-bidprice
	//
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// The bid price, as a percentage of On-Demand price, for each Amazon EC2 Spot Instance as defined by `InstanceType` .
	//
	// Expressed as a number (for example, 20 specifies 20%). If neither `BidPrice` nor `BidPriceAsPercentageOfOnDemandPrice` is provided, `BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancetypeconfig.html#cfn-emr-instancefleetconfig-instancetypeconfig-bidpriceaspercentageofondemandprice
	//
	BidPriceAsPercentageOfOnDemandPrice *float64 `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// > Amazon EMR releases 4.x or later.
	//
	// An optional configuration specification to be used when provisioning cluster instances, which can include configurations for applications and software bundled with Amazon EMR. A configuration consists of a classification, properties, and optional nested configurations. A classification refers to an application-specific configuration file. Properties are the settings you want to change in that file. For more information, see [Configuring Applications](https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-configure-apps.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancetypeconfig.html#cfn-emr-instancefleetconfig-instancetypeconfig-configurations
	//
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// The custom AMI ID to use for the instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancetypeconfig.html#cfn-emr-instancefleetconfig-instancetypeconfig-customamiid
	//
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// The configuration of Amazon Elastic Block Store (Amazon EBS) attached to each instance as defined by `InstanceType` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancetypeconfig.html#cfn-emr-instancefleetconfig-instancetypeconfig-ebsconfiguration
	//
	EbsConfiguration interface{} `field:"optional" json:"ebsConfiguration" yaml:"ebsConfiguration"`
	// An Amazon EC2 instance type, such as `m3.xlarge` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancetypeconfig.html#cfn-emr-instancefleetconfig-instancetypeconfig-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The priority at which Amazon EMR launches the Amazon EC2 instances with this instance type.
	//
	// Priority starts at 0, which is the highest priority. Amazon EMR considers the highest priority first.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancetypeconfig.html#cfn-emr-instancefleetconfig-instancetypeconfig-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in `InstanceFleetConfig` .
	//
	// This value is 1 for a master instance fleet, and must be 1 or greater for core and task instance fleets. Defaults to 1 if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-instancefleetconfig-instancetypeconfig.html#cfn-emr-instancefleetconfig-instancetypeconfig-weightedcapacity
	//
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


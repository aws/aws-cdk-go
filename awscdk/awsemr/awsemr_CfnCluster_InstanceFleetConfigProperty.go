package awsemr


// Use `InstanceFleetConfig` to define instance fleets for an EMR cluster.
//
// A cluster can not use both instance fleets and instance groups. For more information, see [Configure Instance Fleets](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the *Amazon EMR Management Guide* .
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
//   instanceFleetConfigProperty := &instanceFleetConfigProperty{
//   	instanceTypeConfigs: []interface{}{
//   		&instanceTypeConfigProperty{
//   			instanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			bidPrice: jsii.String("bidPrice"),
//   			bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   			configurations: []interface{}{
//   				&configurationProperty{
//   					classification: jsii.String("classification"),
//   					configurationProperties: map[string]*string{
//   						"configurationPropertiesKey": jsii.String("configurationProperties"),
//   					},
//   					configurations: []interface{}{
//   						configurationProperty_,
//   					},
//   				},
//   			},
//   			customAmiId: jsii.String("customAmiId"),
//   			ebsConfiguration: &ebsConfigurationProperty{
//   				ebsBlockDeviceConfigs: []interface{}{
//   					&ebsBlockDeviceConfigProperty{
//   						volumeSpecification: &volumeSpecificationProperty{
//   							sizeInGb: jsii.Number(123),
//   							volumeType: jsii.String("volumeType"),
//
//   							// the properties below are optional
//   							iops: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						volumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				ebsOptimized: jsii.Boolean(false),
//   			},
//   			weightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	launchSpecifications: &instanceFleetProvisioningSpecificationsProperty{
//   		onDemandSpecification: &onDemandProvisioningSpecificationProperty{
//   			allocationStrategy: jsii.String("allocationStrategy"),
//   		},
//   		spotSpecification: &spotProvisioningSpecificationProperty{
//   			timeoutAction: jsii.String("timeoutAction"),
//   			timeoutDurationMinutes: jsii.Number(123),
//
//   			// the properties below are optional
//   			allocationStrategy: jsii.String("allocationStrategy"),
//   			blockDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	targetOnDemandCapacity: jsii.Number(123),
//   	targetSpotCapacity: jsii.Number(123),
//   }
//
type CfnCluster_InstanceFleetConfigProperty struct {
	// The instance type configurations that define the EC2 instances in the instance fleet.
	InstanceTypeConfigs interface{} `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// The launch specification for the instance fleet.
	LaunchSpecifications interface{} `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// The friendly name of the instance fleet.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision On-Demand instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When an On-Demand instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only Spot instances are provisioned for the instance fleet using `TargetSpotCapacity` . At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision Spot instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When a Spot instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only On-Demand instances are provisioned for the instance fleet. At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}


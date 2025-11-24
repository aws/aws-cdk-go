package mixinsawsemr


// Use `InstanceFleetConfig` to define instance fleets for an EMR cluster.
//
// A cluster can not use both instance fleets and instance groups. For more information, see [Configure Instance Fleets](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the *Amazon EMR Management Guide* .
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
//   instanceFleetConfigProperty := &InstanceFleetConfigProperty{
//   	InstanceTypeConfigs: []interface{}{
//   		&InstanceTypeConfigProperty{
//   			BidPrice: jsii.String("bidPrice"),
//   			BidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   			Configurations: []interface{}{
//   				&ConfigurationProperty{
//   					Classification: jsii.String("classification"),
//   					ConfigurationProperties: map[string]*string{
//   						"configurationPropertiesKey": jsii.String("configurationProperties"),
//   					},
//   					Configurations: []interface{}{
//   						configurationProperty_,
//   					},
//   				},
//   			},
//   			CustomAmiId: jsii.String("customAmiId"),
//   			EbsConfiguration: &EbsConfigurationProperty{
//   				EbsBlockDeviceConfigs: []interface{}{
//   					&EbsBlockDeviceConfigProperty{
//   						VolumeSpecification: &VolumeSpecificationProperty{
//   							Iops: jsii.Number(123),
//   							SizeInGb: jsii.Number(123),
//   							Throughput: jsii.Number(123),
//   							VolumeType: jsii.String("volumeType"),
//   						},
//   						VolumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				EbsOptimized: jsii.Boolean(false),
//   			},
//   			InstanceType: jsii.String("instanceType"),
//   			Priority: jsii.Number(123),
//   			WeightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	LaunchSpecifications: &InstanceFleetProvisioningSpecificationsProperty{
//   		OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   			AllocationStrategy: jsii.String("allocationStrategy"),
//   			CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   				CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   				CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   				UsageStrategy: jsii.String("usageStrategy"),
//   			},
//   		},
//   		SpotSpecification: &SpotProvisioningSpecificationProperty{
//   			AllocationStrategy: jsii.String("allocationStrategy"),
//   			BlockDurationMinutes: jsii.Number(123),
//   			TimeoutAction: jsii.String("timeoutAction"),
//   			TimeoutDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	ResizeSpecifications: &InstanceFleetResizingSpecificationsProperty{
//   		OnDemandResizeSpecification: &OnDemandResizingSpecificationProperty{
//   			AllocationStrategy: jsii.String("allocationStrategy"),
//   			CapacityReservationOptions: &OnDemandCapacityReservationOptionsProperty{
//   				CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   				CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   				UsageStrategy: jsii.String("usageStrategy"),
//   			},
//   			TimeoutDurationMinutes: jsii.Number(123),
//   		},
//   		SpotResizeSpecification: &SpotResizingSpecificationProperty{
//   			AllocationStrategy: jsii.String("allocationStrategy"),
//   			TimeoutDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	TargetOnDemandCapacity: jsii.Number(123),
//   	TargetSpotCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancefleetconfig.html
//
type CfnClusterPropsMixin_InstanceFleetConfigProperty struct {
	// The instance type configurations that define the Amazon EC2 instances in the instance fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancefleetconfig.html#cfn-emr-cluster-instancefleetconfig-instancetypeconfigs
	//
	InstanceTypeConfigs interface{} `field:"optional" json:"instanceTypeConfigs" yaml:"instanceTypeConfigs"`
	// The launch specification for the instance fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancefleetconfig.html#cfn-emr-cluster-instancefleetconfig-launchspecifications
	//
	LaunchSpecifications interface{} `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// The friendly name of the instance fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancefleetconfig.html#cfn-emr-cluster-instancefleetconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The resize specification for the instance fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancefleetconfig.html#cfn-emr-cluster-instancefleetconfig-resizespecifications
	//
	ResizeSpecifications interface{} `field:"optional" json:"resizeSpecifications" yaml:"resizeSpecifications"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision On-Demand instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When an On-Demand instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only Spot instances are provisioned for the instance fleet using `TargetSpotCapacity` . At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancefleetconfig.html#cfn-emr-cluster-instancefleetconfig-targetondemandcapacity
	//
	TargetOnDemandCapacity *float64 `field:"optional" json:"targetOnDemandCapacity" yaml:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	//
	// When the instance fleet launches, Amazon EMR tries to provision Spot instances as specified by `InstanceTypeConfig` . Each instance configuration has a specified `WeightedCapacity` . When a Spot instance is provisioned, the `WeightedCapacity` units count toward the target capacity. Amazon EMR provisions instances until the target capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EMR can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the target capacity is exceeded by 3 units.
	//
	// > If not specified or set to 0, only On-Demand instances are provisioned for the instance fleet. At least one of `TargetSpotCapacity` and `TargetOnDemandCapacity` should be greater than 0. For a master instance fleet, only one of `TargetSpotCapacity` and `TargetOnDemandCapacity` can be specified, and its value must be 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-instancefleetconfig.html#cfn-emr-cluster-instancefleetconfig-targetspotcapacity
	//
	TargetSpotCapacity *float64 `field:"optional" json:"targetSpotCapacity" yaml:"targetSpotCapacity"`
}


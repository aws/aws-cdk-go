package awsec2


// Properties for defining a `CfnEC2Fleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEC2FleetProps := &CfnEC2FleetProps{
//   	LaunchTemplateConfigs: []interface{}{
//   		&FleetLaunchTemplateConfigRequestProperty{
//   			LaunchTemplateSpecification: &FleetLaunchTemplateSpecificationRequestProperty{
//   				Version: jsii.String("version"),
//
//   				// the properties below are optional
//   				LaunchTemplateId: jsii.String("launchTemplateId"),
//   				LaunchTemplateName: jsii.String("launchTemplateName"),
//   			},
//   			Overrides: []interface{}{
//   				&FleetLaunchTemplateOverridesRequestProperty{
//   					AvailabilityZone: jsii.String("availabilityZone"),
//   					InstanceRequirements: &InstanceRequirementsRequestProperty{
//   						AcceleratorCount: &AcceleratorCountRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						AcceleratorManufacturers: []*string{
//   							jsii.String("acceleratorManufacturers"),
//   						},
//   						AcceleratorNames: []*string{
//   							jsii.String("acceleratorNames"),
//   						},
//   						AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						AcceleratorTypes: []*string{
//   							jsii.String("acceleratorTypes"),
//   						},
//   						AllowedInstanceTypes: []*string{
//   							jsii.String("allowedInstanceTypes"),
//   						},
//   						BareMetal: jsii.String("bareMetal"),
//   						BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						BurstablePerformance: jsii.String("burstablePerformance"),
//   						CpuManufacturers: []*string{
//   							jsii.String("cpuManufacturers"),
//   						},
//   						ExcludedInstanceTypes: []*string{
//   							jsii.String("excludedInstanceTypes"),
//   						},
//   						InstanceGenerations: []*string{
//   							jsii.String("instanceGenerations"),
//   						},
//   						LocalStorage: jsii.String("localStorage"),
//   						LocalStorageTypes: []*string{
//   							jsii.String("localStorageTypes"),
//   						},
//   						MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   						MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						MemoryMiB: &MemoryMiBRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						RequireHibernateSupport: jsii.Boolean(false),
//   						SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						VCpuCount: &VCpuCountRangeRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   					},
//   					InstanceType: jsii.String("instanceType"),
//   					MaxPrice: jsii.String("maxPrice"),
//   					Placement: &PlacementProperty{
//   						Affinity: jsii.String("affinity"),
//   						AvailabilityZone: jsii.String("availabilityZone"),
//   						GroupName: jsii.String("groupName"),
//   						HostId: jsii.String("hostId"),
//   						HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   						PartitionNumber: jsii.Number(123),
//   						SpreadDomain: jsii.String("spreadDomain"),
//   						Tenancy: jsii.String("tenancy"),
//   					},
//   					Priority: jsii.Number(123),
//   					SubnetId: jsii.String("subnetId"),
//   					WeightedCapacity: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	TargetCapacitySpecification: &TargetCapacitySpecificationRequestProperty{
//   		TotalTargetCapacity: jsii.Number(123),
//
//   		// the properties below are optional
//   		DefaultTargetCapacityType: jsii.String("defaultTargetCapacityType"),
//   		OnDemandTargetCapacity: jsii.Number(123),
//   		SpotTargetCapacity: jsii.Number(123),
//   		TargetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   	},
//
//   	// the properties below are optional
//   	Context: jsii.String("context"),
//   	ExcessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   	OnDemandOptions: &OnDemandOptionsRequestProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		CapacityReservationOptions: &CapacityReservationOptionsRequestProperty{
//   			UsageStrategy: jsii.String("usageStrategy"),
//   		},
//   		MaxTotalPrice: jsii.String("maxTotalPrice"),
//   		MinTargetCapacity: jsii.Number(123),
//   		SingleAvailabilityZone: jsii.Boolean(false),
//   		SingleInstanceType: jsii.Boolean(false),
//   	},
//   	ReplaceUnhealthyInstances: jsii.Boolean(false),
//   	SpotOptions: &SpotOptionsRequestProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		InstancePoolsToUseCount: jsii.Number(123),
//   		MaintenanceStrategies: &MaintenanceStrategiesProperty{
//   			CapacityRebalance: &CapacityRebalanceProperty{
//   				ReplacementStrategy: jsii.String("replacementStrategy"),
//   				TerminationDelay: jsii.Number(123),
//   			},
//   		},
//   		MaxTotalPrice: jsii.String("maxTotalPrice"),
//   		MinTargetCapacity: jsii.Number(123),
//   		SingleAvailabilityZone: jsii.Boolean(false),
//   		SingleInstanceType: jsii.Boolean(false),
//   	},
//   	TagSpecifications: []interface{}{
//   		&TagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []cfnTag{
//   				&cfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	TerminateInstancesWithExpiration: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   	ValidFrom: jsii.String("validFrom"),
//   	ValidUntil: jsii.String("validUntil"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html
//
type CfnEC2FleetProps struct {
	// The configuration for the EC2 Fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-launchtemplateconfigs
	//
	LaunchTemplateConfigs interface{} `field:"required" json:"launchTemplateConfigs" yaml:"launchTemplateConfigs"`
	// The number of units to request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-targetcapacityspecification
	//
	TargetCapacitySpecification interface{} `field:"required" json:"targetCapacitySpecification" yaml:"targetCapacitySpecification"`
	// Reserved.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-context
	//
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Indicates whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2 Fleet.
	//
	// Supported only for fleets of type `maintain` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-excesscapacityterminationpolicy
	//
	ExcessCapacityTerminationPolicy *string `field:"optional" json:"excessCapacityTerminationPolicy" yaml:"excessCapacityTerminationPolicy"`
	// Describes the configuration of On-Demand Instances in an EC2 Fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-ondemandoptions
	//
	OnDemandOptions interface{} `field:"optional" json:"onDemandOptions" yaml:"onDemandOptions"`
	// Indicates whether EC2 Fleet should replace unhealthy Spot Instances.
	//
	// Supported only for fleets of type `maintain` . For more information, see [EC2 Fleet health checks](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#ec2-fleet-health-checks) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-replaceunhealthyinstances
	//
	ReplaceUnhealthyInstances interface{} `field:"optional" json:"replaceUnhealthyInstances" yaml:"replaceUnhealthyInstances"`
	// Describes the configuration of Spot Instances in an EC2 Fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-spotoptions
	//
	SpotOptions interface{} `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// The key-value pair for tagging the EC2 Fleet request on creation. For more information, see [Tag your resources](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#tag-resources) .
	//
	// If the fleet type is `instant` , specify a resource type of `fleet` to tag the fleet or `instance` to tag the instances at launch.
	//
	// If the fleet type is `maintain` or `request` , specify a resource type of `fleet` to tag the fleet. You cannot specify a resource type of `instance` . To tag instances at launch, specify the tags in a [launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#create-launch-template) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-tagspecifications
	//
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Indicates whether running instances should be terminated when the EC2 Fleet expires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-terminateinstanceswithexpiration
	//
	TerminateInstancesWithExpiration interface{} `field:"optional" json:"terminateInstancesWithExpiration" yaml:"terminateInstancesWithExpiration"`
	// The fleet type. The default value is `maintain` .
	//
	// - `maintain` - The EC2 Fleet places an asynchronous request for your desired capacity, and continues to maintain your desired Spot capacity by replenishing interrupted Spot Instances.
	// - `request` - The EC2 Fleet places an asynchronous one-time request for your desired capacity, but does submit Spot requests in alternative capacity pools if Spot capacity is unavailable, and does not maintain Spot capacity if Spot Instances are interrupted.
	// - `instant` - The EC2 Fleet places a synchronous one-time request for your desired capacity, and returns errors for any instances that could not be launched.
	//
	// For more information, see [EC2 Fleet request types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-request-type.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The start date and time of the request, in UTC format (for example, *YYYY* - *MM* - *DD* T *HH* : *MM* : *SS* Z).
	//
	// The default is to start fulfilling the request immediately.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validfrom
	//
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// The end date and time of the request, in UTC format (for example, *YYYY* - *MM* - *DD* T *HH* : *MM* : *SS* Z).
	//
	// At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html#cfn-ec2-ec2fleet-validuntil
	//
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}


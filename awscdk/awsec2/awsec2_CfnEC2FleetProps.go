package awsec2


// Properties for defining a `CfnEC2Fleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEC2FleetProps := &cfnEC2FleetProps{
//   	launchTemplateConfigs: []interface{}{
//   		&fleetLaunchTemplateConfigRequestProperty{
//   			launchTemplateSpecification: &fleetLaunchTemplateSpecificationRequestProperty{
//   				version: jsii.String("version"),
//
//   				// the properties below are optional
//   				launchTemplateId: jsii.String("launchTemplateId"),
//   				launchTemplateName: jsii.String("launchTemplateName"),
//   			},
//   			overrides: []interface{}{
//   				&fleetLaunchTemplateOverridesRequestProperty{
//   					availabilityZone: jsii.String("availabilityZone"),
//   					instanceRequirements: &instanceRequirementsRequestProperty{
//   						acceleratorCount: &acceleratorCountRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						acceleratorManufacturers: []*string{
//   							jsii.String("acceleratorManufacturers"),
//   						},
//   						acceleratorNames: []*string{
//   							jsii.String("acceleratorNames"),
//   						},
//   						acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						acceleratorTypes: []*string{
//   							jsii.String("acceleratorTypes"),
//   						},
//   						allowedInstanceTypes: []*string{
//   							jsii.String("allowedInstanceTypes"),
//   						},
//   						bareMetal: jsii.String("bareMetal"),
//   						baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						burstablePerformance: jsii.String("burstablePerformance"),
//   						cpuManufacturers: []*string{
//   							jsii.String("cpuManufacturers"),
//   						},
//   						excludedInstanceTypes: []*string{
//   							jsii.String("excludedInstanceTypes"),
//   						},
//   						instanceGenerations: []*string{
//   							jsii.String("instanceGenerations"),
//   						},
//   						localStorage: jsii.String("localStorage"),
//   						localStorageTypes: []*string{
//   							jsii.String("localStorageTypes"),
//   						},
//   						memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						memoryMiB: &memoryMiBRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						networkBandwidthGbps: &networkBandwidthGbpsRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						requireHibernateSupport: jsii.Boolean(false),
//   						spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   						vCpuCount: &vCpuCountRangeRequestProperty{
//   							max: jsii.Number(123),
//   							min: jsii.Number(123),
//   						},
//   					},
//   					instanceType: jsii.String("instanceType"),
//   					maxPrice: jsii.String("maxPrice"),
//   					placement: &placementProperty{
//   						affinity: jsii.String("affinity"),
//   						availabilityZone: jsii.String("availabilityZone"),
//   						groupName: jsii.String("groupName"),
//   						hostId: jsii.String("hostId"),
//   						hostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   						partitionNumber: jsii.Number(123),
//   						spreadDomain: jsii.String("spreadDomain"),
//   						tenancy: jsii.String("tenancy"),
//   					},
//   					priority: jsii.Number(123),
//   					subnetId: jsii.String("subnetId"),
//   					weightedCapacity: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	targetCapacitySpecification: &targetCapacitySpecificationRequestProperty{
//   		totalTargetCapacity: jsii.Number(123),
//
//   		// the properties below are optional
//   		defaultTargetCapacityType: jsii.String("defaultTargetCapacityType"),
//   		onDemandTargetCapacity: jsii.Number(123),
//   		spotTargetCapacity: jsii.Number(123),
//   		targetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   	},
//
//   	// the properties below are optional
//   	context: jsii.String("context"),
//   	excessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   	onDemandOptions: &onDemandOptionsRequestProperty{
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		capacityReservationOptions: &capacityReservationOptionsRequestProperty{
//   			usageStrategy: jsii.String("usageStrategy"),
//   		},
//   		maxTotalPrice: jsii.String("maxTotalPrice"),
//   		minTargetCapacity: jsii.Number(123),
//   		singleAvailabilityZone: jsii.Boolean(false),
//   		singleInstanceType: jsii.Boolean(false),
//   	},
//   	replaceUnhealthyInstances: jsii.Boolean(false),
//   	spotOptions: &spotOptionsRequestProperty{
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		instanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		instancePoolsToUseCount: jsii.Number(123),
//   		maintenanceStrategies: &maintenanceStrategiesProperty{
//   			capacityRebalance: &capacityRebalanceProperty{
//   				replacementStrategy: jsii.String("replacementStrategy"),
//   				terminationDelay: jsii.Number(123),
//   			},
//   		},
//   		maxTotalPrice: jsii.String("maxTotalPrice"),
//   		minTargetCapacity: jsii.Number(123),
//   		singleAvailabilityZone: jsii.Boolean(false),
//   		singleInstanceType: jsii.Boolean(false),
//   	},
//   	tagSpecifications: []interface{}{
//   		&tagSpecificationProperty{
//   			resourceType: jsii.String("resourceType"),
//   			tags: []cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	terminateInstancesWithExpiration: jsii.Boolean(false),
//   	type: jsii.String("type"),
//   	validFrom: jsii.String("validFrom"),
//   	validUntil: jsii.String("validUntil"),
//   }
//
type CfnEC2FleetProps struct {
	// The configuration for the EC2 Fleet.
	LaunchTemplateConfigs interface{} `field:"required" json:"launchTemplateConfigs" yaml:"launchTemplateConfigs"`
	// The number of units to request.
	TargetCapacitySpecification interface{} `field:"required" json:"targetCapacitySpecification" yaml:"targetCapacitySpecification"`
	// Reserved.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Indicates whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2 Fleet.
	ExcessCapacityTerminationPolicy *string `field:"optional" json:"excessCapacityTerminationPolicy" yaml:"excessCapacityTerminationPolicy"`
	// Describes the configuration of On-Demand Instances in an EC2 Fleet.
	OnDemandOptions interface{} `field:"optional" json:"onDemandOptions" yaml:"onDemandOptions"`
	// Indicates whether EC2 Fleet should replace unhealthy Spot Instances.
	//
	// Supported only for fleets of type `maintain` . For more information, see [EC2 Fleet health checks](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#ec2-fleet-health-checks) in the *Amazon EC2 User Guide* .
	ReplaceUnhealthyInstances interface{} `field:"optional" json:"replaceUnhealthyInstances" yaml:"replaceUnhealthyInstances"`
	// Describes the configuration of Spot Instances in an EC2 Fleet.
	SpotOptions interface{} `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// The key-value pair for tagging the EC2 Fleet request on creation. For more information, see [Tagging your resources](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#tag-resources) .
	//
	// If the fleet type is `instant` , specify a resource type of `fleet` to tag the fleet or `instance` to tag the instances at launch.
	//
	// If the fleet type is `maintain` or `request` , specify a resource type of `fleet` to tag the fleet. You cannot specify a resource type of `instance` . To tag instances at launch, specify the tags in a [launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#create-launch-template) .
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Indicates whether running instances should be terminated when the EC2 Fleet expires.
	TerminateInstancesWithExpiration interface{} `field:"optional" json:"terminateInstancesWithExpiration" yaml:"terminateInstancesWithExpiration"`
	// The fleet type. The default value is `maintain` .
	//
	// - `maintain` - The EC2 Fleet places an asynchronous request for your desired capacity, and continues to maintain your desired Spot capacity by replenishing interrupted Spot Instances.
	// - `request` - The EC2 Fleet places an asynchronous one-time request for your desired capacity, but does submit Spot requests in alternative capacity pools if Spot capacity is unavailable, and does not maintain Spot capacity if Spot Instances are interrupted.
	// - `instant` - The EC2 Fleet places a synchronous one-time request for your desired capacity, and returns errors for any instances that could not be launched.
	//
	// For more information, see [EC2 Fleet request types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-request-type.html) in the *Amazon EC2 User Guide* .
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The start date and time of the request, in UTC format (for example, *YYYY* - *MM* - *DD* T *HH* : *MM* : *SS* Z).
	//
	// The default is to start fulfilling the request immediately.
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// The end date and time of the request, in UTC format (for example, *YYYY* - *MM* - *DD* T *HH* : *MM* : *SS* Z).
	//
	// At this point, no new EC2 Fleet requests are placed or able to fulfill the request. If no value is specified, the request remains until you cancel it.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}


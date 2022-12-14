package awsec2


// Specifies the configuration of a Spot Fleet request.
//
// For more information, see [How Spot Fleet Works](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html) in the *Amazon EC2 User Guide for Linux Instances* .
//
// You must specify either `LaunchSpecifications` or `LaunchTemplateConfigs` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotFleetRequestConfigDataProperty := &spotFleetRequestConfigDataProperty{
//   	iamFleetRole: jsii.String("iamFleetRole"),
//   	targetCapacity: jsii.Number(123),
//
//   	// the properties below are optional
//   	allocationStrategy: jsii.String("allocationStrategy"),
//   	context: jsii.String("context"),
//   	excessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   	instanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   	instancePoolsToUseCount: jsii.Number(123),
//   	launchSpecifications: []interface{}{
//   		&spotFleetLaunchSpecificationProperty{
//   			imageId: jsii.String("imageId"),
//
//   			// the properties below are optional
//   			blockDeviceMappings: []interface{}{
//   				&blockDeviceMappingProperty{
//   					deviceName: jsii.String("deviceName"),
//
//   					// the properties below are optional
//   					ebs: &ebsBlockDeviceProperty{
//   						deleteOnTermination: jsii.Boolean(false),
//   						encrypted: jsii.Boolean(false),
//   						iops: jsii.Number(123),
//   						snapshotId: jsii.String("snapshotId"),
//   						volumeSize: jsii.Number(123),
//   						volumeType: jsii.String("volumeType"),
//   					},
//   					noDevice: jsii.String("noDevice"),
//   					virtualName: jsii.String("virtualName"),
//   				},
//   			},
//   			ebsOptimized: jsii.Boolean(false),
//   			iamInstanceProfile: &iamInstanceProfileSpecificationProperty{
//   				arn: jsii.String("arn"),
//   			},
//   			instanceRequirements: &instanceRequirementsRequestProperty{
//   				acceleratorCount: &acceleratorCountRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				acceleratorManufacturers: []*string{
//   					jsii.String("acceleratorManufacturers"),
//   				},
//   				acceleratorNames: []*string{
//   					jsii.String("acceleratorNames"),
//   				},
//   				acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				acceleratorTypes: []*string{
//   					jsii.String("acceleratorTypes"),
//   				},
//   				allowedInstanceTypes: []*string{
//   					jsii.String("allowedInstanceTypes"),
//   				},
//   				bareMetal: jsii.String("bareMetal"),
//   				baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				burstablePerformance: jsii.String("burstablePerformance"),
//   				cpuManufacturers: []*string{
//   					jsii.String("cpuManufacturers"),
//   				},
//   				excludedInstanceTypes: []*string{
//   					jsii.String("excludedInstanceTypes"),
//   				},
//   				instanceGenerations: []*string{
//   					jsii.String("instanceGenerations"),
//   				},
//   				localStorage: jsii.String("localStorage"),
//   				localStorageTypes: []*string{
//   					jsii.String("localStorageTypes"),
//   				},
//   				memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				memoryMiB: &memoryMiBRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				networkBandwidthGbps: &networkBandwidthGbpsRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				requireHibernateSupport: jsii.Boolean(false),
//   				spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				vCpuCount: &vCpuCountRangeRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   			},
//   			instanceType: jsii.String("instanceType"),
//   			kernelId: jsii.String("kernelId"),
//   			keyName: jsii.String("keyName"),
//   			monitoring: &spotFleetMonitoringProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   			networkInterfaces: []interface{}{
//   				&instanceNetworkInterfaceSpecificationProperty{
//   					associatePublicIpAddress: jsii.Boolean(false),
//   					deleteOnTermination: jsii.Boolean(false),
//   					description: jsii.String("description"),
//   					deviceIndex: jsii.Number(123),
//   					groups: []*string{
//   						jsii.String("groups"),
//   					},
//   					ipv6AddressCount: jsii.Number(123),
//   					ipv6Addresses: []interface{}{
//   						&instanceIpv6AddressProperty{
//   							ipv6Address: jsii.String("ipv6Address"),
//   						},
//   					},
//   					networkInterfaceId: jsii.String("networkInterfaceId"),
//   					privateIpAddresses: []interface{}{
//   						&privateIpAddressSpecificationProperty{
//   							privateIpAddress: jsii.String("privateIpAddress"),
//
//   							// the properties below are optional
//   							primary: jsii.Boolean(false),
//   						},
//   					},
//   					secondaryPrivateIpAddressCount: jsii.Number(123),
//   					subnetId: jsii.String("subnetId"),
//   				},
//   			},
//   			placement: &spotPlacementProperty{
//   				availabilityZone: jsii.String("availabilityZone"),
//   				groupName: jsii.String("groupName"),
//   				tenancy: jsii.String("tenancy"),
//   			},
//   			ramdiskId: jsii.String("ramdiskId"),
//   			securityGroups: []interface{}{
//   				&groupIdentifierProperty{
//   					groupId: jsii.String("groupId"),
//   				},
//   			},
//   			spotPrice: jsii.String("spotPrice"),
//   			subnetId: jsii.String("subnetId"),
//   			tagSpecifications: []interface{}{
//   				&spotFleetTagSpecificationProperty{
//   					resourceType: jsii.String("resourceType"),
//   					tags: []cfnTag{
//   						&cfnTag{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			userData: jsii.String("userData"),
//   			weightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	launchTemplateConfigs: []interface{}{
//   		&launchTemplateConfigProperty{
//   			launchTemplateSpecification: &fleetLaunchTemplateSpecificationProperty{
//   				version: jsii.String("version"),
//
//   				// the properties below are optional
//   				launchTemplateId: jsii.String("launchTemplateId"),
//   				launchTemplateName: jsii.String("launchTemplateName"),
//   			},
//   			overrides: []interface{}{
//   				&launchTemplateOverridesProperty{
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
//   					priority: jsii.Number(123),
//   					spotPrice: jsii.String("spotPrice"),
//   					subnetId: jsii.String("subnetId"),
//   					weightedCapacity: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	loadBalancersConfig: &loadBalancersConfigProperty{
//   		classicLoadBalancersConfig: &classicLoadBalancersConfigProperty{
//   			classicLoadBalancers: []interface{}{
//   				&classicLoadBalancerProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		targetGroupsConfig: &targetGroupsConfigProperty{
//   			targetGroups: []interface{}{
//   				&targetGroupProperty{
//   					arn: jsii.String("arn"),
//   				},
//   			},
//   		},
//   	},
//   	onDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   	onDemandMaxTotalPrice: jsii.String("onDemandMaxTotalPrice"),
//   	onDemandTargetCapacity: jsii.Number(123),
//   	replaceUnhealthyInstances: jsii.Boolean(false),
//   	spotMaintenanceStrategies: &spotMaintenanceStrategiesProperty{
//   		capacityRebalance: &spotCapacityRebalanceProperty{
//   			replacementStrategy: jsii.String("replacementStrategy"),
//   			terminationDelay: jsii.Number(123),
//   		},
//   	},
//   	spotMaxTotalPrice: jsii.String("spotMaxTotalPrice"),
//   	spotPrice: jsii.String("spotPrice"),
//   	tagSpecifications: []interface{}{
//   		&spotFleetTagSpecificationProperty{
//   			resourceType: jsii.String("resourceType"),
//   			tags: []*cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	targetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   	terminateInstancesWithExpiration: jsii.Boolean(false),
//   	type: jsii.String("type"),
//   	validFrom: jsii.String("validFrom"),
//   	validUntil: jsii.String("validUntil"),
//   }
//
type CfnSpotFleet_SpotFleetRequestConfigDataProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM) role that grants the Spot Fleet the permission to request, launch, terminate, and tag instances on your behalf.
	//
	// For more information, see [Spot Fleet Prerequisites](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-requests.html#spot-fleet-prerequisites) in the *Amazon EC2 User Guide for Linux Instances* . Spot Fleet can terminate Spot Instances on your behalf when you cancel its Spot Fleet request or when the Spot Fleet request expires, if you set `TerminateInstancesWithExpiration` .
	IamFleetRole *string `field:"required" json:"iamFleetRole" yaml:"iamFleetRole"`
	// The number of units to request for the Spot Fleet.
	//
	// You can choose to set the target capacity in terms of instances or a performance characteristic that is important to your application workload, such as vCPUs, memory, or I/O. If the request type is `maintain` , you can specify a target capacity of 0 and add capacity later.
	TargetCapacity *float64 `field:"required" json:"targetCapacity" yaml:"targetCapacity"`
	// Indicates how to allocate the target Spot Instance capacity across the Spot Instance pools specified by the Spot Fleet request.
	//
	// If the allocation strategy is `lowestPrice` , Spot Fleet launches instances from the Spot Instance pools with the lowest price. This is the default allocation strategy.
	//
	// If the allocation strategy is `diversified` , Spot Fleet launches instances from all the Spot Instance pools that you specify.
	//
	// If the allocation strategy is `capacityOptimized` (recommended), Spot Fleet launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching. To give certain instance types a higher chance of launching first, use `capacityOptimizedPrioritized` . Set a priority for each instance type by using the `Priority` parameter for `LaunchTemplateOverrides` . You can assign the same priority to different `LaunchTemplateOverrides` . EC2 implements the priorities on a best-effort basis, but optimizes for capacity first. `capacityOptimizedPrioritized` is supported only if your Spot Fleet uses a launch template. Note that if the `OnDemandAllocationStrategy` is set to `prioritized` , the same priority is applied when fulfilling On-Demand capacity.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Reserved.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Indicates whether running Spot Instances should be terminated if you decrease the target capacity of the Spot Fleet request below the current size of the Spot Fleet.
	ExcessCapacityTerminationPolicy *string `field:"optional" json:"excessCapacityTerminationPolicy" yaml:"excessCapacityTerminationPolicy"`
	// The behavior when a Spot Instance is interrupted.
	//
	// The default is `terminate` .
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// The number of Spot pools across which to allocate your target Spot capacity.
	//
	// Valid only when Spot *AllocationStrategy* is set to `lowest-price` . Spot Fleet selects the cheapest Spot pools and evenly allocates your target Spot capacity across the number of Spot pools that you specify.
	//
	// Note that Spot Fleet attempts to draw Spot Instances from the number of pools that you specify on a best effort basis. If a pool runs out of Spot capacity before fulfilling your target capacity, Spot Fleet will continue to fulfill your request by drawing from the next cheapest pool. To ensure that your target capacity is met, you might receive Spot Instances from more than the number of pools that you specified. Similarly, if most of the pools have no Spot capacity, you might receive your full target capacity from fewer than the number of pools that you specified.
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// The launch specifications for the Spot Fleet request.
	//
	// If you specify `LaunchSpecifications` , you can't specify `LaunchTemplateConfigs` .
	LaunchSpecifications interface{} `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// The launch template and overrides.
	//
	// If you specify `LaunchTemplateConfigs` , you can't specify `LaunchSpecifications` .
	LaunchTemplateConfigs interface{} `field:"optional" json:"launchTemplateConfigs" yaml:"launchTemplateConfigs"`
	// One or more Classic Load Balancers and target groups to attach to the Spot Fleet request.
	//
	// Spot Fleet registers the running Spot Instances with the specified Classic Load Balancers and target groups.
	//
	// With Network Load Balancers, Spot Fleet cannot register instances that have the following instance types: C1, CC1, CC2, CG1, CG2, CR1, CS1, G1, G2, HI1, HS1, M1, M2, M3, and T1.
	LoadBalancersConfig interface{} `field:"optional" json:"loadBalancersConfig" yaml:"loadBalancersConfig"`
	// The order of the launch template overrides to use in fulfilling On-Demand capacity.
	//
	// If you specify `lowestPrice` , Spot Fleet uses price to determine the order, launching the lowest price first. If you specify `prioritized` , Spot Fleet uses the priority that you assign to each Spot Fleet launch template override, launching the highest priority first. If you do not specify a value, Spot Fleet defaults to `lowestPrice` .
	OnDemandAllocationStrategy *string `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// The maximum amount per hour for On-Demand Instances that you're willing to pay.
	//
	// You can use the `onDemandMaxTotalPrice` parameter, the `spotMaxTotalPrice` parameter, or both parameters to ensure that your fleet cost does not exceed your budget. If you set a maximum price per hour for the On-Demand Instances and Spot Instances in your request, Spot Fleet will launch instances until it reaches the maximum amount you're willing to pay. When the maximum amount you're willing to pay is reached, the fleet stops launching instances even if it hasn’t met the target capacity.
	OnDemandMaxTotalPrice *string `field:"optional" json:"onDemandMaxTotalPrice" yaml:"onDemandMaxTotalPrice"`
	// The number of On-Demand units to request.
	//
	// You can choose to set the target capacity in terms of instances or a performance characteristic that is important to your application workload, such as vCPUs, memory, or I/O. If the request type is `maintain` , you can specify a target capacity of 0 and add capacity later.
	OnDemandTargetCapacity *float64 `field:"optional" json:"onDemandTargetCapacity" yaml:"onDemandTargetCapacity"`
	// Indicates whether Spot Fleet should replace unhealthy instances.
	ReplaceUnhealthyInstances interface{} `field:"optional" json:"replaceUnhealthyInstances" yaml:"replaceUnhealthyInstances"`
	// The strategies for managing your Spot Instances that are at an elevated risk of being interrupted.
	SpotMaintenanceStrategies interface{} `field:"optional" json:"spotMaintenanceStrategies" yaml:"spotMaintenanceStrategies"`
	// The maximum amount per hour for Spot Instances that you're willing to pay.
	//
	// You can use the `spotdMaxTotalPrice` parameter, the `onDemandMaxTotalPrice` parameter, or both parameters to ensure that your fleet cost does not exceed your budget. If you set a maximum price per hour for the On-Demand Instances and Spot Instances in your request, Spot Fleet will launch instances until it reaches the maximum amount you're willing to pay. When the maximum amount you're willing to pay is reached, the fleet stops launching instances even if it hasn’t met the target capacity.
	SpotMaxTotalPrice *string `field:"optional" json:"spotMaxTotalPrice" yaml:"spotMaxTotalPrice"`
	// The maximum price per unit hour that you are willing to pay for a Spot Instance.
	//
	// The default is the On-Demand price.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// `CfnSpotFleet.SpotFleetRequestConfigDataProperty.TagSpecifications`.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The unit for the target capacity.
	//
	// Default: `units` (translates to number of instances).
	TargetCapacityUnitType *string `field:"optional" json:"targetCapacityUnitType" yaml:"targetCapacityUnitType"`
	// Indicates whether running Spot Instances are terminated when the Spot Fleet request expires.
	TerminateInstancesWithExpiration interface{} `field:"optional" json:"terminateInstancesWithExpiration" yaml:"terminateInstancesWithExpiration"`
	// The type of request.
	//
	// Indicates whether the Spot Fleet only requests the target capacity or also attempts to maintain it. When this value is `request` , the Spot Fleet only places the required requests. It does not attempt to replenish Spot Instances if capacity is diminished, nor does it submit requests in alternative Spot pools if capacity is not available. When this value is `maintain` , the Spot Fleet maintains the target capacity. The Spot Fleet places the required requests to meet capacity and automatically replenishes any interrupted instances. Default: `maintain` . `instant` is listed but is not used by Spot Fleet.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The start date and time of the request, in UTC format ( *YYYY* - *MM* - *DD* T *HH* : *MM* : *SS* Z).
	//
	// By default, Amazon EC2 starts fulfilling the request immediately.
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// The end date and time of the request, in UTC format ( *YYYY* - *MM* - *DD* T *HH* : *MM* : *SS* Z).
	//
	// After the end date and time, no new Spot Instance requests are placed or able to fulfill the request. If no value is specified, the Spot Fleet request remains until you cancel it.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}


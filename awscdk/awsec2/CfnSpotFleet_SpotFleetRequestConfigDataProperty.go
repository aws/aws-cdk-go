package awsec2


// Specifies the configuration of a Spot Fleet request.
//
// For more information, see [Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html) in the *Amazon EC2 User Guide* .
//
// You must specify either `LaunchSpecifications` or `LaunchTemplateConfigs` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotFleetRequestConfigDataProperty := &SpotFleetRequestConfigDataProperty{
//   	IamFleetRole: jsii.String("iamFleetRole"),
//   	TargetCapacity: jsii.Number(123),
//
//   	// the properties below are optional
//   	AllocationStrategy: jsii.String("allocationStrategy"),
//   	Context: jsii.String("context"),
//   	ExcessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   	InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   	InstancePoolsToUseCount: jsii.Number(123),
//   	LaunchSpecifications: []interface{}{
//   		&SpotFleetLaunchSpecificationProperty{
//   			ImageId: jsii.String("imageId"),
//
//   			// the properties below are optional
//   			BlockDeviceMappings: []interface{}{
//   				&BlockDeviceMappingProperty{
//   					DeviceName: jsii.String("deviceName"),
//
//   					// the properties below are optional
//   					Ebs: &EbsBlockDeviceProperty{
//   						DeleteOnTermination: jsii.Boolean(false),
//   						Encrypted: jsii.Boolean(false),
//   						Iops: jsii.Number(123),
//   						SnapshotId: jsii.String("snapshotId"),
//   						VolumeSize: jsii.Number(123),
//   						VolumeType: jsii.String("volumeType"),
//   					},
//   					NoDevice: jsii.String("noDevice"),
//   					VirtualName: jsii.String("virtualName"),
//   				},
//   			},
//   			EbsOptimized: jsii.Boolean(false),
//   			IamInstanceProfile: &IamInstanceProfileSpecificationProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			InstanceRequirements: &InstanceRequirementsRequestProperty{
//   				AcceleratorCount: &AcceleratorCountRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorManufacturers: []*string{
//   					jsii.String("acceleratorManufacturers"),
//   				},
//   				AcceleratorNames: []*string{
//   					jsii.String("acceleratorNames"),
//   				},
//   				AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorTypes: []*string{
//   					jsii.String("acceleratorTypes"),
//   				},
//   				AllowedInstanceTypes: []*string{
//   					jsii.String("allowedInstanceTypes"),
//   				},
//   				BareMetal: jsii.String("bareMetal"),
//   				BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				BurstablePerformance: jsii.String("burstablePerformance"),
//   				CpuManufacturers: []*string{
//   					jsii.String("cpuManufacturers"),
//   				},
//   				ExcludedInstanceTypes: []*string{
//   					jsii.String("excludedInstanceTypes"),
//   				},
//   				InstanceGenerations: []*string{
//   					jsii.String("instanceGenerations"),
//   				},
//   				LocalStorage: jsii.String("localStorage"),
//   				LocalStorageTypes: []*string{
//   					jsii.String("localStorageTypes"),
//   				},
//   				MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   				MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				MemoryMiB: &MemoryMiBRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				RequireHibernateSupport: jsii.Boolean(false),
//   				SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				VCpuCount: &VCpuCountRangeRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   			},
//   			InstanceType: jsii.String("instanceType"),
//   			KernelId: jsii.String("kernelId"),
//   			KeyName: jsii.String("keyName"),
//   			Monitoring: &SpotFleetMonitoringProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			NetworkInterfaces: []interface{}{
//   				&InstanceNetworkInterfaceSpecificationProperty{
//   					AssociatePublicIpAddress: jsii.Boolean(false),
//   					DeleteOnTermination: jsii.Boolean(false),
//   					Description: jsii.String("description"),
//   					DeviceIndex: jsii.Number(123),
//   					Groups: []*string{
//   						jsii.String("groups"),
//   					},
//   					Ipv6AddressCount: jsii.Number(123),
//   					Ipv6Addresses: []interface{}{
//   						&InstanceIpv6AddressProperty{
//   							Ipv6Address: jsii.String("ipv6Address"),
//   						},
//   					},
//   					NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   					PrivateIpAddresses: []interface{}{
//   						&PrivateIpAddressSpecificationProperty{
//   							PrivateIpAddress: jsii.String("privateIpAddress"),
//
//   							// the properties below are optional
//   							Primary: jsii.Boolean(false),
//   						},
//   					},
//   					SecondaryPrivateIpAddressCount: jsii.Number(123),
//   					SubnetId: jsii.String("subnetId"),
//   				},
//   			},
//   			Placement: &SpotPlacementProperty{
//   				AvailabilityZone: jsii.String("availabilityZone"),
//   				GroupName: jsii.String("groupName"),
//   				Tenancy: jsii.String("tenancy"),
//   			},
//   			RamdiskId: jsii.String("ramdiskId"),
//   			SecurityGroups: []interface{}{
//   				&GroupIdentifierProperty{
//   					GroupId: jsii.String("groupId"),
//   				},
//   			},
//   			SpotPrice: jsii.String("spotPrice"),
//   			SubnetId: jsii.String("subnetId"),
//   			TagSpecifications: []interface{}{
//   				&SpotFleetTagSpecificationProperty{
//   					ResourceType: jsii.String("resourceType"),
//   					Tags: []cfnTag{
//   						&cfnTag{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			UserData: jsii.String("userData"),
//   			WeightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	LaunchTemplateConfigs: []interface{}{
//   		&LaunchTemplateConfigProperty{
//   			LaunchTemplateSpecification: &FleetLaunchTemplateSpecificationProperty{
//   				Version: jsii.String("version"),
//
//   				// the properties below are optional
//   				LaunchTemplateId: jsii.String("launchTemplateId"),
//   				LaunchTemplateName: jsii.String("launchTemplateName"),
//   			},
//   			Overrides: []interface{}{
//   				&LaunchTemplateOverridesProperty{
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
//   					Priority: jsii.Number(123),
//   					SpotPrice: jsii.String("spotPrice"),
//   					SubnetId: jsii.String("subnetId"),
//   					WeightedCapacity: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	LoadBalancersConfig: &LoadBalancersConfigProperty{
//   		ClassicLoadBalancersConfig: &ClassicLoadBalancersConfigProperty{
//   			ClassicLoadBalancers: []interface{}{
//   				&ClassicLoadBalancerProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		TargetGroupsConfig: &TargetGroupsConfigProperty{
//   			TargetGroups: []interface{}{
//   				&TargetGroupProperty{
//   					Arn: jsii.String("arn"),
//   				},
//   			},
//   		},
//   	},
//   	OnDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   	OnDemandMaxTotalPrice: jsii.String("onDemandMaxTotalPrice"),
//   	OnDemandTargetCapacity: jsii.Number(123),
//   	ReplaceUnhealthyInstances: jsii.Boolean(false),
//   	SpotMaintenanceStrategies: &SpotMaintenanceStrategiesProperty{
//   		CapacityRebalance: &SpotCapacityRebalanceProperty{
//   			ReplacementStrategy: jsii.String("replacementStrategy"),
//   			TerminationDelay: jsii.Number(123),
//   		},
//   	},
//   	SpotMaxTotalPrice: jsii.String("spotMaxTotalPrice"),
//   	SpotPrice: jsii.String("spotPrice"),
//   	TagSpecifications: []interface{}{
//   		&SpotFleetTagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []*cfnTag{
//   				&cfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	TargetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   	TerminateInstancesWithExpiration: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   	ValidFrom: jsii.String("validFrom"),
//   	ValidUntil: jsii.String("validUntil"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html
//
type CfnSpotFleet_SpotFleetRequestConfigDataProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM) role that grants the Spot Fleet the permission to request, launch, terminate, and tag instances on your behalf.
	//
	// For more information, see [Spot Fleet Prerequisites](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-requests.html#spot-fleet-prerequisites) in the *Amazon EC2 User Guide for Linux Instances* . Spot Fleet can terminate Spot Instances on your behalf when you cancel its Spot Fleet request or when the Spot Fleet request expires, if you set `TerminateInstancesWithExpiration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-iamfleetrole
	//
	IamFleetRole *string `field:"required" json:"iamFleetRole" yaml:"iamFleetRole"`
	// The number of units to request for the Spot Fleet.
	//
	// You can choose to set the target capacity in terms of instances or a performance characteristic that is important to your application workload, such as vCPUs, memory, or I/O. If the request type is `maintain` , you can specify a target capacity of 0 and add capacity later.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-targetcapacity
	//
	TargetCapacity *float64 `field:"required" json:"targetCapacity" yaml:"targetCapacity"`
	// The strategy that determines how to allocate the target Spot Instance capacity across the Spot Instance pools specified by the Spot Fleet launch configuration.
	//
	// For more information, see [Allocation strategies for Spot Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-allocation-strategy.html) in the *Amazon EC2 User Guide* .
	//
	// - **priceCapacityOptimized (recommended)** - Spot Fleet identifies the pools with the highest capacity availability for the number of instances that are launching. This means that we will request Spot Instances from the pools that we believe have the lowest chance of interruption in the near term. Spot Fleet then requests Spot Instances from the lowest priced of these pools.
	// - **capacityOptimized** - Spot Fleet identifies the pools with the highest capacity availability for the number of instances that are launching. This means that we will request Spot Instances from the pools that we believe have the lowest chance of interruption in the near term. To give certain instance types a higher chance of launching first, use `capacityOptimizedPrioritized` . Set a priority for each instance type by using the `Priority` parameter for `LaunchTemplateOverrides` . You can assign the same priority to different `LaunchTemplateOverrides` . EC2 implements the priorities on a best-effort basis, but optimizes for capacity first. `capacityOptimizedPrioritized` is supported only if your Spot Fleet uses a launch template. Note that if the `OnDemandAllocationStrategy` is set to `prioritized` , the same priority is applied when fulfilling On-Demand capacity.
	// - **diversified** - Spot Fleet requests instances from all of the Spot Instance pools that you specify.
	// - **lowestPrice** - Spot Fleet requests instances from the lowest priced Spot Instance pool that has available capacity. If the lowest priced pool doesn't have available capacity, the Spot Instances come from the next lowest priced pool that has available capacity. If a pool runs out of capacity before fulfilling your desired capacity, Spot Fleet will continue to fulfill your request by drawing from the next lowest priced pool. To ensure that your desired capacity is met, you might receive Spot Instances from several pools. Because this strategy only considers instance price and not capacity availability, it might lead to high interruption rates.
	//
	// Default: `lowestPrice`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-allocationstrategy
	//
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Reserved.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-context
	//
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Indicates whether running Spot Instances should be terminated if you decrease the target capacity of the Spot Fleet request below the current size of the Spot Fleet.
	//
	// Supported only for fleets of type `maintain` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-excesscapacityterminationpolicy
	//
	ExcessCapacityTerminationPolicy *string `field:"optional" json:"excessCapacityTerminationPolicy" yaml:"excessCapacityTerminationPolicy"`
	// The behavior when a Spot Instance is interrupted.
	//
	// The default is `terminate` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-instanceinterruptionbehavior
	//
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// The number of Spot pools across which to allocate your target Spot capacity.
	//
	// Valid only when Spot *AllocationStrategy* is set to `lowest-price` . Spot Fleet selects the cheapest Spot pools and evenly allocates your target Spot capacity across the number of Spot pools that you specify.
	//
	// Note that Spot Fleet attempts to draw Spot Instances from the number of pools that you specify on a best effort basis. If a pool runs out of Spot capacity before fulfilling your target capacity, Spot Fleet will continue to fulfill your request by drawing from the next cheapest pool. To ensure that your target capacity is met, you might receive Spot Instances from more than the number of pools that you specified. Similarly, if most of the pools have no Spot capacity, you might receive your full target capacity from fewer than the number of pools that you specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-instancepoolstousecount
	//
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// The launch specifications for the Spot Fleet request.
	//
	// If you specify `LaunchSpecifications` , you can't specify `LaunchTemplateConfigs` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications
	//
	LaunchSpecifications interface{} `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// The launch template and overrides.
	//
	// If you specify `LaunchTemplateConfigs` , you can't specify `LaunchSpecifications` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-launchtemplateconfigs
	//
	LaunchTemplateConfigs interface{} `field:"optional" json:"launchTemplateConfigs" yaml:"launchTemplateConfigs"`
	// One or more Classic Load Balancers and target groups to attach to the Spot Fleet request.
	//
	// Spot Fleet registers the running Spot Instances with the specified Classic Load Balancers and target groups.
	//
	// With Network Load Balancers, Spot Fleet cannot register instances that have the following instance types: C1, CC1, CC2, CG1, CG2, CR1, CS1, G1, G2, HI1, HS1, M1, M2, M3, and T1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-loadbalancersconfig
	//
	LoadBalancersConfig interface{} `field:"optional" json:"loadBalancersConfig" yaml:"loadBalancersConfig"`
	// The order of the launch template overrides to use in fulfilling On-Demand capacity.
	//
	// If you specify `lowestPrice` , Spot Fleet uses price to determine the order, launching the lowest price first. If you specify `prioritized` , Spot Fleet uses the priority that you assign to each Spot Fleet launch template override, launching the highest priority first. If you do not specify a value, Spot Fleet defaults to `lowestPrice` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-ondemandallocationstrategy
	//
	OnDemandAllocationStrategy *string `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// The maximum amount per hour for On-Demand Instances that you're willing to pay.
	//
	// You can use the `onDemandMaxTotalPrice` parameter, the `spotMaxTotalPrice` parameter, or both parameters to ensure that your fleet cost does not exceed your budget. If you set a maximum price per hour for the On-Demand Instances and Spot Instances in your request, Spot Fleet will launch instances until it reaches the maximum amount you're willing to pay. When the maximum amount you're willing to pay is reached, the fleet stops launching instances even if it hasn’t met the target capacity.
	//
	// > If your fleet includes T instances that are configured as `unlimited` , and if their average CPU usage exceeds the baseline utilization, you will incur a charge for surplus credits. The `onDemandMaxTotalPrice` does not account for surplus credits, and, if you use surplus credits, your final cost might be higher than what you specified for `onDemandMaxTotalPrice` . For more information, see [Surplus credits can incur charges](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#unlimited-mode-surplus-credits) in the *EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-ondemandmaxtotalprice
	//
	OnDemandMaxTotalPrice *string `field:"optional" json:"onDemandMaxTotalPrice" yaml:"onDemandMaxTotalPrice"`
	// The number of On-Demand units to request.
	//
	// You can choose to set the target capacity in terms of instances or a performance characteristic that is important to your application workload, such as vCPUs, memory, or I/O. If the request type is `maintain` , you can specify a target capacity of 0 and add capacity later.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-ondemandtargetcapacity
	//
	OnDemandTargetCapacity *float64 `field:"optional" json:"onDemandTargetCapacity" yaml:"onDemandTargetCapacity"`
	// Indicates whether Spot Fleet should replace unhealthy instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-replaceunhealthyinstances
	//
	ReplaceUnhealthyInstances interface{} `field:"optional" json:"replaceUnhealthyInstances" yaml:"replaceUnhealthyInstances"`
	// The strategies for managing your Spot Instances that are at an elevated risk of being interrupted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-spotmaintenancestrategies
	//
	SpotMaintenanceStrategies interface{} `field:"optional" json:"spotMaintenanceStrategies" yaml:"spotMaintenanceStrategies"`
	// The maximum amount per hour for Spot Instances that you're willing to pay.
	//
	// You can use the `spotMaxTotalPrice` parameter, the `onDemandMaxTotalPrice` parameter, or both parameters to ensure that your fleet cost does not exceed your budget. If you set a maximum price per hour for the On-Demand Instances and Spot Instances in your request, Spot Fleet will launch instances until it reaches the maximum amount you're willing to pay. When the maximum amount you're willing to pay is reached, the fleet stops launching instances even if it hasn’t met the target capacity.
	//
	// > If your fleet includes T instances that are configured as `unlimited` , and if their average CPU usage exceeds the baseline utilization, you will incur a charge for surplus credits. The `spotMaxTotalPrice` does not account for surplus credits, and, if you use surplus credits, your final cost might be higher than what you specified for `spotMaxTotalPrice` . For more information, see [Surplus credits can incur charges](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#unlimited-mode-surplus-credits) in the *EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-spotmaxtotalprice
	//
	SpotMaxTotalPrice *string `field:"optional" json:"spotMaxTotalPrice" yaml:"spotMaxTotalPrice"`
	// The maximum price per unit hour that you are willing to pay for a Spot Instance.
	//
	// We do not recommend using this parameter because it can lead to increased interruptions. If you do not specify this parameter, you will pay the current Spot price.
	//
	// > If you specify a maximum price, your instances will be interrupted more frequently than if you do not specify this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-spotprice
	//
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// The key-value pair for tagging the Spot Fleet request on creation.
	//
	// The value for `ResourceType` must be `spot-fleet-request` , otherwise the Spot Fleet request fails. To tag instances at launch, specify the tags in the [launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html#create-launch-template) (valid only if you use `LaunchTemplateConfigs` ) or in the `[SpotFleetTagSpecification](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotFleetTagSpecification.html)` (valid only if you use `LaunchSpecifications` ). For information about tagging after launch, see [Tag your resources](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#tag-resources) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-tagspecifications
	//
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The unit for the target capacity. You can specify this parameter only when using attribute-based instance type selection.
	//
	// Default: `units` (the number of instances).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-targetcapacityunittype
	//
	TargetCapacityUnitType *string `field:"optional" json:"targetCapacityUnitType" yaml:"targetCapacityUnitType"`
	// Indicates whether running Spot Instances are terminated when the Spot Fleet request expires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-terminateinstanceswithexpiration
	//
	TerminateInstancesWithExpiration interface{} `field:"optional" json:"terminateInstancesWithExpiration" yaml:"terminateInstancesWithExpiration"`
	// The type of request.
	//
	// Indicates whether the Spot Fleet only requests the target capacity or also attempts to maintain it. When this value is `request` , the Spot Fleet only places the required requests. It does not attempt to replenish Spot Instances if capacity is diminished, nor does it submit requests in alternative Spot pools if capacity is not available. When this value is `maintain` , the Spot Fleet maintains the target capacity. The Spot Fleet places the required requests to meet capacity and automatically replenishes any interrupted instances. Default: `maintain` . `instant` is listed but is not used by Spot Fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The start date and time of the request, in UTC format ( *YYYY* - *MM* - *DD* T *HH* : *MM* : *SS* Z).
	//
	// By default, Amazon EC2 starts fulfilling the request immediately.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-validfrom
	//
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// The end date and time of the request, in UTC format ( *YYYY* - *MM* - *DD* T *HH* : *MM* : *SS* Z).
	//
	// After the end date and time, no new Spot Instance requests are placed or able to fulfill the request. If no value is specified, the Spot Fleet request remains until you cancel it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata-validuntil
	//
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}


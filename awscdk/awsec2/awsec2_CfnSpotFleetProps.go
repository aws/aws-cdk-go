package awsec2


// Properties for defining a `CfnSpotFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSpotFleetProps := &cfnSpotFleetProps{
//   	spotFleetRequestConfigData: &spotFleetRequestConfigDataProperty{
//   		iamFleetRole: jsii.String("iamFleetRole"),
//   		targetCapacity: jsii.Number(123),
//
//   		// the properties below are optional
//   		allocationStrategy: jsii.String("allocationStrategy"),
//   		context: jsii.String("context"),
//   		excessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   		instanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		instancePoolsToUseCount: jsii.Number(123),
//   		launchSpecifications: []interface{}{
//   			&spotFleetLaunchSpecificationProperty{
//   				imageId: jsii.String("imageId"),
//
//   				// the properties below are optional
//   				blockDeviceMappings: []interface{}{
//   					&blockDeviceMappingProperty{
//   						deviceName: jsii.String("deviceName"),
//
//   						// the properties below are optional
//   						ebs: &ebsBlockDeviceProperty{
//   							deleteOnTermination: jsii.Boolean(false),
//   							encrypted: jsii.Boolean(false),
//   							iops: jsii.Number(123),
//   							snapshotId: jsii.String("snapshotId"),
//   							volumeSize: jsii.Number(123),
//   							volumeType: jsii.String("volumeType"),
//   						},
//   						noDevice: jsii.String("noDevice"),
//   						virtualName: jsii.String("virtualName"),
//   					},
//   				},
//   				ebsOptimized: jsii.Boolean(false),
//   				iamInstanceProfile: &iamInstanceProfileSpecificationProperty{
//   					arn: jsii.String("arn"),
//   				},
//   				instanceRequirements: &instanceRequirementsRequestProperty{
//   					acceleratorCount: &acceleratorCountRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					acceleratorManufacturers: []*string{
//   						jsii.String("acceleratorManufacturers"),
//   					},
//   					acceleratorNames: []*string{
//   						jsii.String("acceleratorNames"),
//   					},
//   					acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					acceleratorTypes: []*string{
//   						jsii.String("acceleratorTypes"),
//   					},
//   					allowedInstanceTypes: []*string{
//   						jsii.String("allowedInstanceTypes"),
//   					},
//   					bareMetal: jsii.String("bareMetal"),
//   					baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					burstablePerformance: jsii.String("burstablePerformance"),
//   					cpuManufacturers: []*string{
//   						jsii.String("cpuManufacturers"),
//   					},
//   					excludedInstanceTypes: []*string{
//   						jsii.String("excludedInstanceTypes"),
//   					},
//   					instanceGenerations: []*string{
//   						jsii.String("instanceGenerations"),
//   					},
//   					localStorage: jsii.String("localStorage"),
//   					localStorageTypes: []*string{
//   						jsii.String("localStorageTypes"),
//   					},
//   					memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					memoryMiB: &memoryMiBRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					networkBandwidthGbps: &networkBandwidthGbpsRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					requireHibernateSupport: jsii.Boolean(false),
//   					spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					vCpuCount: &vCpuCountRangeRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   				},
//   				instanceType: jsii.String("instanceType"),
//   				kernelId: jsii.String("kernelId"),
//   				keyName: jsii.String("keyName"),
//   				monitoring: &spotFleetMonitoringProperty{
//   					enabled: jsii.Boolean(false),
//   				},
//   				networkInterfaces: []interface{}{
//   					&instanceNetworkInterfaceSpecificationProperty{
//   						associatePublicIpAddress: jsii.Boolean(false),
//   						deleteOnTermination: jsii.Boolean(false),
//   						description: jsii.String("description"),
//   						deviceIndex: jsii.Number(123),
//   						groups: []*string{
//   							jsii.String("groups"),
//   						},
//   						ipv6AddressCount: jsii.Number(123),
//   						ipv6Addresses: []interface{}{
//   							&instanceIpv6AddressProperty{
//   								ipv6Address: jsii.String("ipv6Address"),
//   							},
//   						},
//   						networkInterfaceId: jsii.String("networkInterfaceId"),
//   						privateIpAddresses: []interface{}{
//   							&privateIpAddressSpecificationProperty{
//   								privateIpAddress: jsii.String("privateIpAddress"),
//
//   								// the properties below are optional
//   								primary: jsii.Boolean(false),
//   							},
//   						},
//   						secondaryPrivateIpAddressCount: jsii.Number(123),
//   						subnetId: jsii.String("subnetId"),
//   					},
//   				},
//   				placement: &spotPlacementProperty{
//   					availabilityZone: jsii.String("availabilityZone"),
//   					groupName: jsii.String("groupName"),
//   					tenancy: jsii.String("tenancy"),
//   				},
//   				ramdiskId: jsii.String("ramdiskId"),
//   				securityGroups: []interface{}{
//   					&groupIdentifierProperty{
//   						groupId: jsii.String("groupId"),
//   					},
//   				},
//   				spotPrice: jsii.String("spotPrice"),
//   				subnetId: jsii.String("subnetId"),
//   				tagSpecifications: []interface{}{
//   					&spotFleetTagSpecificationProperty{
//   						resourceType: jsii.String("resourceType"),
//   						tags: []cfnTag{
//   							&cfnTag{
//   								key: jsii.String("key"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				userData: jsii.String("userData"),
//   				weightedCapacity: jsii.Number(123),
//   			},
//   		},
//   		launchTemplateConfigs: []interface{}{
//   			&launchTemplateConfigProperty{
//   				launchTemplateSpecification: &fleetLaunchTemplateSpecificationProperty{
//   					version: jsii.String("version"),
//
//   					// the properties below are optional
//   					launchTemplateId: jsii.String("launchTemplateId"),
//   					launchTemplateName: jsii.String("launchTemplateName"),
//   				},
//   				overrides: []interface{}{
//   					&launchTemplateOverridesProperty{
//   						availabilityZone: jsii.String("availabilityZone"),
//   						instanceRequirements: &instanceRequirementsRequestProperty{
//   							acceleratorCount: &acceleratorCountRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							acceleratorManufacturers: []*string{
//   								jsii.String("acceleratorManufacturers"),
//   							},
//   							acceleratorNames: []*string{
//   								jsii.String("acceleratorNames"),
//   							},
//   							acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							acceleratorTypes: []*string{
//   								jsii.String("acceleratorTypes"),
//   							},
//   							allowedInstanceTypes: []*string{
//   								jsii.String("allowedInstanceTypes"),
//   							},
//   							bareMetal: jsii.String("bareMetal"),
//   							baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							burstablePerformance: jsii.String("burstablePerformance"),
//   							cpuManufacturers: []*string{
//   								jsii.String("cpuManufacturers"),
//   							},
//   							excludedInstanceTypes: []*string{
//   								jsii.String("excludedInstanceTypes"),
//   							},
//   							instanceGenerations: []*string{
//   								jsii.String("instanceGenerations"),
//   							},
//   							localStorage: jsii.String("localStorage"),
//   							localStorageTypes: []*string{
//   								jsii.String("localStorageTypes"),
//   							},
//   							memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							memoryMiB: &memoryMiBRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							networkBandwidthGbps: &networkBandwidthGbpsRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   							requireHibernateSupport: jsii.Boolean(false),
//   							spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   							totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   							vCpuCount: &vCpuCountRangeRequestProperty{
//   								max: jsii.Number(123),
//   								min: jsii.Number(123),
//   							},
//   						},
//   						instanceType: jsii.String("instanceType"),
//   						priority: jsii.Number(123),
//   						spotPrice: jsii.String("spotPrice"),
//   						subnetId: jsii.String("subnetId"),
//   						weightedCapacity: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		loadBalancersConfig: &loadBalancersConfigProperty{
//   			classicLoadBalancersConfig: &classicLoadBalancersConfigProperty{
//   				classicLoadBalancers: []interface{}{
//   					&classicLoadBalancerProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			targetGroupsConfig: &targetGroupsConfigProperty{
//   				targetGroups: []interface{}{
//   					&targetGroupProperty{
//   						arn: jsii.String("arn"),
//   					},
//   				},
//   			},
//   		},
//   		onDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   		onDemandMaxTotalPrice: jsii.String("onDemandMaxTotalPrice"),
//   		onDemandTargetCapacity: jsii.Number(123),
//   		replaceUnhealthyInstances: jsii.Boolean(false),
//   		spotMaintenanceStrategies: &spotMaintenanceStrategiesProperty{
//   			capacityRebalance: &spotCapacityRebalanceProperty{
//   				replacementStrategy: jsii.String("replacementStrategy"),
//   				terminationDelay: jsii.Number(123),
//   			},
//   		},
//   		spotMaxTotalPrice: jsii.String("spotMaxTotalPrice"),
//   		spotPrice: jsii.String("spotPrice"),
//   		tagSpecifications: []interface{}{
//   			&spotFleetTagSpecificationProperty{
//   				resourceType: jsii.String("resourceType"),
//   				tags: []*cfnTag{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		targetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   		terminateInstancesWithExpiration: jsii.Boolean(false),
//   		type: jsii.String("type"),
//   		validFrom: jsii.String("validFrom"),
//   		validUntil: jsii.String("validUntil"),
//   	},
//   }
//
type CfnSpotFleetProps struct {
	// Describes the configuration of a Spot Fleet request.
	SpotFleetRequestConfigData interface{} `field:"required" json:"spotFleetRequestConfigData" yaml:"spotFleetRequestConfigData"`
}


package awsec2


// Properties for defining a `CfnSpotFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSpotFleetProps := &CfnSpotFleetProps{
//   	SpotFleetRequestConfigData: &SpotFleetRequestConfigDataProperty{
//   		IamFleetRole: jsii.String("iamFleetRole"),
//   		TargetCapacity: jsii.Number(123),
//
//   		// the properties below are optional
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		Context: jsii.String("context"),
//   		ExcessCapacityTerminationPolicy: jsii.String("excessCapacityTerminationPolicy"),
//   		InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   		InstancePoolsToUseCount: jsii.Number(123),
//   		LaunchSpecifications: []interface{}{
//   			&SpotFleetLaunchSpecificationProperty{
//   				ImageId: jsii.String("imageId"),
//
//   				// the properties below are optional
//   				BlockDeviceMappings: []interface{}{
//   					&BlockDeviceMappingProperty{
//   						DeviceName: jsii.String("deviceName"),
//
//   						// the properties below are optional
//   						Ebs: &EbsBlockDeviceProperty{
//   							DeleteOnTermination: jsii.Boolean(false),
//   							Encrypted: jsii.Boolean(false),
//   							Iops: jsii.Number(123),
//   							SnapshotId: jsii.String("snapshotId"),
//   							VolumeSize: jsii.Number(123),
//   							VolumeType: jsii.String("volumeType"),
//   						},
//   						NoDevice: jsii.String("noDevice"),
//   						VirtualName: jsii.String("virtualName"),
//   					},
//   				},
//   				EbsOptimized: jsii.Boolean(false),
//   				IamInstanceProfile: &IamInstanceProfileSpecificationProperty{
//   					Arn: jsii.String("arn"),
//   				},
//   				InstanceRequirements: &InstanceRequirementsRequestProperty{
//   					AcceleratorCount: &AcceleratorCountRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					AcceleratorManufacturers: []*string{
//   						jsii.String("acceleratorManufacturers"),
//   					},
//   					AcceleratorNames: []*string{
//   						jsii.String("acceleratorNames"),
//   					},
//   					AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					AcceleratorTypes: []*string{
//   						jsii.String("acceleratorTypes"),
//   					},
//   					AllowedInstanceTypes: []*string{
//   						jsii.String("allowedInstanceTypes"),
//   					},
//   					BareMetal: jsii.String("bareMetal"),
//   					BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					BurstablePerformance: jsii.String("burstablePerformance"),
//   					CpuManufacturers: []*string{
//   						jsii.String("cpuManufacturers"),
//   					},
//   					ExcludedInstanceTypes: []*string{
//   						jsii.String("excludedInstanceTypes"),
//   					},
//   					InstanceGenerations: []*string{
//   						jsii.String("instanceGenerations"),
//   					},
//   					LocalStorage: jsii.String("localStorage"),
//   					LocalStorageTypes: []*string{
//   						jsii.String("localStorageTypes"),
//   					},
//   					MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   					MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					MemoryMiB: &MemoryMiBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					RequireHibernateSupport: jsii.Boolean(false),
//   					SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					VCpuCount: &VCpuCountRangeRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   				},
//   				InstanceType: jsii.String("instanceType"),
//   				KernelId: jsii.String("kernelId"),
//   				KeyName: jsii.String("keyName"),
//   				Monitoring: &SpotFleetMonitoringProperty{
//   					Enabled: jsii.Boolean(false),
//   				},
//   				NetworkInterfaces: []interface{}{
//   					&InstanceNetworkInterfaceSpecificationProperty{
//   						AssociatePublicIpAddress: jsii.Boolean(false),
//   						DeleteOnTermination: jsii.Boolean(false),
//   						Description: jsii.String("description"),
//   						DeviceIndex: jsii.Number(123),
//   						Groups: []*string{
//   							jsii.String("groups"),
//   						},
//   						Ipv6AddressCount: jsii.Number(123),
//   						Ipv6Addresses: []interface{}{
//   							&InstanceIpv6AddressProperty{
//   								Ipv6Address: jsii.String("ipv6Address"),
//   							},
//   						},
//   						NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   						PrivateIpAddresses: []interface{}{
//   							&PrivateIpAddressSpecificationProperty{
//   								PrivateIpAddress: jsii.String("privateIpAddress"),
//
//   								// the properties below are optional
//   								Primary: jsii.Boolean(false),
//   							},
//   						},
//   						SecondaryPrivateIpAddressCount: jsii.Number(123),
//   						SubnetId: jsii.String("subnetId"),
//   					},
//   				},
//   				Placement: &SpotPlacementProperty{
//   					AvailabilityZone: jsii.String("availabilityZone"),
//   					GroupName: jsii.String("groupName"),
//   					Tenancy: jsii.String("tenancy"),
//   				},
//   				RamdiskId: jsii.String("ramdiskId"),
//   				SecurityGroups: []interface{}{
//   					&GroupIdentifierProperty{
//   						GroupId: jsii.String("groupId"),
//   					},
//   				},
//   				SpotPrice: jsii.String("spotPrice"),
//   				SubnetId: jsii.String("subnetId"),
//   				TagSpecifications: []interface{}{
//   					&SpotFleetTagSpecificationProperty{
//   						ResourceType: jsii.String("resourceType"),
//   						Tags: []cfnTag{
//   							&cfnTag{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				UserData: jsii.String("userData"),
//   				WeightedCapacity: jsii.Number(123),
//   			},
//   		},
//   		LaunchTemplateConfigs: []interface{}{
//   			&LaunchTemplateConfigProperty{
//   				LaunchTemplateSpecification: &FleetLaunchTemplateSpecificationProperty{
//   					Version: jsii.String("version"),
//
//   					// the properties below are optional
//   					LaunchTemplateId: jsii.String("launchTemplateId"),
//   					LaunchTemplateName: jsii.String("launchTemplateName"),
//   				},
//   				Overrides: []interface{}{
//   					&LaunchTemplateOverridesProperty{
//   						AvailabilityZone: jsii.String("availabilityZone"),
//   						InstanceRequirements: &InstanceRequirementsRequestProperty{
//   							AcceleratorCount: &AcceleratorCountRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							AcceleratorManufacturers: []*string{
//   								jsii.String("acceleratorManufacturers"),
//   							},
//   							AcceleratorNames: []*string{
//   								jsii.String("acceleratorNames"),
//   							},
//   							AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							AcceleratorTypes: []*string{
//   								jsii.String("acceleratorTypes"),
//   							},
//   							AllowedInstanceTypes: []*string{
//   								jsii.String("allowedInstanceTypes"),
//   							},
//   							BareMetal: jsii.String("bareMetal"),
//   							BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							BurstablePerformance: jsii.String("burstablePerformance"),
//   							CpuManufacturers: []*string{
//   								jsii.String("cpuManufacturers"),
//   							},
//   							ExcludedInstanceTypes: []*string{
//   								jsii.String("excludedInstanceTypes"),
//   							},
//   							InstanceGenerations: []*string{
//   								jsii.String("instanceGenerations"),
//   							},
//   							LocalStorage: jsii.String("localStorage"),
//   							LocalStorageTypes: []*string{
//   								jsii.String("localStorageTypes"),
//   							},
//   							MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   							MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							MemoryMiB: &MemoryMiBRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   							RequireHibernateSupport: jsii.Boolean(false),
//   							SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   							TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   							VCpuCount: &VCpuCountRangeRequestProperty{
//   								Max: jsii.Number(123),
//   								Min: jsii.Number(123),
//   							},
//   						},
//   						InstanceType: jsii.String("instanceType"),
//   						Priority: jsii.Number(123),
//   						SpotPrice: jsii.String("spotPrice"),
//   						SubnetId: jsii.String("subnetId"),
//   						WeightedCapacity: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		LoadBalancersConfig: &LoadBalancersConfigProperty{
//   			ClassicLoadBalancersConfig: &ClassicLoadBalancersConfigProperty{
//   				ClassicLoadBalancers: []interface{}{
//   					&ClassicLoadBalancerProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			TargetGroupsConfig: &TargetGroupsConfigProperty{
//   				TargetGroups: []interface{}{
//   					&TargetGroupProperty{
//   						Arn: jsii.String("arn"),
//   					},
//   				},
//   			},
//   		},
//   		OnDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   		OnDemandMaxTotalPrice: jsii.String("onDemandMaxTotalPrice"),
//   		OnDemandTargetCapacity: jsii.Number(123),
//   		ReplaceUnhealthyInstances: jsii.Boolean(false),
//   		SpotMaintenanceStrategies: &SpotMaintenanceStrategiesProperty{
//   			CapacityRebalance: &SpotCapacityRebalanceProperty{
//   				ReplacementStrategy: jsii.String("replacementStrategy"),
//   				TerminationDelay: jsii.Number(123),
//   			},
//   		},
//   		SpotMaxTotalPrice: jsii.String("spotMaxTotalPrice"),
//   		SpotPrice: jsii.String("spotPrice"),
//   		TagSpecifications: []interface{}{
//   			&SpotFleetTagSpecificationProperty{
//   				ResourceType: jsii.String("resourceType"),
//   				Tags: []*cfnTag{
//   					&cfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		TargetCapacityUnitType: jsii.String("targetCapacityUnitType"),
//   		TerminateInstancesWithExpiration: jsii.Boolean(false),
//   		Type: jsii.String("type"),
//   		ValidFrom: jsii.String("validFrom"),
//   		ValidUntil: jsii.String("validUntil"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html
//
type CfnSpotFleetProps struct {
	// Describes the configuration of a Spot Fleet request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata
	//
	SpotFleetRequestConfigData interface{} `field:"required" json:"spotFleetRequestConfigData" yaml:"spotFleetRequestConfigData"`
}


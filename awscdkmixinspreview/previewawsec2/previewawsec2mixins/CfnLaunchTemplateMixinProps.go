package previewawsec2mixins


// Properties for CfnLaunchTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLaunchTemplateMixinProps := &CfnLaunchTemplateMixinProps{
//   	LaunchTemplateData: &LaunchTemplateDataProperty{
//   		BlockDeviceMappings: []interface{}{
//   			&BlockDeviceMappingProperty{
//   				DeviceName: jsii.String("deviceName"),
//   				Ebs: &EbsProperty{
//   					DeleteOnTermination: jsii.Boolean(false),
//   					EbsCardIndex: jsii.Number(123),
//   					Encrypted: jsii.Boolean(false),
//   					Iops: jsii.Number(123),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					SnapshotId: jsii.String("snapshotId"),
//   					Throughput: jsii.Number(123),
//   					VolumeInitializationRate: jsii.Number(123),
//   					VolumeSize: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//   				},
//   				NoDevice: jsii.String("noDevice"),
//   				VirtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		CapacityReservationSpecification: &CapacityReservationSpecificationProperty{
//   			CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   			CapacityReservationTarget: &CapacityReservationTargetProperty{
//   				CapacityReservationId: jsii.String("capacityReservationId"),
//   				CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   			},
//   		},
//   		CpuOptions: &CpuOptionsProperty{
//   			AmdSevSnp: jsii.String("amdSevSnp"),
//   			CoreCount: jsii.Number(123),
//   			ThreadsPerCore: jsii.Number(123),
//   		},
//   		CreditSpecification: &CreditSpecificationProperty{
//   			CpuCredits: jsii.String("cpuCredits"),
//   		},
//   		DisableApiStop: jsii.Boolean(false),
//   		DisableApiTermination: jsii.Boolean(false),
//   		EbsOptimized: jsii.Boolean(false),
//   		ElasticGpuSpecifications: []interface{}{
//   			&ElasticGpuSpecificationProperty{
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		ElasticInferenceAccelerators: []interface{}{
//   			&LaunchTemplateElasticInferenceAcceleratorProperty{
//   				Count: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		EnclaveOptions: &EnclaveOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		HibernationOptions: &HibernationOptionsProperty{
//   			Configured: jsii.Boolean(false),
//   		},
//   		IamInstanceProfile: &IamInstanceProfileProperty{
//   			Arn: jsii.String("arn"),
//   			Name: jsii.String("name"),
//   		},
//   		ImageId: jsii.String("imageId"),
//   		InstanceInitiatedShutdownBehavior: jsii.String("instanceInitiatedShutdownBehavior"),
//   		InstanceMarketOptions: &InstanceMarketOptionsProperty{
//   			MarketType: jsii.String("marketType"),
//   			SpotOptions: &SpotOptionsProperty{
//   				BlockDurationMinutes: jsii.Number(123),
//   				InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   				MaxPrice: jsii.String("maxPrice"),
//   				SpotInstanceType: jsii.String("spotInstanceType"),
//   				ValidUntil: jsii.String("validUntil"),
//   			},
//   		},
//   		InstanceRequirements: &InstanceRequirementsProperty{
//   			AcceleratorCount: &AcceleratorCountProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			AcceleratorManufacturers: []*string{
//   				jsii.String("acceleratorManufacturers"),
//   			},
//   			AcceleratorNames: []*string{
//   				jsii.String("acceleratorNames"),
//   			},
//   			AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			AcceleratorTypes: []*string{
//   				jsii.String("acceleratorTypes"),
//   			},
//   			AllowedInstanceTypes: []*string{
//   				jsii.String("allowedInstanceTypes"),
//   			},
//   			BareMetal: jsii.String("bareMetal"),
//   			BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			BaselinePerformanceFactors: &BaselinePerformanceFactorsProperty{
//   				Cpu: &CpuProperty{
//   					References: []interface{}{
//   						&ReferenceProperty{
//   							InstanceFamily: jsii.String("instanceFamily"),
//   						},
//   					},
//   				},
//   			},
//   			BurstablePerformance: jsii.String("burstablePerformance"),
//   			CpuManufacturers: []*string{
//   				jsii.String("cpuManufacturers"),
//   			},
//   			ExcludedInstanceTypes: []*string{
//   				jsii.String("excludedInstanceTypes"),
//   			},
//   			InstanceGenerations: []*string{
//   				jsii.String("instanceGenerations"),
//   			},
//   			LocalStorage: jsii.String("localStorage"),
//   			LocalStorageTypes: []*string{
//   				jsii.String("localStorageTypes"),
//   			},
//   			MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   			MemoryGiBPerVCpu: &MemoryGiBPerVCpuProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			MemoryMiB: &MemoryMiBProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			NetworkBandwidthGbps: &NetworkBandwidthGbpsProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			NetworkInterfaceCount: &NetworkInterfaceCountProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   			RequireHibernateSupport: jsii.Boolean(false),
//   			SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   			TotalLocalStorageGb: &TotalLocalStorageGBProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			VCpuCount: &VCpuCountProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   		},
//   		InstanceType: jsii.String("instanceType"),
//   		KernelId: jsii.String("kernelId"),
//   		KeyName: jsii.String("keyName"),
//   		LicenseSpecifications: []interface{}{
//   			&LicenseSpecificationProperty{
//   				LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   			},
//   		},
//   		MaintenanceOptions: &MaintenanceOptionsProperty{
//   			AutoRecovery: jsii.String("autoRecovery"),
//   		},
//   		MetadataOptions: &MetadataOptionsProperty{
//   			HttpEndpoint: jsii.String("httpEndpoint"),
//   			HttpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   			HttpPutResponseHopLimit: jsii.Number(123),
//   			HttpTokens: jsii.String("httpTokens"),
//   			InstanceMetadataTags: jsii.String("instanceMetadataTags"),
//   		},
//   		Monitoring: &MonitoringProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		NetworkInterfaces: []interface{}{
//   			&NetworkInterfaceProperty{
//   				AssociateCarrierIpAddress: jsii.Boolean(false),
//   				AssociatePublicIpAddress: jsii.Boolean(false),
//   				ConnectionTrackingSpecification: &ConnectionTrackingSpecificationProperty{
//   					TcpEstablishedTimeout: jsii.Number(123),
//   					UdpStreamTimeout: jsii.Number(123),
//   					UdpTimeout: jsii.Number(123),
//   				},
//   				DeleteOnTermination: jsii.Boolean(false),
//   				Description: jsii.String("description"),
//   				DeviceIndex: jsii.Number(123),
//   				EnaQueueCount: jsii.Number(123),
//   				EnaSrdSpecification: &EnaSrdSpecificationProperty{
//   					EnaSrdEnabled: jsii.Boolean(false),
//   					EnaSrdUdpSpecification: &EnaSrdUdpSpecificationProperty{
//   						EnaSrdUdpEnabled: jsii.Boolean(false),
//   					},
//   				},
//   				Groups: []*string{
//   					jsii.String("groups"),
//   				},
//   				InterfaceType: jsii.String("interfaceType"),
//   				Ipv4PrefixCount: jsii.Number(123),
//   				Ipv4Prefixes: []interface{}{
//   					&Ipv4PrefixSpecificationProperty{
//   						Ipv4Prefix: jsii.String("ipv4Prefix"),
//   					},
//   				},
//   				Ipv6AddressCount: jsii.Number(123),
//   				Ipv6Addresses: []interface{}{
//   					&Ipv6AddProperty{
//   						Ipv6Address: jsii.String("ipv6Address"),
//   					},
//   				},
//   				Ipv6PrefixCount: jsii.Number(123),
//   				Ipv6Prefixes: []interface{}{
//   					&Ipv6PrefixSpecificationProperty{
//   						Ipv6Prefix: jsii.String("ipv6Prefix"),
//   					},
//   				},
//   				NetworkCardIndex: jsii.Number(123),
//   				NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   				PrimaryIpv6: jsii.Boolean(false),
//   				PrivateIpAddress: jsii.String("privateIpAddress"),
//   				PrivateIpAddresses: []interface{}{
//   					&PrivateIpAddProperty{
//   						Primary: jsii.Boolean(false),
//   						PrivateIpAddress: jsii.String("privateIpAddress"),
//   					},
//   				},
//   				SecondaryPrivateIpAddressCount: jsii.Number(123),
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   		NetworkPerformanceOptions: &NetworkPerformanceOptionsProperty{
//   			BandwidthWeighting: jsii.String("bandwidthWeighting"),
//   		},
//   		Placement: &PlacementProperty{
//   			Affinity: jsii.String("affinity"),
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			GroupId: jsii.String("groupId"),
//   			GroupName: jsii.String("groupName"),
//   			HostId: jsii.String("hostId"),
//   			HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   			PartitionNumber: jsii.Number(123),
//   			SpreadDomain: jsii.String("spreadDomain"),
//   			Tenancy: jsii.String("tenancy"),
//   		},
//   		PrivateDnsNameOptions: &PrivateDnsNameOptionsProperty{
//   			EnableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   			EnableResourceNameDnsARecord: jsii.Boolean(false),
//   			HostnameType: jsii.String("hostnameType"),
//   		},
//   		RamDiskId: jsii.String("ramDiskId"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		TagSpecifications: []interface{}{
//   			&TagSpecificationProperty{
//   				ResourceType: jsii.String("resourceType"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		UserData: jsii.String("userData"),
//   	},
//   	LaunchTemplateName: jsii.String("launchTemplateName"),
//   	TagSpecifications: []interface{}{
//   		&LaunchTemplateTagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	VersionDescription: jsii.String("versionDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html
//
type CfnLaunchTemplateMixinProps struct {
	// The information for the launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-launchtemplatedata
	//
	LaunchTemplateData interface{} `field:"optional" json:"launchTemplateData" yaml:"launchTemplateData"`
	// A name for the launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-launchtemplatename
	//
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The tags to apply to the launch template on creation.
	//
	// To tag the launch template, the resource type must be `launch-template` .
	//
	// To specify the tags for resources that are created during instance launch, use [TagSpecifications](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-tagspecifications) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-tagspecifications
	//
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// A description for the first version of the launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-versiondescription
	//
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}


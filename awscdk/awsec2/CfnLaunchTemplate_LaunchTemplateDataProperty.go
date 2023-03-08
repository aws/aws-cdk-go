package awsec2


// The information to include in the launch template.
//
// > You must specify at least one parameter for the launch template data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateDataProperty := &LaunchTemplateDataProperty{
//   	BlockDeviceMappings: []interface{}{
//   		&BlockDeviceMappingProperty{
//   			DeviceName: jsii.String("deviceName"),
//   			Ebs: &EbsProperty{
//   				DeleteOnTermination: jsii.Boolean(false),
//   				Encrypted: jsii.Boolean(false),
//   				Iops: jsii.Number(123),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				SnapshotId: jsii.String("snapshotId"),
//   				Throughput: jsii.Number(123),
//   				VolumeSize: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   			NoDevice: jsii.String("noDevice"),
//   			VirtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	CapacityReservationSpecification: &CapacityReservationSpecificationProperty{
//   		CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   		CapacityReservationTarget: &CapacityReservationTargetProperty{
//   			CapacityReservationId: jsii.String("capacityReservationId"),
//   			CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   		},
//   	},
//   	CpuOptions: &CpuOptionsProperty{
//   		CoreCount: jsii.Number(123),
//   		ThreadsPerCore: jsii.Number(123),
//   	},
//   	CreditSpecification: &CreditSpecificationProperty{
//   		CpuCredits: jsii.String("cpuCredits"),
//   	},
//   	DisableApiStop: jsii.Boolean(false),
//   	DisableApiTermination: jsii.Boolean(false),
//   	EbsOptimized: jsii.Boolean(false),
//   	ElasticGpuSpecifications: []interface{}{
//   		&ElasticGpuSpecificationProperty{
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ElasticInferenceAccelerators: []interface{}{
//   		&LaunchTemplateElasticInferenceAcceleratorProperty{
//   			Count: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	EnclaveOptions: &EnclaveOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	HibernationOptions: &HibernationOptionsProperty{
//   		Configured: jsii.Boolean(false),
//   	},
//   	IamInstanceProfile: &IamInstanceProfileProperty{
//   		Arn: jsii.String("arn"),
//   		Name: jsii.String("name"),
//   	},
//   	ImageId: jsii.String("imageId"),
//   	InstanceInitiatedShutdownBehavior: jsii.String("instanceInitiatedShutdownBehavior"),
//   	InstanceMarketOptions: &InstanceMarketOptionsProperty{
//   		MarketType: jsii.String("marketType"),
//   		SpotOptions: &SpotOptionsProperty{
//   			BlockDurationMinutes: jsii.Number(123),
//   			InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   			MaxPrice: jsii.String("maxPrice"),
//   			SpotInstanceType: jsii.String("spotInstanceType"),
//   			ValidUntil: jsii.String("validUntil"),
//   		},
//   	},
//   	InstanceRequirements: &InstanceRequirementsProperty{
//   		AcceleratorCount: &AcceleratorCountProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		AcceleratorManufacturers: []*string{
//   			jsii.String("acceleratorManufacturers"),
//   		},
//   		AcceleratorNames: []*string{
//   			jsii.String("acceleratorNames"),
//   		},
//   		AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		AcceleratorTypes: []*string{
//   			jsii.String("acceleratorTypes"),
//   		},
//   		AllowedInstanceTypes: []*string{
//   			jsii.String("allowedInstanceTypes"),
//   		},
//   		BareMetal: jsii.String("bareMetal"),
//   		BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		BurstablePerformance: jsii.String("burstablePerformance"),
//   		CpuManufacturers: []*string{
//   			jsii.String("cpuManufacturers"),
//   		},
//   		ExcludedInstanceTypes: []*string{
//   			jsii.String("excludedInstanceTypes"),
//   		},
//   		InstanceGenerations: []*string{
//   			jsii.String("instanceGenerations"),
//   		},
//   		LocalStorage: jsii.String("localStorage"),
//   		LocalStorageTypes: []*string{
//   			jsii.String("localStorageTypes"),
//   		},
//   		MemoryGiBPerVCpu: &MemoryGiBPerVCpuProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		MemoryMiB: &MemoryMiBProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		NetworkBandwidthGbps: &NetworkBandwidthGbpsProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		NetworkInterfaceCount: &NetworkInterfaceCountProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   		RequireHibernateSupport: jsii.Boolean(false),
//   		SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   		TotalLocalStorageGb: &TotalLocalStorageGBProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		VCpuCount: &VCpuCountProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	KernelId: jsii.String("kernelId"),
//   	KeyName: jsii.String("keyName"),
//   	LicenseSpecifications: []interface{}{
//   		&LicenseSpecificationProperty{
//   			LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   		},
//   	},
//   	MaintenanceOptions: &MaintenanceOptionsProperty{
//   		AutoRecovery: jsii.String("autoRecovery"),
//   	},
//   	MetadataOptions: &MetadataOptionsProperty{
//   		HttpEndpoint: jsii.String("httpEndpoint"),
//   		HttpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   		HttpPutResponseHopLimit: jsii.Number(123),
//   		HttpTokens: jsii.String("httpTokens"),
//   		InstanceMetadataTags: jsii.String("instanceMetadataTags"),
//   	},
//   	Monitoring: &MonitoringProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	NetworkInterfaces: []interface{}{
//   		&NetworkInterfaceProperty{
//   			AssociateCarrierIpAddress: jsii.Boolean(false),
//   			AssociatePublicIpAddress: jsii.Boolean(false),
//   			DeleteOnTermination: jsii.Boolean(false),
//   			Description: jsii.String("description"),
//   			DeviceIndex: jsii.Number(123),
//   			Groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			InterfaceType: jsii.String("interfaceType"),
//   			Ipv4PrefixCount: jsii.Number(123),
//   			Ipv4Prefixes: []interface{}{
//   				&Ipv4PrefixSpecificationProperty{
//   					Ipv4Prefix: jsii.String("ipv4Prefix"),
//   				},
//   			},
//   			Ipv6AddressCount: jsii.Number(123),
//   			Ipv6Addresses: []interface{}{
//   				&Ipv6AddProperty{
//   					Ipv6Address: jsii.String("ipv6Address"),
//   				},
//   			},
//   			Ipv6PrefixCount: jsii.Number(123),
//   			Ipv6Prefixes: []interface{}{
//   				&Ipv6PrefixSpecificationProperty{
//   					Ipv6Prefix: jsii.String("ipv6Prefix"),
//   				},
//   			},
//   			NetworkCardIndex: jsii.Number(123),
//   			NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   			PrivateIpAddress: jsii.String("privateIpAddress"),
//   			PrivateIpAddresses: []interface{}{
//   				&PrivateIpAddProperty{
//   					Primary: jsii.Boolean(false),
//   					PrivateIpAddress: jsii.String("privateIpAddress"),
//   				},
//   			},
//   			SecondaryPrivateIpAddressCount: jsii.Number(123),
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	Placement: &PlacementProperty{
//   		Affinity: jsii.String("affinity"),
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		GroupId: jsii.String("groupId"),
//   		GroupName: jsii.String("groupName"),
//   		HostId: jsii.String("hostId"),
//   		HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   		PartitionNumber: jsii.Number(123),
//   		SpreadDomain: jsii.String("spreadDomain"),
//   		Tenancy: jsii.String("tenancy"),
//   	},
//   	PrivateDnsNameOptions: &PrivateDnsNameOptionsProperty{
//   		EnableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   		EnableResourceNameDnsARecord: jsii.Boolean(false),
//   		HostnameType: jsii.String("hostnameType"),
//   	},
//   	RamDiskId: jsii.String("ramDiskId"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
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
//   	UserData: jsii.String("userData"),
//   }
//
type CfnLaunchTemplate_LaunchTemplateDataProperty struct {
	// The block device mapping.
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The Capacity Reservation targeting option.
	//
	// If you do not specify this parameter, the instance's Capacity Reservation preference defaults to `open` , which enables it to run in any open Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).
	CapacityReservationSpecification interface{} `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// The CPU options for the instance.
	//
	// For more information, see [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the *Amazon Elastic Compute Cloud User Guide* .
	CpuOptions interface{} `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// The credit option for CPU usage of the instance.
	//
	// Valid only for T instances.
	CreditSpecification interface{} `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Indicates whether to enable the instance for stop protection.
	//
	// For more information, see [Stop protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html#Using_StopProtection) in the *Amazon Elastic Compute Cloud User Guide* .
	DisableApiStop interface{} `field:"optional" json:"disableApiStop" yaml:"disableApiStop"`
	// If you set this parameter to `true` , you can't terminate the instance using the Amazon EC2 console, CLI, or API;
	//
	// otherwise, you can. To change this attribute after launch, use [ModifyInstanceAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html) . Alternatively, if you set `InstanceInitiatedShutdownBehavior` to `terminate` , you can terminate the instance by running the shutdown command from the instance.
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// An elastic GPU to associate with the instance.
	ElasticGpuSpecifications interface{} `field:"optional" json:"elasticGpuSpecifications" yaml:"elasticGpuSpecifications"`
	// The elastic inference accelerator for the instance.
	ElasticInferenceAccelerators interface{} `field:"optional" json:"elasticInferenceAccelerators" yaml:"elasticInferenceAccelerators"`
	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	//
	// For more information, see [What is AWS Nitro Enclaves?](https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html) in the *AWS Nitro Enclaves User Guide* .
	//
	// You can't enable AWS Nitro Enclaves and hibernation on the same instance.
	EnclaveOptions interface{} `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// Indicates whether an instance is enabled for hibernation.
	//
	// This parameter is valid only if the instance meets the [hibernation prerequisites](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/hibernating-prerequisites.html) . For more information, see [Hibernate your instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the *Amazon Elastic Compute Cloud User Guide* .
	HibernationOptions interface{} `field:"optional" json:"hibernationOptions" yaml:"hibernationOptions"`
	// The name or Amazon Resource Name (ARN) of an IAM instance profile.
	IamInstanceProfile interface{} `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// The ID of the AMI.
	//
	// Alternatively, you can specify a Systems Manager parameter, which will resolve to an AMI ID on launch.
	//
	// Valid formats:
	//
	// - `ami-17characters00000`
	// - `resolve:ssm:parameter-name`
	// - `resolve:ssm:parameter-name:version-number`
	// - `resolve:ssm:parameter-name:label`
	//
	// For more information, see [Use a Systems Manager parameter to find an AMI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html#using-systems-manager-parameter-to-find-AMI) in the *Amazon Elastic Compute Cloud User Guide* .
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	//
	// Default: `stop`.
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// The market (purchasing) option for the instances.
	InstanceMarketOptions interface{} `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// The attributes for the instance types.
	//
	// When you specify instance attributes, Amazon EC2 will identify instance types with these attributes.
	//
	// If you specify `InstanceRequirements` , you can't specify `InstanceType` .
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon Elastic Compute Cloud User Guide* .
	//
	// If you specify `InstanceType` , you can't specify `InstanceRequirements` .
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ID of the kernel.
	//
	// We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User Provided Kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon EC2 User Guide* .
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// The name of the key pair. You can create a key pair using [CreateKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html) or [ImportKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html) .
	//
	// > If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The license configurations.
	LicenseSpecifications interface{} `field:"optional" json:"licenseSpecifications" yaml:"licenseSpecifications"`
	// The maintenance options of your instance.
	MaintenanceOptions interface{} `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// The metadata options for the instance.
	//
	// For more information, see [Instance metadata and user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) in the *Amazon Elastic Compute Cloud User Guide* .
	MetadataOptions interface{} `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// The monitoring for the instance.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// One or more network interfaces.
	//
	// If you specify a network interface, you must specify any security groups and subnets as part of the network interface.
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The placement for the instance.
	Placement interface{} `field:"optional" json:"placement" yaml:"placement"`
	// The options for the instance hostname.
	//
	// The default values are inherited from the subnet.
	PrivateDnsNameOptions interface{} `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// The ID of the RAM disk.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User provided kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon Elastic Compute Cloud User Guide* .
	RamDiskId *string `field:"optional" json:"ramDiskId" yaml:"ramDiskId"`
	// One or more security group IDs.
	//
	// You can create a security group using [CreateSecurityGroup](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html) . You cannot specify both a security group ID and security name in the same request.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// One or more security group names.
	//
	// For a nondefault VPC, you must use security group IDs instead. You cannot specify both a security group ID and security name in the same request.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The tags to apply to the resources that are created during instance launch.
	//
	// You can specify tags for the following resources only:
	//
	// - Instances
	// - Volumes
	// - Elastic graphics
	// - Spot Instance requests
	// - Network interfaces
	//
	// To tag a resource after it has been created, see [CreateTags](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html) .
	//
	// > To tag the launch template itself, you must use the [TagSpecification](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateLaunchTemplate.html) parameter.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The user data to make available to the instance.
	//
	// You must provide base64-encoded text. User data is limited to 16 KB. For more information, see [Run commands on your Linux instance at launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) (Linux) or [Work with instance user data](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/instancedata-add-user-data.html) (Windows) in the *Amazon Elastic Compute Cloud User Guide* .
	//
	// If you are creating the launch template for use with AWS Batch , the user data must be provided in the [MIME multi-part archive format](https://docs.aws.amazon.com/https://cloudinit.readthedocs.io/en/latest/topics/format.html#mime-multi-part-archive) . For more information, see [Amazon EC2 user data in launch templates](https://docs.aws.amazon.com/batch/latest/userguide/launch-templates.html) in the *AWS Batch User Guide* .
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}


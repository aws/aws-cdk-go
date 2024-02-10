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
//   		AmdSevSnp: jsii.String("amdSevSnp"),
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
//   		MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
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
//   		RebootMigration: jsii.String("rebootMigration"),
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
//   			ConnectionTrackingSpecification: &ConnectionTrackingSpecificationProperty{
//   				TcpEstablishedTimeout: jsii.Number(123),
//   				UdpStreamTimeout: jsii.Number(123),
//   				UdpTimeout: jsii.Number(123),
//   			},
//   			DeleteOnTermination: jsii.Boolean(false),
//   			Description: jsii.String("description"),
//   			DeviceIndex: jsii.Number(123),
//   			EnaSrdSpecification: &EnaSrdSpecificationProperty{
//   				EnaSrdEnabled: jsii.Boolean(false),
//   				EnaSrdUdpSpecification: &EnaSrdUdpSpecificationProperty{
//   					EnaSrdUdpEnabled: jsii.Boolean(false),
//   				},
//   			},
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
//   			PrimaryIpv6: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html
//
type CfnLaunchTemplate_LaunchTemplateDataProperty struct {
	// The block device mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-blockdevicemappings
	//
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The Capacity Reservation targeting option.
	//
	// If you do not specify this parameter, the instance's Capacity Reservation preference defaults to `open` , which enables it to run in any open Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-capacityreservationspecification
	//
	CapacityReservationSpecification interface{} `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// The CPU options for the instance.
	//
	// For more information, see [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the *Amazon Elastic Compute Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-cpuoptions
	//
	CpuOptions interface{} `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// The credit option for CPU usage of the instance.
	//
	// Valid only for T instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-creditspecification
	//
	CreditSpecification interface{} `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Indicates whether to enable the instance for stop protection.
	//
	// For more information, see [Stop protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html#Using_StopProtection) in the *Amazon Elastic Compute Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-disableapistop
	//
	DisableApiStop interface{} `field:"optional" json:"disableApiStop" yaml:"disableApiStop"`
	// If you set this parameter to `true` , you can't terminate the instance using the Amazon EC2 console, CLI, or API;
	//
	// otherwise, you can. To change this attribute after launch, use [ModifyInstanceAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html) . Alternatively, if you set `InstanceInitiatedShutdownBehavior` to `terminate` , you can terminate the instance by running the shutdown command from the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-disableapitermination
	//
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-ebsoptimized
	//
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Deprecated.
	//
	// > Amazon Elastic Graphics reached end of life on January 8, 2024. For workloads that require graphics acceleration, we recommend that you use Amazon EC2 G4ad, G4dn, or G5 instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-elasticgpuspecifications
	//
	ElasticGpuSpecifications interface{} `field:"optional" json:"elasticGpuSpecifications" yaml:"elasticGpuSpecifications"`
	// An elastic inference accelerator to associate with the instance.
	//
	// Elastic inference accelerators are a resource you can attach to your Amazon EC2 instances to accelerate your Deep Learning (DL) inference workloads.
	//
	// You cannot specify accelerators from different generations in the same request.
	//
	// > Starting April 15, 2023, AWS will not onboard new customers to Amazon Elastic Inference (EI), and will help current customers migrate their workloads to options that offer better price and performance. After April 15, 2023, new customers will not be able to launch instances with Amazon EI accelerators in Amazon SageMaker, Amazon ECS, or Amazon EC2. However, customers who have used Amazon EI at least once during the past 30-day period are considered current customers and will be able to continue using the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-elasticinferenceaccelerators
	//
	ElasticInferenceAccelerators interface{} `field:"optional" json:"elasticInferenceAccelerators" yaml:"elasticInferenceAccelerators"`
	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	//
	// For more information, see [What is AWS Nitro Enclaves?](https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html) in the *AWS Nitro Enclaves User Guide* .
	//
	// You can't enable AWS Nitro Enclaves and hibernation on the same instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-enclaveoptions
	//
	EnclaveOptions interface{} `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// Indicates whether an instance is enabled for hibernation.
	//
	// This parameter is valid only if the instance meets the [hibernation prerequisites](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/hibernating-prerequisites.html) . For more information, see [Hibernate your instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the *Amazon Elastic Compute Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-hibernationoptions
	//
	HibernationOptions interface{} `field:"optional" json:"hibernationOptions" yaml:"hibernationOptions"`
	// The name or Amazon Resource Name (ARN) of an IAM instance profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-iaminstanceprofile
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-imageid
	//
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	//
	// Default: `stop`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-instanceinitiatedshutdownbehavior
	//
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// The market (purchasing) option for the instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-instancemarketoptions
	//
	InstanceMarketOptions interface{} `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// The attributes for the instance types.
	//
	// When you specify instance attributes, Amazon EC2 will identify instance types with these attributes.
	//
	// You must specify `VCpuCount` and `MemoryMiB` . All other attributes are optional. Any unspecified optional attribute is set to its default.
	//
	// When you specify multiple attributes, you get instance types that satisfy all of the specified attributes. If you specify multiple values for an attribute, you get instance types that satisfy any of the specified values.
	//
	// To limit the list of instance types from which Amazon EC2 can identify matching instance types, you can use one of the following parameters, but not both in the same request:
	//
	// - `AllowedInstanceTypes` - The instance types to include in the list. All other instance types are ignored, even if they match your specified attributes.
	// - `ExcludedInstanceTypes` - The instance types to exclude from the list, even if they match your specified attributes.
	//
	// > If you specify `InstanceRequirements` , you can't specify `InstanceType` .
	// >
	// > Attribute-based instance type selection is only supported when using Auto Scaling groups, EC2 Fleet, and Spot Fleet to launch instances. If you plan to use the launch template in the [launch instance wizard](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-instance-wizard.html) , or with the [RunInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html) API or [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) AWS CloudFormation resource, you can't specify `InstanceRequirements` .
	//
	// For more information, see [Attribute-based instance type selection for EC2 Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-attribute-based-instance-type-selection.html) , [Attribute-based instance type selection for Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-attribute-based-instance-type-selection.html) , and [Spot placement score](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-placement-score.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-instancerequirements
	//
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon Elastic Compute Cloud User Guide* .
	//
	// If you specify `InstanceType` , you can't specify `InstanceRequirements` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ID of the kernel.
	//
	// We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User Provided Kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-kernelid
	//
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// The name of the key pair. You can create a key pair using [CreateKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html) or [ImportKeyPair](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html) .
	//
	// > If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-keyname
	//
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The license configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-licensespecifications
	//
	LicenseSpecifications interface{} `field:"optional" json:"licenseSpecifications" yaml:"licenseSpecifications"`
	// The maintenance options of your instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-maintenanceoptions
	//
	MaintenanceOptions interface{} `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// The metadata options for the instance.
	//
	// For more information, see [Instance metadata and user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) in the *Amazon Elastic Compute Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions
	//
	MetadataOptions interface{} `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// The monitoring for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-monitoring
	//
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// One or more network interfaces.
	//
	// If you specify a network interface, you must specify any security groups and subnets as part of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-networkinterfaces
	//
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The placement for the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-placement
	//
	Placement interface{} `field:"optional" json:"placement" yaml:"placement"`
	// The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries should be handled.
	//
	// For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *Amazon Elastic Compute Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-privatednsnameoptions
	//
	PrivateDnsNameOptions interface{} `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// The ID of the RAM disk.
	//
	// > We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see [User provided kernels](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the *Amazon Elastic Compute Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-ramdiskid
	//
	RamDiskId *string `field:"optional" json:"ramDiskId" yaml:"ramDiskId"`
	// The IDs of the security groups.
	//
	// You can specify the IDs of existing security groups and references to resources created by the stack template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// One or more security group names.
	//
	// For a nondefault VPC, you must use security group IDs instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The tags to apply to the resources that are created during instance launch.
	//
	// To tag a resource after it has been created, see [CreateTags](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html) .
	//
	// To tag the launch template itself, use [TagSpecifications](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html#cfn-ec2-launchtemplate-tagspecifications) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-tagspecifications
	//
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The user data to make available to the instance.
	//
	// You must provide base64-encoded text. User data is limited to 16 KB. For more information, see [Run commands on your Linux instance at launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) (Linux) or [Work with instance user data](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/instancedata-add-user-data.html) (Windows) in the *Amazon Elastic Compute Cloud User Guide* .
	//
	// If you are creating the launch template for use with AWS Batch , the user data must be provided in the [MIME multi-part archive format](https://docs.aws.amazon.com/https://cloudinit.readthedocs.io/en/latest/topics/format.html#mime-multi-part-archive) . For more information, see [Amazon EC2 user data in launch templates](https://docs.aws.amazon.com/batch/latest/userguide/launch-templates.html) in the *AWS Batch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html#cfn-ec2-launchtemplate-launchtemplatedata-userdata
	//
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}


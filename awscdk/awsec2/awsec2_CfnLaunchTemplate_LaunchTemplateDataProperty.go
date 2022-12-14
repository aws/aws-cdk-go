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
//   launchTemplateDataProperty := &launchTemplateDataProperty{
//   	blockDeviceMappings: []interface{}{
//   		&blockDeviceMappingProperty{
//   			deviceName: jsii.String("deviceName"),
//   			ebs: &ebsProperty{
//   				deleteOnTermination: jsii.Boolean(false),
//   				encrypted: jsii.Boolean(false),
//   				iops: jsii.Number(123),
//   				kmsKeyId: jsii.String("kmsKeyId"),
//   				snapshotId: jsii.String("snapshotId"),
//   				throughput: jsii.Number(123),
//   				volumeSize: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//   			},
//   			noDevice: jsii.String("noDevice"),
//   			virtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	capacityReservationSpecification: &capacityReservationSpecificationProperty{
//   		capacityReservationPreference: jsii.String("capacityReservationPreference"),
//   		capacityReservationTarget: &capacityReservationTargetProperty{
//   			capacityReservationId: jsii.String("capacityReservationId"),
//   			capacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   		},
//   	},
//   	cpuOptions: &cpuOptionsProperty{
//   		coreCount: jsii.Number(123),
//   		threadsPerCore: jsii.Number(123),
//   	},
//   	creditSpecification: &creditSpecificationProperty{
//   		cpuCredits: jsii.String("cpuCredits"),
//   	},
//   	disableApiStop: jsii.Boolean(false),
//   	disableApiTermination: jsii.Boolean(false),
//   	ebsOptimized: jsii.Boolean(false),
//   	elasticGpuSpecifications: []interface{}{
//   		&elasticGpuSpecificationProperty{
//   			type: jsii.String("type"),
//   		},
//   	},
//   	elasticInferenceAccelerators: []interface{}{
//   		&launchTemplateElasticInferenceAcceleratorProperty{
//   			count: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	enclaveOptions: &enclaveOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	hibernationOptions: &hibernationOptionsProperty{
//   		configured: jsii.Boolean(false),
//   	},
//   	iamInstanceProfile: &iamInstanceProfileProperty{
//   		arn: jsii.String("arn"),
//   		name: jsii.String("name"),
//   	},
//   	imageId: jsii.String("imageId"),
//   	instanceInitiatedShutdownBehavior: jsii.String("instanceInitiatedShutdownBehavior"),
//   	instanceMarketOptions: &instanceMarketOptionsProperty{
//   		marketType: jsii.String("marketType"),
//   		spotOptions: &spotOptionsProperty{
//   			blockDurationMinutes: jsii.Number(123),
//   			instanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   			maxPrice: jsii.String("maxPrice"),
//   			spotInstanceType: jsii.String("spotInstanceType"),
//   			validUntil: jsii.String("validUntil"),
//   		},
//   	},
//   	instanceRequirements: &instanceRequirementsProperty{
//   		acceleratorCount: &acceleratorCountProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		acceleratorManufacturers: []*string{
//   			jsii.String("acceleratorManufacturers"),
//   		},
//   		acceleratorNames: []*string{
//   			jsii.String("acceleratorNames"),
//   		},
//   		acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		acceleratorTypes: []*string{
//   			jsii.String("acceleratorTypes"),
//   		},
//   		allowedInstanceTypes: []*string{
//   			jsii.String("allowedInstanceTypes"),
//   		},
//   		bareMetal: jsii.String("bareMetal"),
//   		baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		burstablePerformance: jsii.String("burstablePerformance"),
//   		cpuManufacturers: []*string{
//   			jsii.String("cpuManufacturers"),
//   		},
//   		excludedInstanceTypes: []*string{
//   			jsii.String("excludedInstanceTypes"),
//   		},
//   		instanceGenerations: []*string{
//   			jsii.String("instanceGenerations"),
//   		},
//   		localStorage: jsii.String("localStorage"),
//   		localStorageTypes: []*string{
//   			jsii.String("localStorageTypes"),
//   		},
//   		memoryGiBPerVCpu: &memoryGiBPerVCpuProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		memoryMiB: &memoryMiBProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		networkBandwidthGbps: &networkBandwidthGbpsProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		networkInterfaceCount: &networkInterfaceCountProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   		requireHibernateSupport: jsii.Boolean(false),
//   		spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   		totalLocalStorageGb: &totalLocalStorageGBProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		vCpuCount: &vCpuCountProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   	},
//   	instanceType: jsii.String("instanceType"),
//   	kernelId: jsii.String("kernelId"),
//   	keyName: jsii.String("keyName"),
//   	licenseSpecifications: []interface{}{
//   		&licenseSpecificationProperty{
//   			licenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   		},
//   	},
//   	maintenanceOptions: &maintenanceOptionsProperty{
//   		autoRecovery: jsii.String("autoRecovery"),
//   	},
//   	metadataOptions: &metadataOptionsProperty{
//   		httpEndpoint: jsii.String("httpEndpoint"),
//   		httpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   		httpPutResponseHopLimit: jsii.Number(123),
//   		httpTokens: jsii.String("httpTokens"),
//   		instanceMetadataTags: jsii.String("instanceMetadataTags"),
//   	},
//   	monitoring: &monitoringProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	networkInterfaces: []interface{}{
//   		&networkInterfaceProperty{
//   			associateCarrierIpAddress: jsii.Boolean(false),
//   			associatePublicIpAddress: jsii.Boolean(false),
//   			deleteOnTermination: jsii.Boolean(false),
//   			description: jsii.String("description"),
//   			deviceIndex: jsii.Number(123),
//   			groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			interfaceType: jsii.String("interfaceType"),
//   			ipv4PrefixCount: jsii.Number(123),
//   			ipv4Prefixes: []interface{}{
//   				&ipv4PrefixSpecificationProperty{
//   					ipv4Prefix: jsii.String("ipv4Prefix"),
//   				},
//   			},
//   			ipv6AddressCount: jsii.Number(123),
//   			ipv6Addresses: []interface{}{
//   				&ipv6AddProperty{
//   					ipv6Address: jsii.String("ipv6Address"),
//   				},
//   			},
//   			ipv6PrefixCount: jsii.Number(123),
//   			ipv6Prefixes: []interface{}{
//   				&ipv6PrefixSpecificationProperty{
//   					ipv6Prefix: jsii.String("ipv6Prefix"),
//   				},
//   			},
//   			networkCardIndex: jsii.Number(123),
//   			networkInterfaceId: jsii.String("networkInterfaceId"),
//   			privateIpAddress: jsii.String("privateIpAddress"),
//   			privateIpAddresses: []interface{}{
//   				&privateIpAddProperty{
//   					primary: jsii.Boolean(false),
//   					privateIpAddress: jsii.String("privateIpAddress"),
//   				},
//   			},
//   			secondaryPrivateIpAddressCount: jsii.Number(123),
//   			subnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	placement: &placementProperty{
//   		affinity: jsii.String("affinity"),
//   		availabilityZone: jsii.String("availabilityZone"),
//   		groupId: jsii.String("groupId"),
//   		groupName: jsii.String("groupName"),
//   		hostId: jsii.String("hostId"),
//   		hostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   		partitionNumber: jsii.Number(123),
//   		spreadDomain: jsii.String("spreadDomain"),
//   		tenancy: jsii.String("tenancy"),
//   	},
//   	privateDnsNameOptions: &privateDnsNameOptionsProperty{
//   		enableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   		enableResourceNameDnsARecord: jsii.Boolean(false),
//   		hostnameType: jsii.String("hostnameType"),
//   	},
//   	ramDiskId: jsii.String("ramDiskId"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
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
//   	userData: jsii.String("userData"),
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
	// Valid for T2, T3, or T3a instances only.
	CreditSpecification interface{} `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// `CfnLaunchTemplate.LaunchTemplateDataProperty.DisableApiStop`.
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
	// If you specify `InstanceRequirements` , you can't specify `InstanceTypes` .
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon Elastic Compute Cloud User Guide* .
	//
	// If you specify `InstanceTypes` , you can't specify `InstanceRequirements` .
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
	// `CfnLaunchTemplate.LaunchTemplateDataProperty.MaintenanceOptions`.
	MaintenanceOptions interface{} `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// The metadata options for the instance.
	//
	// For more information, see [Instance Metadata and User Data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) in the *Amazon Elastic Compute Cloud User Guide* .
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
	// [EC2-Classic, default VPC] One or more security group names.
	//
	// For a nondefault VPC, you must use security group IDs instead. You cannot specify both a security group ID and security name in the same request.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The tags to apply to the resources during launch.
	//
	// You can only tag instances and volumes on launch. The specified tags are applied to all instances or volumes that are created during launch.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The user data to make available to the instance.
	//
	// You must provide base64-encoded text. User data is limited to 16 KB. For more information, see [Run commands on your Linux instance at launch](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) (Linux) or [Work with instance user data](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/instancedata-add-user-data.html) (Windows) in the *Amazon Elastic Compute Cloud User Guide* .
	//
	// If you are creating the launch template for use with AWS Batch , the user data must be provided in the [MIME multi-part archive format](https://docs.aws.amazon.com/https://cloudinit.readthedocs.io/en/latest/topics/format.html#mime-multi-part-archive) . For more information, see [Amazon EC2 user data in launch templates](https://docs.aws.amazon.com/batch/latest/userguide/launch-templates.html) in the *AWS Batch User Guide* .
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}


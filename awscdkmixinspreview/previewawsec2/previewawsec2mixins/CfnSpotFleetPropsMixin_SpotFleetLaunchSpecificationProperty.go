package previewawsec2mixins


// Specifies the launch specification for one or more Spot Instances.
//
// If you include On-Demand capacity in your fleet request, you can't use `SpotFleetLaunchSpecification` ; you must use [LaunchTemplateConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-launchtemplateconfig.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spotFleetLaunchSpecificationProperty := &SpotFleetLaunchSpecificationProperty{
//   	BlockDeviceMappings: []interface{}{
//   		&BlockDeviceMappingProperty{
//   			DeviceName: jsii.String("deviceName"),
//   			Ebs: &EbsBlockDeviceProperty{
//   				DeleteOnTermination: jsii.Boolean(false),
//   				Encrypted: jsii.Boolean(false),
//   				Iops: jsii.Number(123),
//   				SnapshotId: jsii.String("snapshotId"),
//   				VolumeSize: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   			NoDevice: jsii.String("noDevice"),
//   			VirtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	EbsOptimized: jsii.Boolean(false),
//   	IamInstanceProfile: &IamInstanceProfileSpecificationProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	ImageId: jsii.String("imageId"),
//   	InstanceRequirements: &InstanceRequirementsRequestProperty{
//   		AcceleratorCount: &AcceleratorCountRequestProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		AcceleratorManufacturers: []*string{
//   			jsii.String("acceleratorManufacturers"),
//   		},
//   		AcceleratorNames: []*string{
//   			jsii.String("acceleratorNames"),
//   		},
//   		AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
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
//   		BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		BaselinePerformanceFactors: &BaselinePerformanceFactorsRequestProperty{
//   			Cpu: &CpuPerformanceFactorRequestProperty{
//   				References: []interface{}{
//   					&PerformanceFactorReferenceRequestProperty{
//   						InstanceFamily: jsii.String("instanceFamily"),
//   					},
//   				},
//   			},
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
//   		MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		MemoryMiB: &MemoryMiBRequestProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   		RequireHibernateSupport: jsii.Boolean(false),
//   		SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   		TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		VCpuCount: &VCpuCountRangeRequestProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	KernelId: jsii.String("kernelId"),
//   	KeyName: jsii.String("keyName"),
//   	Monitoring: &SpotFleetMonitoringProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	NetworkInterfaces: []interface{}{
//   		&InstanceNetworkInterfaceSpecificationProperty{
//   			AssociatePublicIpAddress: jsii.Boolean(false),
//   			DeleteOnTermination: jsii.Boolean(false),
//   			Description: jsii.String("description"),
//   			DeviceIndex: jsii.Number(123),
//   			Groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			Ipv6AddressCount: jsii.Number(123),
//   			Ipv6Addresses: []interface{}{
//   				&InstanceIpv6AddressProperty{
//   					Ipv6Address: jsii.String("ipv6Address"),
//   				},
//   			},
//   			NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   			PrivateIpAddresses: []interface{}{
//   				&PrivateIpAddressSpecificationProperty{
//   					Primary: jsii.Boolean(false),
//   					PrivateIpAddress: jsii.String("privateIpAddress"),
//   				},
//   			},
//   			SecondaryPrivateIpAddressCount: jsii.Number(123),
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	Placement: &SpotPlacementProperty{
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		GroupName: jsii.String("groupName"),
//   		Tenancy: jsii.String("tenancy"),
//   	},
//   	RamdiskId: jsii.String("ramdiskId"),
//   	SecurityGroups: []interface{}{
//   		&GroupIdentifierProperty{
//   			GroupId: jsii.String("groupId"),
//   		},
//   	},
//   	SpotPrice: jsii.String("spotPrice"),
//   	SubnetId: jsii.String("subnetId"),
//   	TagSpecifications: []interface{}{
//   		&SpotFleetTagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	UserData: jsii.String("userData"),
//   	WeightedCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html
//
type CfnSpotFleetPropsMixin_SpotFleetLaunchSpecificationProperty struct {
	// One or more block devices that are mapped to the Spot Instances.
	//
	// You can't specify both a snapshot ID and an encryption value. This is because only blank volumes can be encrypted on creation. If a snapshot is the basis for a volume, it is not blank and its encryption status is used for the volume encryption status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-blockdevicemappings
	//
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// Indicates whether the instances are optimized for EBS I/O.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS Optimized instance.
	//
	// Default: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-ebsoptimized
	//
	// Default: - false.
	//
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// The IAM instance profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-iaminstanceprofile
	//
	IamInstanceProfile interface{} `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// The ID of the AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-imageid
	//
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// The attributes for the instance types.
	//
	// When you specify instance attributes, Amazon EC2 will identify instance types with those attributes.
	//
	// > If you specify `InstanceRequirements` , you can't specify `InstanceType` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-instancerequirements
	//
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ID of the kernel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-kernelid
	//
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// The name of the key pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-keyname
	//
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Enable or disable monitoring for the instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-monitoring
	//
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// The network interfaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-networkinterfaces
	//
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The placement information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-placement
	//
	Placement interface{} `field:"optional" json:"placement" yaml:"placement"`
	// The ID of the RAM disk.
	//
	// Some kernels require additional drivers at launch. Check the kernel requirements for information about whether you need to specify a RAM disk. To find kernel requirements, refer to the AWS Resource Center and search for the kernel ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-ramdiskid
	//
	RamdiskId *string `field:"optional" json:"ramdiskId" yaml:"ramdiskId"`
	// The security groups.
	//
	// If you specify a network interface, you must specify any security groups as part of the network interface instead of using this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-securitygroups
	//
	SecurityGroups interface{} `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The maximum price per unit hour that you are willing to pay for a Spot Instance.
	//
	// We do not recommend using this parameter because it can lead to increased interruptions. If you do not specify this parameter, you will pay the current Spot price.
	//
	// > If you specify a maximum price, your instances will be interrupted more frequently than if you do not specify this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-spotprice
	//
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// The IDs of the subnets in which to launch the instances.
	//
	// To specify multiple subnets, separate them using commas; for example, "subnet-1234abcdeexample1, subnet-0987cdef6example2".
	//
	// If you specify a network interface, you must specify any subnets as part of the network interface instead of using this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The tags to apply during creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-tagspecifications
	//
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The base64-encoded user data that instances use when starting up.
	//
	// User data is limited to 16 KB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-userdata
	//
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// The number of units provided by the specified instance type.
	//
	// These are the same units that you chose to set the target capacity in terms of instances, or a performance characteristic such as vCPUs, memory, or I/O.
	//
	// If the target capacity divided by this value is not a whole number, Amazon EC2 rounds the number of instances to the next whole number. If this value is not specified, the default is 1.
	//
	// > When specifying weights, the price used in the `lowestPrice` and `priceCapacityOptimized` allocation strategies is per *unit* hour (where the instance price is divided by the specified weight). However, if all the specified weights are above the requested `TargetCapacity` , resulting in only 1 instance being launched, the price used is per *instance* hour.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetlaunchspecification.html#cfn-ec2-spotfleet-spotfleetlaunchspecification-weightedcapacity
	//
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


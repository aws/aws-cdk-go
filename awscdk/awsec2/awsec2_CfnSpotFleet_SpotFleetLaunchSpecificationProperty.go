package awsec2


// Specifies the launch specification for one or more Spot Instances.
//
// If you include On-Demand capacity in your fleet request, you can't use `SpotFleetLaunchSpecification` ; you must use [LaunchTemplateConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-launchtemplateconfig.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spotFleetLaunchSpecificationProperty := &spotFleetLaunchSpecificationProperty{
//   	imageId: jsii.String("imageId"),
//
//   	// the properties below are optional
//   	blockDeviceMappings: []interface{}{
//   		&blockDeviceMappingProperty{
//   			deviceName: jsii.String("deviceName"),
//
//   			// the properties below are optional
//   			ebs: &ebsBlockDeviceProperty{
//   				deleteOnTermination: jsii.Boolean(false),
//   				encrypted: jsii.Boolean(false),
//   				iops: jsii.Number(123),
//   				snapshotId: jsii.String("snapshotId"),
//   				volumeSize: jsii.Number(123),
//   				volumeType: jsii.String("volumeType"),
//   			},
//   			noDevice: jsii.String("noDevice"),
//   			virtualName: jsii.String("virtualName"),
//   		},
//   	},
//   	ebsOptimized: jsii.Boolean(false),
//   	iamInstanceProfile: &iamInstanceProfileSpecificationProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	instanceRequirements: &instanceRequirementsRequestProperty{
//   		acceleratorCount: &acceleratorCountRequestProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		acceleratorManufacturers: []*string{
//   			jsii.String("acceleratorManufacturers"),
//   		},
//   		acceleratorNames: []*string{
//   			jsii.String("acceleratorNames"),
//   		},
//   		acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
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
//   		baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
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
//   		memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		memoryMiB: &memoryMiBRequestProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		networkBandwidthGbps: &networkBandwidthGbpsRequestProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   		requireHibernateSupport: jsii.Boolean(false),
//   		spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   		totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   		vCpuCount: &vCpuCountRangeRequestProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   	},
//   	instanceType: jsii.String("instanceType"),
//   	kernelId: jsii.String("kernelId"),
//   	keyName: jsii.String("keyName"),
//   	monitoring: &spotFleetMonitoringProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	networkInterfaces: []interface{}{
//   		&instanceNetworkInterfaceSpecificationProperty{
//   			associatePublicIpAddress: jsii.Boolean(false),
//   			deleteOnTermination: jsii.Boolean(false),
//   			description: jsii.String("description"),
//   			deviceIndex: jsii.Number(123),
//   			groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			ipv6AddressCount: jsii.Number(123),
//   			ipv6Addresses: []interface{}{
//   				&instanceIpv6AddressProperty{
//   					ipv6Address: jsii.String("ipv6Address"),
//   				},
//   			},
//   			networkInterfaceId: jsii.String("networkInterfaceId"),
//   			privateIpAddresses: []interface{}{
//   				&privateIpAddressSpecificationProperty{
//   					privateIpAddress: jsii.String("privateIpAddress"),
//
//   					// the properties below are optional
//   					primary: jsii.Boolean(false),
//   				},
//   			},
//   			secondaryPrivateIpAddressCount: jsii.Number(123),
//   			subnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	placement: &spotPlacementProperty{
//   		availabilityZone: jsii.String("availabilityZone"),
//   		groupName: jsii.String("groupName"),
//   		tenancy: jsii.String("tenancy"),
//   	},
//   	ramdiskId: jsii.String("ramdiskId"),
//   	securityGroups: []interface{}{
//   		&groupIdentifierProperty{
//   			groupId: jsii.String("groupId"),
//   		},
//   	},
//   	spotPrice: jsii.String("spotPrice"),
//   	subnetId: jsii.String("subnetId"),
//   	tagSpecifications: []interface{}{
//   		&spotFleetTagSpecificationProperty{
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
//   	weightedCapacity: jsii.Number(123),
//   }
//
type CfnSpotFleet_SpotFleetLaunchSpecificationProperty struct {
	// The ID of the AMI.
	ImageId *string `field:"required" json:"imageId" yaml:"imageId"`
	// One or more block devices that are mapped to the Spot Instances.
	//
	// You can't specify both a snapshot ID and an encryption value. This is because only blank volumes can be encrypted on creation. If a snapshot is the basis for a volume, it is not blank and its encryption status is used for the volume encryption status.
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// Indicates whether the instances are optimized for EBS I/O.
	//
	// This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS Optimized instance.
	//
	// Default: `false`.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// The IAM instance profile.
	IamInstanceProfile interface{} `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// The attributes for the instance types.
	//
	// When you specify instance attributes, Amazon EC2 will identify instance types with those attributes.
	//
	// > If you specify `InstanceRequirements` , you can't specify `InstanceTypes` .
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ID of the kernel.
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// The name of the key pair.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Enable or disable monitoring for the instances.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// One or more network interfaces.
	//
	// If you specify a network interface, you must specify subnet IDs and security group IDs using the network interface.
	//
	// > `SpotFleetLaunchSpecification` currently does not support Elastic Fabric Adapter (EFA). To specify an EFA, you must use [LaunchTemplateConfig](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchTemplateConfig.html) .
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The placement information.
	Placement interface{} `field:"optional" json:"placement" yaml:"placement"`
	// The ID of the RAM disk.
	//
	// Some kernels require additional drivers at launch. Check the kernel requirements for information about whether you need to specify a RAM disk. To find kernel requirements, refer to the AWS Resource Center and search for the kernel ID.
	RamdiskId *string `field:"optional" json:"ramdiskId" yaml:"ramdiskId"`
	// One or more security groups.
	//
	// When requesting instances in a VPC, you must specify the IDs of the security groups. When requesting instances in EC2-Classic, you can specify the names or the IDs of the security groups.
	SecurityGroups interface{} `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The maximum price per unit hour that you are willing to pay for a Spot Instance.
	//
	// If this value is not specified, the default is the Spot price specified for the fleet. To determine the Spot price per unit hour, divide the Spot price by the value of `WeightedCapacity` .
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// The IDs of the subnets in which to launch the instances.
	//
	// To specify multiple subnets, separate them using commas; for example, "subnet-1234abcdeexample1, subnet-0987cdef6example2".
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The tags to apply during creation.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The Base64-encoded user data that instances use when starting up.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// The number of units provided by the specified instance type.
	//
	// These are the same units that you chose to set the target capacity in terms of instances, or a performance characteristic such as vCPUs, memory, or I/O.
	//
	// If the target capacity divided by this value is not a whole number, Amazon EC2 rounds the number of instances to the next whole number. If this value is not specified, the default is 1.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


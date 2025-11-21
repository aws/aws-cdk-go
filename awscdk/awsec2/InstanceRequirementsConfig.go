package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The attributes for the instance types for a mixed instances policy.
//
// When you specify multiple attributes, you get instance types that satisfy all of the specified attributes. If you specify multiple values for an attribute, you get instance types that satisfy any of the specified values.
//
// To limit the list of instance types from which Amazon EC2 can identify matching instance types, you can use one of the following parameters, but not both in the same request:
// - AllowedInstanceTypes - The instance types to include in the list. All other instance types are ignored, even if they match your specified attributes.
// - ExcludedInstanceTypes - The instance types to exclude from the list, even if they match your specified attributes.
//
// Note: You must specify VCpuCount and MemoryMiB. All other attributes are optional. Any unspecified optional attribute is set to its default.
//
// Example:
//   var vpc Vpc
//   var infrastructureRole Role
//   var instanceProfile InstanceProfile
//
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   // Create a Managed Instances Capacity Provider
//   miCapacityProvider := ecs.NewManagedInstancesCapacityProvider(this, jsii.String("MICapacityProvider"), &ManagedInstancesCapacityProviderProps{
//   	InfrastructureRole: InfrastructureRole,
//   	Ec2InstanceProfile: instanceProfile,
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []ISecurityGroup{
//   		ec2.NewSecurityGroup(this, jsii.String("MISecurityGroup"), &SecurityGroupProps{
//   			Vpc: *Vpc,
//   		}),
//   	},
//   	InstanceRequirements: &InstanceRequirementsConfig{
//   		VCpuCountMin: jsii.Number(1),
//   		MemoryMin: awscdk.Size_Gibibytes(jsii.Number(2)),
//   		CpuManufacturers: []CpuManufacturer{
//   			ec2.CpuManufacturer_INTEL,
//   		},
//   		AcceleratorManufacturers: []AcceleratorManufacturer{
//   			ec2.AcceleratorManufacturer_NVIDIA,
//   		},
//   	},
//   	PropagateTags: ecs.PropagateManagedInstancesTags_CAPACITY_PROVIDER,
//   })
//
//   // Optionally configure security group rules using IConnectable interface
//   miCapacityProvider.Connections.AllowFrom(ec2.Peer_Ipv4(vpc.VpcCidrBlock), ec2.Port_Tcp(jsii.Number(80)))
//
//   // Add the capacity provider to the cluster
//   cluster.AddManagedInstancesCapacityProvider(miCapacityProvider)
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TaskDef"), &TaskDefinitionProps{
//   	MemoryMiB: jsii.String("512"),
//   	Cpu: jsii.String("256"),
//   	NetworkMode: ecs.NetworkMode_AWS_VPC,
//   	Compatibility: ecs.Compatibility_MANAGED_INSTANCES,
//   })
//
//   taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryReservationMiB: jsii.Number(256),
//   })
//
//   ecs.NewFargateService(this, jsii.String("FargateService"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	MinHealthyPercent: jsii.Number(100),
//   	CapacityProviderStrategies: []CapacityProviderStrategy{
//   		&CapacityProviderStrategy{
//   			CapacityProvider: miCapacityProvider.CapacityProviderName,
//   			Weight: jsii.Number(1),
//   		},
//   	},
//   })
//
type InstanceRequirementsConfig struct {
	// The minimum instance memory size for an instance type, in MiB.
	//
	// Required: Yes.
	MemoryMin awscdk.Size `field:"required" json:"memoryMin" yaml:"memoryMin"`
	// The minimum number of vCPUs for an instance type.
	//
	// Required: Yes.
	VCpuCountMin *float64 `field:"required" json:"vCpuCountMin" yaml:"vCpuCountMin"`
	// The maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) for an instance type.
	//
	// To exclude accelerator-enabled instance types, set Max to 0.
	// Default: - No minimum or maximum limits.
	//
	AcceleratorCountMax *float64 `field:"optional" json:"acceleratorCountMax" yaml:"acceleratorCountMax"`
	// The minimum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) for an instance type.
	//
	// To exclude accelerator-enabled instance types, set acceleratorCountMax to 0.
	// Default: - No minimum or maximum limits.
	//
	AcceleratorCountMin *float64 `field:"optional" json:"acceleratorCountMin" yaml:"acceleratorCountMin"`
	// Indicates whether instance types must have accelerators by specific manufacturers.
	//
	// - For instance types with NVIDIA devices, specify nvidia.
	// - For instance types with AMD devices, specify amd.
	// - For instance types with AWS devices, specify amazon-web-services.
	// - For instance types with Xilinx devices, specify xilinx.
	// Default: - Any manufacturer.
	//
	AcceleratorManufacturers *[]AcceleratorManufacturer `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Lists the accelerators that must be on an instance type.
	//
	// - For instance types with NVIDIA A100 GPUs, specify a100.
	// - For instance types with NVIDIA V100 GPUs, specify v100.
	// - For instance types with NVIDIA K80 GPUs, specify k80.
	// - For instance types with NVIDIA T4 GPUs, specify t4.
	// - For instance types with NVIDIA M60 GPUs, specify m60.
	// - For instance types with AMD Radeon Pro V520 GPUs, specify radeon-pro-v520.
	// - For instance types with Xilinx VU9P FPGAs, specify vu9p.
	// Default: - Any accelerator.
	//
	AcceleratorNames *[]AcceleratorName `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// The maximum total memory size for the accelerators on an instance type, in MiB.
	// Default: - No minimum or maximum limits.
	//
	AcceleratorTotalMemoryMax awscdk.Size `field:"optional" json:"acceleratorTotalMemoryMax" yaml:"acceleratorTotalMemoryMax"`
	// The minimum total memory size for the accelerators on an instance type, in MiB.
	// Default: - No minimum or maximum limits.
	//
	AcceleratorTotalMemoryMin awscdk.Size `field:"optional" json:"acceleratorTotalMemoryMin" yaml:"acceleratorTotalMemoryMin"`
	// Lists the accelerator types that must be on an instance type.
	//
	// - For instance types with GPU accelerators, specify gpu.
	// - For instance types with FPGA accelerators, specify fpga.
	// - For instance types with inference accelerators, specify inference.
	// Default: - Any accelerator type.
	//
	AcceleratorTypes *[]AcceleratorType `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// The instance types to apply your specified attributes against.
	//
	// All other instance types are ignored, even if they match your specified attributes.
	//
	// You can use strings with one or more wild cards, represented by an asterisk (*), to allow an instance type, size, or generation. The following are examples: m5.8xlarge, c5*.*, m5a.*, r*, *3*.
	//
	// For example, if you specify c5*, Amazon EC2 Auto Scaling will allow the entire C5 instance family, which includes all C5a and C5n instance types. If you specify m5a.*, Amazon EC2 Auto Scaling will allow all the M5a instance types, but not the M5n instance types.
	//
	// Note: If you specify AllowedInstanceTypes, you can't specify ExcludedInstanceTypes.
	// Default: - All instance types.
	//
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Indicates whether bare metal instance types are included, excluded, or required.
	// Default: - excluded.
	//
	BareMetal BareMetal `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// The maximum baseline bandwidth performance for an instance type, in Mbps.
	//
	// For more information, see Amazon EBS–optimized instances in the Amazon EC2 User Guide.
	// Default: - No minimum or maximum limits.
	//
	BaselineEbsBandwidthMbpsMax *float64 `field:"optional" json:"baselineEbsBandwidthMbpsMax" yaml:"baselineEbsBandwidthMbpsMax"`
	// The minimum baseline bandwidth performance for an instance type, in Mbps.
	//
	// For more information, see Amazon EBS–optimized instances in the Amazon EC2 User Guide.
	// Default: - No minimum or maximum limits.
	//
	BaselineEbsBandwidthMbpsMin *float64 `field:"optional" json:"baselineEbsBandwidthMbpsMin" yaml:"baselineEbsBandwidthMbpsMin"`
	// Indicates whether burstable performance instance types are included, excluded, or required.
	//
	// For more information, see Burstable performance instances in the Amazon EC2 User Guide.
	// Default: - excluded.
	//
	BurstablePerformance BurstablePerformance `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Lists which specific CPU manufacturers to include.
	//
	// - For instance types with Intel CPUs, specify intel.
	// - For instance types with AMD CPUs, specify amd.
	// - For instance types with AWS CPUs, specify amazon-web-services.
	// - For instance types with Apple CPUs, specify apple.
	//
	// Note: Don't confuse the CPU hardware manufacturer with the CPU hardware architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template.
	// Default: - Any manufacturer.
	//
	CpuManufacturers *[]CpuManufacturer `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// The instance types to exclude.
	//
	// You can use strings with one or more wild cards, represented by an asterisk (*), to exclude an instance family, type, size, or generation. The following are examples: m5.8xlarge, c5*.*, m5a.*, r*, *3*.
	//
	// For example, if you specify c5*, you are excluding the entire C5 instance family, which includes all C5a and C5n instance types. If you specify m5a.*, Amazon EC2 Auto Scaling will exclude all the M5a instance types, but not the M5n instance types.
	//
	// Note: If you specify ExcludedInstanceTypes, you can't specify AllowedInstanceTypes.
	// Default: - No excluded instance types.
	//
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Indicates whether current or previous generation instance types are included.
	//
	// - For current generation instance types, specify current. The current generation includes EC2 instance types currently recommended for use. This typically includes the latest two to three generations in each instance family. For more information, see Instance types in the Amazon EC2 User Guide.
	// - For previous generation instance types, specify previous.
	// Default: - Any current or previous generation.
	//
	InstanceGenerations *[]InstanceGeneration `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Indicates whether instance types with instance store volumes are included, excluded, or required.
	//
	// For more information, see Amazon EC2 instance store in the Amazon EC2 User Guide.
	// Default: - included.
	//
	LocalStorage LocalStorage `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Indicates the type of local storage that is required.
	//
	// - For instance types with hard disk drive (HDD) storage, specify hdd.
	// - For instance types with solid state drive (SSD) storage, specify ssd.
	// Default: - Any local storage type.
	//
	LocalStorageTypes *[]LocalStorageType `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// [Price protection] The price protection threshold for Spot Instances, as a percentage of an identified On-Demand price.
	//
	// The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from either the lowest priced current generation instance types or, failing that, the lowest priced previous generation instance types that match your attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price exceeds your specified threshold.
	//
	// The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage.
	//
	// If you set DesiredCapacityType to vcpu or memory-mib, the price protection threshold is based on the per-vCPU or per-memory price instead of the per instance price.
	//
	// Note: Only one of SpotMaxPricePercentageOverLowestPrice or MaxSpotPriceAsPercentageOfOptimalOnDemandPrice can be specified. If you don't specify either, Amazon EC2 Auto Scaling will automatically apply optimal price protection to consistently select from a wide range of instance types. To indicate no price protection threshold for Spot Instances, meaning you want to consider all instance types that match your attributes, include one of these parameters and specify a high value, such as 999999.
	// Default: - Automatic optimal price protection.
	//
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// The maximum instance memory size for an instance type, in MiB.
	// Default: - No maximum limit.
	//
	MemoryMax awscdk.Size `field:"optional" json:"memoryMax" yaml:"memoryMax"`
	// The maximum amount of memory per vCPU for an instance type, in GiB.
	// Default: - No minimum or maximum limits.
	//
	MemoryPerVCpuMax awscdk.Size `field:"optional" json:"memoryPerVCpuMax" yaml:"memoryPerVCpuMax"`
	// The minimum amount of memory per vCPU for an instance type, in GiB.
	// Default: - No minimum or maximum limits.
	//
	MemoryPerVCpuMin awscdk.Size `field:"optional" json:"memoryPerVCpuMin" yaml:"memoryPerVCpuMin"`
	// The maximum amount of network bandwidth, in gigabits per second (Gbps).
	// Default: - No minimum or maximum limits.
	//
	NetworkBandwidthGbpsMax *float64 `field:"optional" json:"networkBandwidthGbpsMax" yaml:"networkBandwidthGbpsMax"`
	// The minimum amount of network bandwidth, in gigabits per second (Gbps).
	// Default: - No minimum or maximum limits.
	//
	NetworkBandwidthGbpsMin *float64 `field:"optional" json:"networkBandwidthGbpsMin" yaml:"networkBandwidthGbpsMin"`
	// The maximum number of network interfaces for an instance type.
	// Default: - No minimum or maximum limits.
	//
	NetworkInterfaceCountMax *float64 `field:"optional" json:"networkInterfaceCountMax" yaml:"networkInterfaceCountMax"`
	// The minimum number of network interfaces for an instance type.
	// Default: - No minimum or maximum limits.
	//
	NetworkInterfaceCountMin *float64 `field:"optional" json:"networkInterfaceCountMin" yaml:"networkInterfaceCountMin"`
	// [Price protection] The price protection threshold for On-Demand Instances, as a percentage higher than an identified On-Demand price.
	//
	// The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from either the lowest priced current generation instance types or, failing that, the lowest priced previous generation instance types that match your attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price exceeds your specified threshold.
	//
	// The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage.
	//
	// To turn off price protection, specify a high value, such as 999999.
	//
	// If you set DesiredCapacityType to vcpu or memory-mib, the price protection threshold is applied based on the per-vCPU or per-memory price instead of the per instance price.
	// Default: - 20.
	//
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Indicates whether instance types must provide On-Demand Instance hibernation support.
	// Default: - false.
	//
	RequireHibernateSupport *bool `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// [Price protection] The price protection threshold for Spot Instances, as a percentage higher than an identified Spot price.
	//
	// The identified Spot price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from either the lowest priced current generation instance types or, failing that, the lowest priced previous generation instance types that match your attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price exceeds your specified threshold.
	//
	// The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage.
	//
	// If you set DesiredCapacityType to vcpu or memory-mib, the price protection threshold is based on the per-vCPU or per-memory price instead of the per instance price.
	//
	// Note: Only one of SpotMaxPricePercentageOverLowestPrice or MaxSpotPriceAsPercentageOfOptimalOnDemandPrice can be specified. If you don't specify either, Amazon EC2 Auto Scaling will automatically apply optimal price protection to consistently select from a wide range of instance types. To indicate no price protection threshold for Spot Instances, meaning you want to consider all instance types that match your attributes, include one of these parameters and specify a high value, such as 999999.
	// Default: - Automatic optimal price protection.
	//
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// The maximum total local storage size for an instance type, in GB.
	// Default: - No minimum or maximum limits.
	//
	TotalLocalStorageGBMax *float64 `field:"optional" json:"totalLocalStorageGBMax" yaml:"totalLocalStorageGBMax"`
	// The minimum total local storage size for an instance type, in GB.
	// Default: - No minimum or maximum limits.
	//
	TotalLocalStorageGBMin *float64 `field:"optional" json:"totalLocalStorageGBMin" yaml:"totalLocalStorageGBMin"`
	// The maximum number of vCPUs for an instance type.
	// Default: - No maximum limit.
	//
	VCpuCountMax *float64 `field:"optional" json:"vCpuCountMax" yaml:"vCpuCountMax"`
}


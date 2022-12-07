package awsautoscaling


// `InstanceRequirements` specifies a set of requirements for the types of instances that can be launched by an `AWS::AutoScaling::AutoScalingGroup` resource.
//
// `InstanceRequirements` is a property of the `LaunchTemplateOverrides` property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html) property type.
//
// You must specify `VCpuCount` and `MemoryMiB` , but all other properties are optional. Any unspecified optional property is set to its default.
//
// When you specify multiple properties, you get instance types that satisfy all of the specified properties. If you specify multiple values for a property, you get instance types that satisfy any of the specified values.
//
// For an example template, see [Auto scaling template snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-autoscaling.html) .
//
// For more information, see [Create an Auto Scaling group using attribute-based instance type selection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-instance-type-requirements.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceRequirementsProperty := &instanceRequirementsProperty{
//   	acceleratorCount: &acceleratorCountRequestProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	acceleratorManufacturers: []*string{
//   		jsii.String("acceleratorManufacturers"),
//   	},
//   	acceleratorNames: []*string{
//   		jsii.String("acceleratorNames"),
//   	},
//   	acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	acceleratorTypes: []*string{
//   		jsii.String("acceleratorTypes"),
//   	},
//   	allowedInstanceTypes: []*string{
//   		jsii.String("allowedInstanceTypes"),
//   	},
//   	bareMetal: jsii.String("bareMetal"),
//   	baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	burstablePerformance: jsii.String("burstablePerformance"),
//   	cpuManufacturers: []*string{
//   		jsii.String("cpuManufacturers"),
//   	},
//   	excludedInstanceTypes: []*string{
//   		jsii.String("excludedInstanceTypes"),
//   	},
//   	instanceGenerations: []*string{
//   		jsii.String("instanceGenerations"),
//   	},
//   	localStorage: jsii.String("localStorage"),
//   	localStorageTypes: []*string{
//   		jsii.String("localStorageTypes"),
//   	},
//   	memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	memoryMiB: &memoryMiBRequestProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	networkBandwidthGbps: &networkBandwidthGbpsRequestProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   	requireHibernateSupport: jsii.Boolean(false),
//   	spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   	totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	vCpuCount: &vCpuCountRequestProperty{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   }
//
type CfnAutoScalingGroup_InstanceRequirementsProperty struct {
	// The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) for an instance type.
	//
	// To exclude accelerator-enabled instance types, set `Max` to `0` .
	//
	// Default: No minimum or maximum.
	AcceleratorCount interface{} `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Indicates whether instance types must have accelerators by specific manufacturers.
	//
	// - For instance types with NVIDIA devices, specify `nvidia` .
	// - For instance types with AMD devices, specify `amd` .
	// - For instance types with AWS devices, specify `amazon-web-services` .
	// - For instance types with Xilinx devices, specify `xilinx` .
	//
	// Default: Any manufacturer.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Lists the accelerators that must be on an instance type.
	//
	// - For instance types with NVIDIA A100 GPUs, specify `a100` .
	// - For instance types with NVIDIA V100 GPUs, specify `v100` .
	// - For instance types with NVIDIA K80 GPUs, specify `k80` .
	// - For instance types with NVIDIA T4 GPUs, specify `t4` .
	// - For instance types with NVIDIA M60 GPUs, specify `m60` .
	// - For instance types with AMD Radeon Pro V520 GPUs, specify `radeon-pro-v520` .
	// - For instance types with Xilinx VU9P FPGAs, specify `vu9p` .
	//
	// Default: Any accelerator.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// The minimum and maximum total memory size for the accelerators on an instance type, in MiB.
	//
	// Default: No minimum or maximum.
	AcceleratorTotalMemoryMiB interface{} `field:"optional" json:"acceleratorTotalMemoryMiB" yaml:"acceleratorTotalMemoryMiB"`
	// Lists the accelerator types that must be on an instance type.
	//
	// - For instance types with GPU accelerators, specify `gpu` .
	// - For instance types with FPGA accelerators, specify `fpga` .
	// - For instance types with inference accelerators, specify `inference` .
	//
	// Default: Any accelerator type.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// `CfnAutoScalingGroup.InstanceRequirementsProperty.AllowedInstanceTypes`.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Indicates whether bare metal instance types are included, excluded, or required.
	//
	// Default: `excluded`.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// The minimum and maximum baseline bandwidth performance for an instance type, in Mbps.
	//
	// For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// Default: No minimum or maximum.
	BaselineEbsBandwidthMbps interface{} `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Indicates whether burstable performance instance types are included, excluded, or required.
	//
	// For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// Default: `excluded`.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Lists which specific CPU manufacturers to include.
	//
	// - For instance types with Intel CPUs, specify `intel` .
	// - For instance types with AMD CPUs, specify `amd` .
	// - For instance types with AWS CPUs, specify `amazon-web-services` .
	//
	// > Don't confuse the CPU hardware manufacturer with the CPU hardware architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template.
	//
	// Default: Any manufacturer.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// Lists which instance types to exclude.
	//
	// You can use strings with one or more wild cards, represented by an asterisk ( `*` ). The following are examples: `c5*` , `m5a.*` , `r*` , `*3*` .
	//
	// For example, if you specify `c5*` , you are excluding the entire C5 instance family, which includes all C5a and C5n instance types. If you specify `m5a.*` , you are excluding all the M5a instance types, but not the M5n instance types.
	//
	// Default: No excluded instance types.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Indicates whether current or previous generation instance types are included.
	//
	// - For current generation instance types, specify `current` . The current generation includes EC2 instance types currently recommended for use. This typically includes the latest two to three generations in each instance family. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide for Linux Instances* .
	// - For previous generation instance types, specify `previous` .
	//
	// Default: Any current or previous generation.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Indicates whether instance types with instance store volumes are included, excluded, or required.
	//
	// For more information, see [Amazon EC2 instance store](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html) in the *Amazon EC2 User Guide for Linux Instances* .
	//
	// Default: `included`.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Indicates the type of local storage that is required.
	//
	// - For instance types with hard disk drive (HDD) storage, specify `hdd` .
	// - For instance types with solid state drive (SSD) storage, specify `sdd` .
	//
	// Default: Any local storage type.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// The minimum and maximum amount of memory per vCPU for an instance type, in GiB.
	//
	// Default: No minimum or maximum.
	MemoryGiBPerVCpu interface{} `field:"optional" json:"memoryGiBPerVCpu" yaml:"memoryGiBPerVCpu"`
	// The minimum and maximum instance memory size for an instance type, in MiB.
	MemoryMiB interface{} `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// `CfnAutoScalingGroup.InstanceRequirementsProperty.NetworkBandwidthGbps`.
	NetworkBandwidthGbps interface{} `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// The minimum and maximum number of network interfaces for an instance type.
	//
	// Default: No minimum or maximum.
	NetworkInterfaceCount interface{} `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// The price protection threshold for On-Demand Instances.
	//
	// This is the maximum you’ll pay for an On-Demand Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as `999999` .
	//
	// If you set `DesiredCapacityType` to `vcpu` or `memory-mib` , the price protection threshold is applied based on the per vCPU or per memory price instead of the per instance price.
	//
	// Default: `20`.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Indicates whether instance types must provide On-Demand Instance hibernation support.
	//
	// Default: `false`.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// The price protection threshold for Spot Instances.
	//
	// This is the maximum you’ll pay for a Spot Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as `999999` .
	//
	// If you set `DesiredCapacityType` to `vcpu` or `memory-mib` , the price protection threshold is applied based on the per vCPU or per memory price instead of the per instance price.
	//
	// Default: `100`.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// The minimum and maximum total local storage size for an instance type, in GB.
	//
	// Default: No minimum or maximum.
	TotalLocalStorageGb interface{} `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
	// The minimum and maximum number of vCPUs for an instance type.
	VCpuCount interface{} `field:"optional" json:"vCpuCount" yaml:"vCpuCount"`
}


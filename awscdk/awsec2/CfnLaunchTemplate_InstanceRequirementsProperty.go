package awsec2


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
// For more information, see [Specify attributes for instance type selection for EC2 Fleet or Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-attribute-based-instance-type-selection.html) and [Spot placement score](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-placement-score.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceRequirementsProperty := &InstanceRequirementsProperty{
//   	AcceleratorCount: &AcceleratorCountProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	AcceleratorManufacturers: []*string{
//   		jsii.String("acceleratorManufacturers"),
//   	},
//   	AcceleratorNames: []*string{
//   		jsii.String("acceleratorNames"),
//   	},
//   	AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	AcceleratorTypes: []*string{
//   		jsii.String("acceleratorTypes"),
//   	},
//   	AllowedInstanceTypes: []*string{
//   		jsii.String("allowedInstanceTypes"),
//   	},
//   	BareMetal: jsii.String("bareMetal"),
//   	BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	BaselinePerformanceFactors: &BaselinePerformanceFactorsProperty{
//   		Cpu: &CpuProperty{
//   			References: []interface{}{
//   				&ReferenceProperty{
//   					InstanceFamily: jsii.String("instanceFamily"),
//   				},
//   			},
//   		},
//   	},
//   	BurstablePerformance: jsii.String("burstablePerformance"),
//   	CpuManufacturers: []*string{
//   		jsii.String("cpuManufacturers"),
//   	},
//   	ExcludedInstanceTypes: []*string{
//   		jsii.String("excludedInstanceTypes"),
//   	},
//   	InstanceGenerations: []*string{
//   		jsii.String("instanceGenerations"),
//   	},
//   	LocalStorage: jsii.String("localStorage"),
//   	LocalStorageTypes: []*string{
//   		jsii.String("localStorageTypes"),
//   	},
//   	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   	MemoryGiBPerVCpu: &MemoryGiBPerVCpuProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	MemoryMiB: &MemoryMiBProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	NetworkBandwidthGbps: &NetworkBandwidthGbpsProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	NetworkInterfaceCount: &NetworkInterfaceCountProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   	RequireHibernateSupport: jsii.Boolean(false),
//   	SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   	TotalLocalStorageGb: &TotalLocalStorageGBProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	VCpuCount: &VCpuCountProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html
//
type CfnLaunchTemplate_InstanceRequirementsProperty struct {
	// The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) on an instance.
	//
	// To exclude accelerator-enabled instance types, set `Max` to `0` .
	//
	// Default: No minimum or maximum limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-acceleratorcount
	//
	AcceleratorCount interface{} `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Indicates whether instance types must have accelerators by specific manufacturers.
	//
	// - For instance types with AWS devices, specify `amazon-web-services` .
	// - For instance types with AMD devices, specify `amd` .
	// - For instance types with Habana devices, specify `habana` .
	// - For instance types with NVIDIA devices, specify `nvidia` .
	// - For instance types with Xilinx devices, specify `xilinx` .
	//
	// Default: Any manufacturer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-acceleratormanufacturers
	//
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// The accelerators that must be on the instance type.
	//
	// - For instance types with NVIDIA A10G GPUs, specify `a10g` .
	// - For instance types with NVIDIA A100 GPUs, specify `a100` .
	// - For instance types with NVIDIA H100 GPUs, specify `h100` .
	// - For instance types with AWS Inferentia chips, specify `inferentia` .
	// - For instance types with NVIDIA GRID K520 GPUs, specify `k520` .
	// - For instance types with NVIDIA K80 GPUs, specify `k80` .
	// - For instance types with NVIDIA M60 GPUs, specify `m60` .
	// - For instance types with AMD Radeon Pro V520 GPUs, specify `radeon-pro-v520` .
	// - For instance types with NVIDIA T4 GPUs, specify `t4` .
	// - For instance types with NVIDIA T4G GPUs, specify `t4g` .
	// - For instance types with Xilinx VU9P FPGAs, specify `vu9p` .
	// - For instance types with NVIDIA V100 GPUs, specify `v100` .
	//
	// Default: Any accelerator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-acceleratornames
	//
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// The minimum and maximum amount of total accelerator memory, in MiB.
	//
	// Default: No minimum or maximum limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-acceleratortotalmemorymib
	//
	AcceleratorTotalMemoryMiB interface{} `field:"optional" json:"acceleratorTotalMemoryMiB" yaml:"acceleratorTotalMemoryMiB"`
	// The accelerator types that must be on the instance type.
	//
	// - For instance types with FPGA accelerators, specify `fpga` .
	// - For instance types with GPU accelerators, specify `gpu` .
	//
	// Default: Any accelerator type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-acceleratortypes
	//
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// The instance types to apply your specified attributes against.
	//
	// All other instance types are ignored, even if they match your specified attributes.
	//
	// You can use strings with one or more wild cards, represented by an asterisk ( `*` ), to allow an instance type, size, or generation. The following are examples: `m5.8xlarge` , `c5*.*` , `m5a.*` , `r*` , `*3*` .
	//
	// For example, if you specify `c5*` ,Amazon EC2 will allow the entire C5 instance family, which includes all C5a and C5n instance types. If you specify `m5a.*` , Amazon EC2 will allow all the M5a instance types, but not the M5n instance types.
	//
	// > If you specify `AllowedInstanceTypes` , you can't specify `ExcludedInstanceTypes` .
	//
	// Default: All instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-allowedinstancetypes
	//
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Indicates whether bare metal instance types must be included, excluded, or required.
	//
	// - To include bare metal instance types, specify `included` .
	// - To require only bare metal instance types, specify `required` .
	// - To exclude bare metal instance types, specify `excluded` .
	//
	// Default: `excluded`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-baremetal
	//
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps.
	//
	// For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html) in the *Amazon EC2 User Guide* .
	//
	// Default: No minimum or maximum limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-baselineebsbandwidthmbps
	//
	BaselineEbsBandwidthMbps interface{} `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// The baseline performance to consider, using an instance family as a baseline reference.
	//
	// The instance family establishes the lowest acceptable level of performance. Amazon EC2 uses this baseline to guide instance type selection, but there is no guarantee that the selected instance types will always exceed the baseline for every application. Currently, this parameter only supports CPU performance as a baseline performance factor. For more information, see [Performance protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-attribute-based-instance-type-selection.html#ec2fleet-abis-performance-protection) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-baselineperformancefactors
	//
	BaselinePerformanceFactors interface{} `field:"optional" json:"baselinePerformanceFactors" yaml:"baselinePerformanceFactors"`
	// Indicates whether burstable performance T instance types are included, excluded, or required.
	//
	// For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) .
	//
	// - To include burstable performance instance types, specify `included` .
	// - To require only burstable performance instance types, specify `required` .
	// - To exclude burstable performance instance types, specify `excluded` .
	//
	// Default: `excluded`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-burstableperformance
	//
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// The CPU manufacturers to include.
	//
	// - For instance types with Intel CPUs, specify `intel` .
	// - For instance types with AMD CPUs, specify `amd` .
	// - For instance types with AWS CPUs, specify `amazon-web-services` .
	// - For instance types with Apple CPUs, specify `apple` .
	//
	// > Don't confuse the CPU manufacturer with the CPU architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template.
	//
	// Default: Any manufacturer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-cpumanufacturers
	//
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// The instance types to exclude.
	//
	// You can use strings with one or more wild cards, represented by an asterisk ( `*` ), to exclude an instance type, size, or generation. The following are examples: `m5.8xlarge` , `c5*.*` , `m5a.*` , `r*` , `*3*` .
	//
	// For example, if you specify `c5*` ,Amazon EC2 will exclude the entire C5 instance family, which includes all C5a and C5n instance types. If you specify `m5a.*` , Amazon EC2 will exclude all the M5a instance types, but not the M5n instance types.
	//
	// > If you specify `ExcludedInstanceTypes` , you can't specify `AllowedInstanceTypes` .
	//
	// Default: No excluded instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-excludedinstancetypes
	//
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Indicates whether current or previous generation instance types are included.
	//
	// The current generation instance types are recommended for use. Current generation instance types are typically the latest two to three generations in each instance family. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide* .
	//
	// For current generation instance types, specify `current` .
	//
	// For previous generation instance types, specify `previous` .
	//
	// Default: Current and previous generation instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-instancegenerations
	//
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Indicates whether instance types with instance store volumes are included, excluded, or required.
	//
	// For more information, [Amazon EC2 instance store](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html) in the *Amazon EC2 User Guide* .
	//
	// - To include instance types with instance store volumes, specify `included` .
	// - To require only instance types with instance store volumes, specify `required` .
	// - To exclude instance types with instance store volumes, specify `excluded` .
	//
	// Default: `included`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-localstorage
	//
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// The type of local storage that is required.
	//
	// - For instance types with hard disk drive (HDD) storage, specify `hdd` .
	// - For instance types with solid state drive (SSD) storage, specify `ssd` .
	//
	// Default: `hdd` and `ssd`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-localstoragetypes
	//
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// [Price protection] The price protection threshold for Spot Instances, as a percentage of an identified On-Demand price.
	//
	// The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from the lowest priced current generation instance types, and failing that, from the lowest priced previous generation instance types that match your attributes. When Amazon EC2 selects instance types with your attributes, it will exclude instance types whose price exceeds your specified threshold.
	//
	// The parameter accepts an integer, which Amazon EC2 interprets as a percentage.
	//
	// If you set `TargetCapacityUnitType` to `vcpu` or `memory-mib` , the price protection threshold is based on the per vCPU or per memory price instead of the per instance price.
	//
	// > Only one of `SpotMaxPricePercentageOverLowestPrice` or `MaxSpotPriceAsPercentageOfOptimalOnDemandPrice` can be specified. If you don't specify either, Amazon EC2 will automatically apply optimal price protection to consistently select from a wide range of instance types. To indicate no price protection threshold for Spot Instances, meaning you want to consider all instance types that match your attributes, include one of these parameters and specify a high value, such as `999999` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-maxspotpriceaspercentageofoptimalondemandprice
	//
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// The minimum and maximum amount of memory per vCPU, in GiB.
	//
	// Default: No minimum or maximum limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-memorygibpervcpu
	//
	MemoryGiBPerVCpu interface{} `field:"optional" json:"memoryGiBPerVCpu" yaml:"memoryGiBPerVCpu"`
	// The minimum and maximum amount of memory, in MiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-memorymib
	//
	MemoryMiB interface{} `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).
	//
	// Default: No minimum or maximum limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-networkbandwidthgbps
	//
	NetworkBandwidthGbps interface{} `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// The minimum and maximum number of network interfaces.
	//
	// Default: No minimum or maximum limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-networkinterfacecount
	//
	NetworkInterfaceCount interface{} `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// [Price protection] The price protection threshold for On-Demand Instances, as a percentage higher than an identified On-Demand price.
	//
	// The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. When Amazon EC2 selects instance types with your attributes, it will exclude instance types whose price exceeds your specified threshold.
	//
	// The parameter accepts an integer, which Amazon EC2 interprets as a percentage.
	//
	// To turn off price protection, specify a high value, such as `999999` .
	//
	// This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html) and [GetInstanceTypesFromInstanceRequirements](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceTypesFromInstanceRequirements.html) .
	//
	// > If you set `TargetCapacityUnitType` to `vcpu` or `memory-mib` , the price protection threshold is applied based on the per-vCPU or per-memory price instead of the per-instance price.
	//
	// Default: `20`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-ondemandmaxpricepercentageoverlowestprice
	//
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Indicates whether instance types must support hibernation for On-Demand Instances.
	//
	// This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html) .
	//
	// Default: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-requirehibernatesupport
	//
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// [Price protection] The price protection threshold for Spot Instances, as a percentage higher than an identified Spot price.
	//
	// The identified Spot price is the Spot price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified Spot price is from the lowest priced current generation instance types, and failing that, from the lowest priced previous generation instance types that match your attributes. When Amazon EC2 selects instance types with your attributes, it will exclude instance types whose Spot price exceeds your specified threshold.
	//
	// The parameter accepts an integer, which Amazon EC2 interprets as a percentage.
	//
	// If you set `TargetCapacityUnitType` to `vcpu` or `memory-mib` , the price protection threshold is applied based on the per-vCPU or per-memory price instead of the per-instance price.
	//
	// This parameter is not supported for [GetSpotPlacementScores](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetSpotPlacementScores.html) and [GetInstanceTypesFromInstanceRequirements](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceTypesFromInstanceRequirements.html) .
	//
	// > Only one of `SpotMaxPricePercentageOverLowestPrice` or `MaxSpotPriceAsPercentageOfOptimalOnDemandPrice` can be specified. If you don't specify either, Amazon EC2 will automatically apply optimal price protection to consistently select from a wide range of instance types. To indicate no price protection threshold for Spot Instances, meaning you want to consider all instance types that match your attributes, include one of these parameters and specify a high value, such as `999999` .
	//
	// Default: `100`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-spotmaxpricepercentageoverlowestprice
	//
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// The minimum and maximum amount of total local storage, in GB.
	//
	// Default: No minimum or maximum limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-totallocalstoragegb
	//
	TotalLocalStorageGb interface{} `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
	// The minimum and maximum number of vCPUs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-instancerequirements.html#cfn-ec2-launchtemplate-instancerequirements-vcpucount
	//
	VCpuCount interface{} `field:"optional" json:"vCpuCount" yaml:"vCpuCount"`
}


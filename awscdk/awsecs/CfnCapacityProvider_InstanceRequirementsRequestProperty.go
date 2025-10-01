package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceRequirementsRequestProperty := &InstanceRequirementsRequestProperty{
//   	MemoryMiB: &MemoryMiBRequestProperty{
//   		Min: jsii.Number(123),
//
//   		// the properties below are optional
//   		Max: jsii.Number(123),
//   	},
//   	VCpuCount: &VCpuCountRangeRequestProperty{
//   		Min: jsii.Number(123),
//
//   		// the properties below are optional
//   		Max: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	AcceleratorCount: &AcceleratorCountRequestProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	AcceleratorManufacturers: []*string{
//   		jsii.String("acceleratorManufacturers"),
//   	},
//   	AcceleratorNames: []*string{
//   		jsii.String("acceleratorNames"),
//   	},
//   	AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
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
//   	BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
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
//   	MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   	RequireHibernateSupport: jsii.Boolean(false),
//   	SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   	TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html
//
type CfnCapacityProvider_InstanceRequirementsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-memorymib
	//
	MemoryMiB interface{} `field:"required" json:"memoryMiB" yaml:"memoryMiB"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-vcpucount
	//
	VCpuCount interface{} `field:"required" json:"vCpuCount" yaml:"vCpuCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-acceleratorcount
	//
	AcceleratorCount interface{} `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-acceleratormanufacturers
	//
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-acceleratornames
	//
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-acceleratortotalmemorymib
	//
	AcceleratorTotalMemoryMiB interface{} `field:"optional" json:"acceleratorTotalMemoryMiB" yaml:"acceleratorTotalMemoryMiB"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-acceleratortypes
	//
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-allowedinstancetypes
	//
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-baremetal
	//
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-baselineebsbandwidthmbps
	//
	BaselineEbsBandwidthMbps interface{} `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-burstableperformance
	//
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-cpumanufacturers
	//
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-excludedinstancetypes
	//
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-instancegenerations
	//
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-localstorage
	//
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-localstoragetypes
	//
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-maxspotpriceaspercentageofoptimalondemandprice
	//
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-memorygibpervcpu
	//
	MemoryGiBPerVCpu interface{} `field:"optional" json:"memoryGiBPerVCpu" yaml:"memoryGiBPerVCpu"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-networkbandwidthgbps
	//
	NetworkBandwidthGbps interface{} `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-networkinterfacecount
	//
	NetworkInterfaceCount interface{} `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-ondemandmaxpricepercentageoverlowestprice
	//
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-requirehibernatesupport
	//
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-spotmaxpricepercentageoverlowestprice
	//
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-instancerequirementsrequest.html#cfn-ecs-capacityprovider-instancerequirementsrequest-totallocalstoragegb
	//
	TotalLocalStorageGb interface{} `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
}


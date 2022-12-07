package awsec2


// Specifies overrides for a launch template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateOverridesProperty := &launchTemplateOverridesProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
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
//   	priority: jsii.Number(123),
//   	spotPrice: jsii.String("spotPrice"),
//   	subnetId: jsii.String("subnetId"),
//   	weightedCapacity: jsii.Number(123),
//   }
//
type CfnSpotFleet_LaunchTemplateOverridesProperty struct {
	// The Availability Zone in which to launch the instances.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The instance requirements.
	//
	// When you specify instance requirements, Amazon EC2 will identify instance types with the provided requirements, and then use your On-Demand and Spot allocation strategies to launch instances from these instance types, in the same way as when you specify a list of instance types.
	//
	// > If you specify `InstanceRequirements` , you can't specify `InstanceTypes` .
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The priority for the launch template override. The highest priority is launched first.
	//
	// If `OnDemandAllocationStrategy` is set to `prioritized` , Spot Fleet uses priority to determine which launch template override to use first in fulfilling On-Demand capacity.
	//
	// If the Spot `AllocationStrategy` is set to `capacityOptimizedPrioritized` , Spot Fleet uses priority on a best-effort basis to determine which launch template override to use in fulfilling Spot capacity, but optimizes for capacity first.
	//
	// Valid values are whole numbers starting at `0` . The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority. You can set the same priority for different launch template overrides.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The maximum price per unit hour that you are willing to pay for a Spot Instance.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// The ID of the subnet in which to launch the instances.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The number of units provided by the specified instance type.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


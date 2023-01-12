package awsec2


// Specifies overrides for a launch template for an EC2 Fleet.
//
// `FleetLaunchTemplateOverridesRequest` is a property of the [FleetLaunchTemplateConfigRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplateconfigrequest.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetLaunchTemplateOverridesRequestProperty := &fleetLaunchTemplateOverridesRequestProperty{
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
//   	maxPrice: jsii.String("maxPrice"),
//   	placement: &placementProperty{
//   		affinity: jsii.String("affinity"),
//   		availabilityZone: jsii.String("availabilityZone"),
//   		groupName: jsii.String("groupName"),
//   		hostId: jsii.String("hostId"),
//   		hostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   		partitionNumber: jsii.Number(123),
//   		spreadDomain: jsii.String("spreadDomain"),
//   		tenancy: jsii.String("tenancy"),
//   	},
//   	priority: jsii.Number(123),
//   	subnetId: jsii.String("subnetId"),
//   	weightedCapacity: jsii.Number(123),
//   }
//
type CfnEC2Fleet_FleetLaunchTemplateOverridesRequestProperty struct {
	// The Availability Zone in which to launch the instances.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The attributes for the instance types.
	//
	// When you specify instance attributes, Amazon EC2 will identify instance types with those attributes.
	//
	// > If you specify `InstanceRequirements` , you can't specify `InstanceTypes` .
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type.
	//
	// > If you specify `InstanceTypes` , you can't specify `InstanceRequirements` .
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The maximum price per unit hour that you are willing to pay for a Spot Instance.
	MaxPrice *string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// The location where the instance launched, if applicable.
	Placement interface{} `field:"optional" json:"placement" yaml:"placement"`
	// The priority for the launch template override. The highest priority is launched first.
	//
	// If the On-Demand `AllocationStrategy` is set to `prioritized` , EC2 Fleet uses priority to determine which launch template override to use first in fulfilling On-Demand capacity.
	//
	// If the Spot `AllocationStrategy` is set to `capacity-optimized-prioritized` , EC2 Fleet uses priority on a best-effort basis to determine which launch template override to use in fulfilling Spot capacity, but optimizes for capacity first.
	//
	// Valid values are whole numbers starting at `0` . The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority. You can set the same priority for different launch template overrides.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The IDs of the subnets in which to launch the instances.
	//
	// Separate multiple subnet IDs using commas (for example, `subnet-1234abcdeexample1, subnet-0987cdef6example2` ). A request of type `instant` can have only one subnet ID.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The number of units provided by the specified instance type.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


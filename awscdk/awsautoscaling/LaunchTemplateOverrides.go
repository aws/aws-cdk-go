package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// LaunchTemplateOverrides is a subproperty of LaunchTemplate that describes an override for a launch template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instanceType instanceType
//   var launchTemplate launchTemplate
//
//   launchTemplateOverrides := &LaunchTemplateOverrides{
//   	InstanceRequirements: &InstanceRequirementsProperty{
//   		MemoryMiB: &MemoryMiBRequestProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		VCpuCount: &VCpuCountRequestProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
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
//   	},
//   	InstanceType: instanceType,
//   	LaunchTemplate: launchTemplate,
//   	WeightedCapacity: jsii.Number(123),
//   }
//
type LaunchTemplateOverrides struct {
	// The instance requirements.
	//
	// Amazon EC2 Auto Scaling uses your specified requirements to identify instance types.
	// Then, it uses your On-Demand and Spot allocation strategies to launch instances from these instance types.
	//
	// You can specify up to four separate sets of instance requirements per Auto Scaling group.
	// This is useful for provisioning instances from different Amazon Machine Images (AMIs) in the same Auto Scaling group.
	// To do this, create the AMIs and create a new launch template for each AMI.
	// Then, create a compatible set of instance requirements for each launch template.
	//
	// You must specify one of instanceRequirements or instanceType.
	// Default: - Do not override instance type.
	//
	InstanceRequirements *CfnAutoScalingGroup_InstanceRequirementsProperty `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type, such as m3.xlarge. You must use an instance type that is supported in your requested Region and Availability Zones.
	//
	// You must specify one of instanceRequirements or instanceType.
	// Default: - Do not override instance type.
	//
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Provides the launch template to be used when launching the instance type.
	//
	// For example, some instance types might
	// require a launch template with a different AMI. If not provided, Amazon EC2 Auto Scaling uses the launch template
	// that's defined for your mixed instances policy.
	// Default: - Do not override launch template.
	//
	LaunchTemplate awsec2.ILaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The number of capacity units provided by the specified instance type in terms of virtual CPUs, memory, storage, throughput, or other relative performance characteristic.
	//
	// When a Spot or On-Demand Instance is provisioned, the
	// capacity units count toward the desired capacity. Amazon EC2 Auto Scaling provisions instances until the desired
	// capacity is totally fulfilled, even if this results in an overage. Value must be in the range of 1 to 999.
	//
	// For example, If there are 2 units remaining to fulfill capacity, and Amazon EC2 Auto Scaling can only provision
	// an instance with a WeightedCapacity of 5 units, the instance is provisioned, and the desired capacity is exceeded
	// by 3 units.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-weighting.html
	//
	// Default: - Do not provide weight.
	//
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


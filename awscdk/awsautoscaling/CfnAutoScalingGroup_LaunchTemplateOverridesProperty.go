package awsautoscaling


// Use this structure to let Amazon EC2 Auto Scaling do the following when the Auto Scaling group has a mixed instances policy:  - Override the instance type that is specified in the launch template.
//
// - Use multiple instance types.
//
// Specify the instance types that you want, or define your instance requirements instead and let Amazon EC2 Auto Scaling provision the available instance types that meet your requirements. This can provide Amazon EC2 Auto Scaling with a larger selection of instance types to choose from when fulfilling Spot and On-Demand capacities. You can view which instance types are matched before you apply the instance requirements to your Auto Scaling group.
//
// After you define your instance requirements, you don't have to keep updating these settings to get new EC2 instance types automatically. Amazon EC2 Auto Scaling uses the instance requirements of the Auto Scaling group to determine whether a new EC2 instance type can be used.
//
// `LaunchTemplateOverrides` is a property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateOverridesProperty := &LaunchTemplateOverridesProperty{
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
//   	InstanceType: jsii.String("instanceType"),
//   	LaunchTemplateSpecification: &LaunchTemplateSpecificationProperty{
//   		Version: jsii.String("version"),
//
//   		// the properties below are optional
//   		LaunchTemplateId: jsii.String("launchTemplateId"),
//   		LaunchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	WeightedCapacity: jsii.String("weightedCapacity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html
//
type CfnAutoScalingGroup_LaunchTemplateOverridesProperty struct {
	// The instance requirements.
	//
	// Amazon EC2 Auto Scaling uses your specified requirements to identify instance types. Then, it uses your On-Demand and Spot allocation strategies to launch instances from these instance types.
	//
	// You can specify up to four separate sets of instance requirements per Auto Scaling group. This is useful for provisioning instances from different Amazon Machine Images (AMIs) in the same Auto Scaling group. To do this, create the AMIs and create a new launch template for each AMI. Then, create a compatible set of instance requirements for each launch template.
	//
	// > If you specify `InstanceRequirements` , you can't specify `InstanceType` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html#cfn-autoscaling-autoscalinggroup-launchtemplateoverrides-instancerequirements
	//
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type, such as `m3.xlarge` . You must specify an instance type that is supported in your requested Region and Availability Zones. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon Elastic Compute Cloud User Guide* .
	//
	// You can specify up to 40 instance types per Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html#cfn-autoscaling-autoscalinggroup-launchtemplateoverrides-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Provides a launch template for the specified instance type or set of instance requirements.
	//
	// For example, some instance types might require a launch template with a different AMI. If not provided, Amazon EC2 Auto Scaling uses the launch template that's specified in the `LaunchTemplate` definition. For more information, see [Specifying a different launch template for an instance type](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-launch-template-overrides.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// You can specify up to 20 launch templates per Auto Scaling group. The launch templates specified in the overrides and in the `LaunchTemplate` definition count towards this limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html#cfn-autoscaling-autoscalinggroup-launchtemplateoverrides-launchtemplatespecification
	//
	LaunchTemplateSpecification interface{} `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// If you provide a list of instance types to use, you can specify the number of capacity units provided by each instance type in terms of virtual CPUs, memory, storage, throughput, or other relative performance characteristic.
	//
	// When a Spot or On-Demand Instance is launched, the capacity units count toward the desired capacity. Amazon EC2 Auto Scaling launches instances until the desired capacity is totally fulfilled, even if this results in an overage. For example, if there are two units remaining to fulfill capacity, and Amazon EC2 Auto Scaling can only launch an instance with a `WeightedCapacity` of five units, the instance is launched, and the desired capacity is exceeded by three units. For more information, see [Configure instance weighting for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-instance-weighting.html) in the *Amazon EC2 Auto Scaling User Guide* . Value must be in the range of 1-999.
	//
	// If you specify a value for `WeightedCapacity` for one instance type, you must specify a value for `WeightedCapacity` for all of them.
	//
	// > Every Auto Scaling group has three size parameters ( `DesiredCapacity` , `MaxSize` , and `MinSize` ). Usually, you set these sizes based on a specific number of instances. However, if you configure a mixed instances policy that defines weights for the instance types, you must specify these sizes with the same units that you use for weighting instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.html#cfn-autoscaling-autoscalinggroup-launchtemplateoverrides-weightedcapacity
	//
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


package awsautoscaling


// `LaunchTemplateOverrides` is a property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html) property type that describes an override for a launch template.
//
// If you supply your own instance types, the maximum number of instance types that can be associated with an Auto Scaling group is 40. The maximum number of distinct launch templates you can define for an Auto Scaling group is 20.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateOverridesProperty := &launchTemplateOverridesProperty{
//   	instanceRequirements: &instanceRequirementsProperty{
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
//   		vCpuCount: &vCpuCountRequestProperty{
//   			max: jsii.Number(123),
//   			min: jsii.Number(123),
//   		},
//   	},
//   	instanceType: jsii.String("instanceType"),
//   	launchTemplateSpecification: &launchTemplateSpecificationProperty{
//   		version: jsii.String("version"),
//
//   		// the properties below are optional
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	weightedCapacity: jsii.String("weightedCapacity"),
//   }
//
type CfnAutoScalingGroup_LaunchTemplateOverridesProperty struct {
	// The instance requirements.
	//
	// When you specify instance requirements, Amazon EC2 Auto Scaling finds instance types that satisfy your requirements, and then uses your On-Demand and Spot allocation strategies to launch instances from these instance types, in the same way as when you specify a list of specific instance types.
	//
	// > `InstanceRequirements` are incompatible with the `InstanceType` and `WeightedCapacity` properties.
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type, such as `m3.xlarge` . You must use an instance type that is supported in your requested Region and Availability Zones. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon Elastic Compute Cloud User Guide* .
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Provides a launch template for the specified instance type or instance requirements.
	//
	// For example, some instance types might require a launch template with a different AMI. If not provided, Amazon EC2 Auto Scaling uses the launch template that's defined for your mixed instances policy. For more information, see [Specifying a different launch template for an instance type](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-launch-template-overrides.html) in the *Amazon EC2 Auto Scaling User Guide* .
	LaunchTemplateSpecification interface{} `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// The number of capacity units provided by the instance type specified in `InstanceType` in terms of virtual CPUs, memory, storage, throughput, or other relative performance characteristic.
	//
	// When a Spot or On-Demand Instance is provisioned, the capacity units count toward the desired capacity. Amazon EC2 Auto Scaling provisions instances until the desired capacity is totally fulfilled, even if this results in an overage. For example, if there are 2 units remaining to fulfill capacity, and Amazon EC2 Auto Scaling can only provision an instance with a `WeightedCapacity` of 5 units, the instance is provisioned, and the desired capacity is exceeded by 3 units. For more information, see [Configure instance weighting for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-instance-weighting.html) in the *Amazon EC2 Auto Scaling User Guide* . Value must be in the range of 1-999.
	//
	// > Every Auto Scaling group has three size parameters ( `DesiredCapacity` , `MaxSize` , and `MinSize` ). Usually, you set these sizes based on a specific number of instances. However, if you configure a mixed instances policy that defines weights for the instance types, you must specify these sizes with the same units that you use for weighting instances.
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}


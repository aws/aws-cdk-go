package awsautoscaling


// Use this structure to launch multiple instance types and On-Demand Instances and Spot Instances within a single Auto Scaling group.
//
// A mixed instances policy contains information that Amazon EC2 Auto Scaling can use to launch instances and help optimize your costs. For more information, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// You can create a mixed instances policy for new and existing Auto Scaling groups. You must use a launch template to configure the policy. You cannot use a launch configuration.
//
// There are key differences between Spot Instances and On-Demand Instances:
//
// - The price for Spot Instances varies based on demand
// - Amazon EC2 can terminate an individual Spot Instance as the availability of, or price for, Spot Instances changes
//
// When a Spot Instance is terminated, Amazon EC2 Auto Scaling group attempts to launch a replacement instance to maintain the desired capacity for the group.
//
// `MixedInstancesPolicy` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mixedInstancesPolicyProperty := &MixedInstancesPolicyProperty{
//   	LaunchTemplate: &LaunchTemplateProperty{
//   		LaunchTemplateSpecification: &LaunchTemplateSpecificationProperty{
//   			Version: jsii.String("version"),
//
//   			// the properties below are optional
//   			LaunchTemplateId: jsii.String("launchTemplateId"),
//   			LaunchTemplateName: jsii.String("launchTemplateName"),
//   		},
//
//   		// the properties below are optional
//   		Overrides: []interface{}{
//   			&LaunchTemplateOverridesProperty{
//   				InstanceRequirements: &InstanceRequirementsProperty{
//   					MemoryMiB: &MemoryMiBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					VCpuCount: &VCpuCountRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					AcceleratorCount: &AcceleratorCountRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					AcceleratorManufacturers: []*string{
//   						jsii.String("acceleratorManufacturers"),
//   					},
//   					AcceleratorNames: []*string{
//   						jsii.String("acceleratorNames"),
//   					},
//   					AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					AcceleratorTypes: []*string{
//   						jsii.String("acceleratorTypes"),
//   					},
//   					AllowedInstanceTypes: []*string{
//   						jsii.String("allowedInstanceTypes"),
//   					},
//   					BareMetal: jsii.String("bareMetal"),
//   					BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					BurstablePerformance: jsii.String("burstablePerformance"),
//   					CpuManufacturers: []*string{
//   						jsii.String("cpuManufacturers"),
//   					},
//   					ExcludedInstanceTypes: []*string{
//   						jsii.String("excludedInstanceTypes"),
//   					},
//   					InstanceGenerations: []*string{
//   						jsii.String("instanceGenerations"),
//   					},
//   					LocalStorage: jsii.String("localStorage"),
//   					LocalStorageTypes: []*string{
//   						jsii.String("localStorageTypes"),
//   					},
//   					MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   					MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   					OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					RequireHibernateSupport: jsii.Boolean(false),
//   					SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   						Max: jsii.Number(123),
//   						Min: jsii.Number(123),
//   					},
//   				},
//   				InstanceType: jsii.String("instanceType"),
//   				LaunchTemplateSpecification: &LaunchTemplateSpecificationProperty{
//   					Version: jsii.String("version"),
//
//   					// the properties below are optional
//   					LaunchTemplateId: jsii.String("launchTemplateId"),
//   					LaunchTemplateName: jsii.String("launchTemplateName"),
//   				},
//   				WeightedCapacity: jsii.String("weightedCapacity"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	InstancesDistribution: &InstancesDistributionProperty{
//   		OnDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   		OnDemandBaseCapacity: jsii.Number(123),
//   		OnDemandPercentageAboveBaseCapacity: jsii.Number(123),
//   		SpotAllocationStrategy: jsii.String("spotAllocationStrategy"),
//   		SpotInstancePools: jsii.Number(123),
//   		SpotMaxPrice: jsii.String("spotMaxPrice"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-mixedinstancespolicy.html
//
type CfnAutoScalingGroup_MixedInstancesPolicyProperty struct {
	// One or more launch templates and the instance types (overrides) that are used to launch EC2 instances to fulfill On-Demand and Spot capacities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-mixedinstancespolicy.html#cfn-autoscaling-autoscalinggroup-mixedinstancespolicy-launchtemplate
	//
	LaunchTemplate interface{} `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// The instances distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-mixedinstancespolicy.html#cfn-autoscaling-autoscalinggroup-mixedinstancespolicy-instancesdistribution
	//
	InstancesDistribution interface{} `field:"optional" json:"instancesDistribution" yaml:"instancesDistribution"`
}


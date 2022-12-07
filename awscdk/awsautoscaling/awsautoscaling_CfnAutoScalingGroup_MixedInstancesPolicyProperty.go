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
// `MixedInstancesPolicy` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mixedInstancesPolicyProperty := &mixedInstancesPolicyProperty{
//   	launchTemplate: &launchTemplateProperty{
//   		launchTemplateSpecification: &launchTemplateSpecificationProperty{
//   			version: jsii.String("version"),
//
//   			// the properties below are optional
//   			launchTemplateId: jsii.String("launchTemplateId"),
//   			launchTemplateName: jsii.String("launchTemplateName"),
//   		},
//
//   		// the properties below are optional
//   		overrides: []interface{}{
//   			&launchTemplateOverridesProperty{
//   				instanceRequirements: &instanceRequirementsProperty{
//   					acceleratorCount: &acceleratorCountRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					acceleratorManufacturers: []*string{
//   						jsii.String("acceleratorManufacturers"),
//   					},
//   					acceleratorNames: []*string{
//   						jsii.String("acceleratorNames"),
//   					},
//   					acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					acceleratorTypes: []*string{
//   						jsii.String("acceleratorTypes"),
//   					},
//   					allowedInstanceTypes: []*string{
//   						jsii.String("allowedInstanceTypes"),
//   					},
//   					bareMetal: jsii.String("bareMetal"),
//   					baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					burstablePerformance: jsii.String("burstablePerformance"),
//   					cpuManufacturers: []*string{
//   						jsii.String("cpuManufacturers"),
//   					},
//   					excludedInstanceTypes: []*string{
//   						jsii.String("excludedInstanceTypes"),
//   					},
//   					instanceGenerations: []*string{
//   						jsii.String("instanceGenerations"),
//   					},
//   					localStorage: jsii.String("localStorage"),
//   					localStorageTypes: []*string{
//   						jsii.String("localStorageTypes"),
//   					},
//   					memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					memoryMiB: &memoryMiBRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					networkBandwidthGbps: &networkBandwidthGbpsRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					requireHibernateSupport: jsii.Boolean(false),
//   					spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   					totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   					vCpuCount: &vCpuCountRequestProperty{
//   						max: jsii.Number(123),
//   						min: jsii.Number(123),
//   					},
//   				},
//   				instanceType: jsii.String("instanceType"),
//   				launchTemplateSpecification: &launchTemplateSpecificationProperty{
//   					version: jsii.String("version"),
//
//   					// the properties below are optional
//   					launchTemplateId: jsii.String("launchTemplateId"),
//   					launchTemplateName: jsii.String("launchTemplateName"),
//   				},
//   				weightedCapacity: jsii.String("weightedCapacity"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	instancesDistribution: &instancesDistributionProperty{
//   		onDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   		onDemandBaseCapacity: jsii.Number(123),
//   		onDemandPercentageAboveBaseCapacity: jsii.Number(123),
//   		spotAllocationStrategy: jsii.String("spotAllocationStrategy"),
//   		spotInstancePools: jsii.Number(123),
//   		spotMaxPrice: jsii.String("spotMaxPrice"),
//   	},
//   }
//
type CfnAutoScalingGroup_MixedInstancesPolicyProperty struct {
	// One or more launch templates and the instance types (overrides) that are used to launch EC2 instances to fulfill On-Demand and Spot capacities.
	LaunchTemplate interface{} `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// The instances distribution.
	InstancesDistribution interface{} `field:"optional" json:"instancesDistribution" yaml:"instancesDistribution"`
}


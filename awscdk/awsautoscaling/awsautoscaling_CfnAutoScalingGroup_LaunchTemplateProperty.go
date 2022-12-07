package awsautoscaling


// `LaunchTemplate` is a property of the [AWS::AutoScaling::AutoScalingGroup MixedInstancesPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-mixedinstancespolicy.html) property type that describes a launch template and overrides. The overrides are used to override the instance type specified by the launch template with multiple instance types that can be used to launch On-Demand Instances and Spot Instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateProperty := &launchTemplateProperty{
//   	launchTemplateSpecification: &launchTemplateSpecificationProperty{
//   		version: jsii.String("version"),
//
//   		// the properties below are optional
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   	},
//
//   	// the properties below are optional
//   	overrides: []interface{}{
//   		&launchTemplateOverridesProperty{
//   			instanceRequirements: &instanceRequirementsProperty{
//   				acceleratorCount: &acceleratorCountRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				acceleratorManufacturers: []*string{
//   					jsii.String("acceleratorManufacturers"),
//   				},
//   				acceleratorNames: []*string{
//   					jsii.String("acceleratorNames"),
//   				},
//   				acceleratorTotalMemoryMiB: &acceleratorTotalMemoryMiBRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				acceleratorTypes: []*string{
//   					jsii.String("acceleratorTypes"),
//   				},
//   				allowedInstanceTypes: []*string{
//   					jsii.String("allowedInstanceTypes"),
//   				},
//   				bareMetal: jsii.String("bareMetal"),
//   				baselineEbsBandwidthMbps: &baselineEbsBandwidthMbpsRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				burstablePerformance: jsii.String("burstablePerformance"),
//   				cpuManufacturers: []*string{
//   					jsii.String("cpuManufacturers"),
//   				},
//   				excludedInstanceTypes: []*string{
//   					jsii.String("excludedInstanceTypes"),
//   				},
//   				instanceGenerations: []*string{
//   					jsii.String("instanceGenerations"),
//   				},
//   				localStorage: jsii.String("localStorage"),
//   				localStorageTypes: []*string{
//   					jsii.String("localStorageTypes"),
//   				},
//   				memoryGiBPerVCpu: &memoryGiBPerVCpuRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				memoryMiB: &memoryMiBRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				networkBandwidthGbps: &networkBandwidthGbpsRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				networkInterfaceCount: &networkInterfaceCountRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				onDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				requireHibernateSupport: jsii.Boolean(false),
//   				spotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				totalLocalStorageGb: &totalLocalStorageGBRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   				vCpuCount: &vCpuCountRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   			},
//   			instanceType: jsii.String("instanceType"),
//   			launchTemplateSpecification: &launchTemplateSpecificationProperty{
//   				version: jsii.String("version"),
//
//   				// the properties below are optional
//   				launchTemplateId: jsii.String("launchTemplateId"),
//   				launchTemplateName: jsii.String("launchTemplateName"),
//   			},
//   			weightedCapacity: jsii.String("weightedCapacity"),
//   		},
//   	},
//   }
//
type CfnAutoScalingGroup_LaunchTemplateProperty struct {
	// The launch template to use.
	LaunchTemplateSpecification interface{} `field:"required" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Any properties that you specify override the same properties in the launch template.
	//
	// If not provided, Amazon EC2 Auto Scaling uses the instance type or instance type requirements specified in the launch template when it launches an instance.
	//
	// The overrides can include either one or more instance types or a set of instance requirements, but not both.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}


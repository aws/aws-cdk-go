package awsec2


// Specifies a launch template and overrides.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateConfigProperty := &launchTemplateConfigProperty{
//   	launchTemplateSpecification: &fleetLaunchTemplateSpecificationProperty{
//   		version: jsii.String("version"),
//
//   		// the properties below are optional
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	overrides: []interface{}{
//   		&launchTemplateOverridesProperty{
//   			availabilityZone: jsii.String("availabilityZone"),
//   			instanceRequirements: &instanceRequirementsRequestProperty{
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
//   				vCpuCount: &vCpuCountRangeRequestProperty{
//   					max: jsii.Number(123),
//   					min: jsii.Number(123),
//   				},
//   			},
//   			instanceType: jsii.String("instanceType"),
//   			priority: jsii.Number(123),
//   			spotPrice: jsii.String("spotPrice"),
//   			subnetId: jsii.String("subnetId"),
//   			weightedCapacity: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnSpotFleet_LaunchTemplateConfigProperty struct {
	// The launch template.
	LaunchTemplateSpecification interface{} `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Any parameters that you specify override the same parameters in the launch template.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}


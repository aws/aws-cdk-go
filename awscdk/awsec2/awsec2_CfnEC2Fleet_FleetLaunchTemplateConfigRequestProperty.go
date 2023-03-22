package awsec2


// Specifies a launch template and overrides for an EC2 Fleet.
//
// `FleetLaunchTemplateConfigRequest` is a property of the [AWS::EC2::EC2Fleet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetLaunchTemplateConfigRequestProperty := &fleetLaunchTemplateConfigRequestProperty{
//   	launchTemplateSpecification: &fleetLaunchTemplateSpecificationRequestProperty{
//   		version: jsii.String("version"),
//
//   		// the properties below are optional
//   		launchTemplateId: jsii.String("launchTemplateId"),
//   		launchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	overrides: []interface{}{
//   		&fleetLaunchTemplateOverridesRequestProperty{
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
//   			maxPrice: jsii.String("maxPrice"),
//   			placement: &placementProperty{
//   				affinity: jsii.String("affinity"),
//   				availabilityZone: jsii.String("availabilityZone"),
//   				groupName: jsii.String("groupName"),
//   				hostId: jsii.String("hostId"),
//   				hostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   				partitionNumber: jsii.Number(123),
//   				spreadDomain: jsii.String("spreadDomain"),
//   				tenancy: jsii.String("tenancy"),
//   			},
//   			priority: jsii.Number(123),
//   			subnetId: jsii.String("subnetId"),
//   			weightedCapacity: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnEC2Fleet_FleetLaunchTemplateConfigRequestProperty struct {
	// The launch template to use.
	//
	// You must specify either the launch template ID or launch template name in the request.
	LaunchTemplateSpecification interface{} `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Any parameters that you specify override the same parameters in the launch template.
	//
	// For fleets of type `request` and `maintain` , a maximum of 300 items is allowed across all launch templates.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}


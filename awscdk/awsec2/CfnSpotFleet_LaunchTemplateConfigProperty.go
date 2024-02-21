package awsec2


// Specifies a launch template and overrides.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateConfigProperty := &LaunchTemplateConfigProperty{
//   	LaunchTemplateSpecification: &FleetLaunchTemplateSpecificationProperty{
//   		Version: jsii.String("version"),
//
//   		// the properties below are optional
//   		LaunchTemplateId: jsii.String("launchTemplateId"),
//   		LaunchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	Overrides: []interface{}{
//   		&LaunchTemplateOverridesProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			InstanceRequirements: &InstanceRequirementsRequestProperty{
//   				AcceleratorCount: &AcceleratorCountRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorManufacturers: []*string{
//   					jsii.String("acceleratorManufacturers"),
//   				},
//   				AcceleratorNames: []*string{
//   					jsii.String("acceleratorNames"),
//   				},
//   				AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				AcceleratorTypes: []*string{
//   					jsii.String("acceleratorTypes"),
//   				},
//   				AllowedInstanceTypes: []*string{
//   					jsii.String("allowedInstanceTypes"),
//   				},
//   				BareMetal: jsii.String("bareMetal"),
//   				BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				BurstablePerformance: jsii.String("burstablePerformance"),
//   				CpuManufacturers: []*string{
//   					jsii.String("cpuManufacturers"),
//   				},
//   				ExcludedInstanceTypes: []*string{
//   					jsii.String("excludedInstanceTypes"),
//   				},
//   				InstanceGenerations: []*string{
//   					jsii.String("instanceGenerations"),
//   				},
//   				LocalStorage: jsii.String("localStorage"),
//   				LocalStorageTypes: []*string{
//   					jsii.String("localStorageTypes"),
//   				},
//   				MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   				MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				MemoryMiB: &MemoryMiBRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				RequireHibernateSupport: jsii.Boolean(false),
//   				SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   				TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				VCpuCount: &VCpuCountRangeRequestProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   			},
//   			InstanceType: jsii.String("instanceType"),
//   			Priority: jsii.Number(123),
//   			SpotPrice: jsii.String("spotPrice"),
//   			SubnetId: jsii.String("subnetId"),
//   			WeightedCapacity: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-launchtemplateconfig.html
//
type CfnSpotFleet_LaunchTemplateConfigProperty struct {
	// The launch template to use.
	//
	// Make sure that the launch template does not contain the `NetworkInterfaceId` parameter because you can't specify a network interface ID in a Spot Fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-launchtemplateconfig.html#cfn-ec2-spotfleet-launchtemplateconfig-launchtemplatespecification
	//
	LaunchTemplateSpecification interface{} `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Any parameters that you specify override the same parameters in the launch template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-launchtemplateconfig.html#cfn-ec2-spotfleet-launchtemplateconfig-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}


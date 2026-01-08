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
//   fleetLaunchTemplateConfigRequestProperty := &FleetLaunchTemplateConfigRequestProperty{
//   	LaunchTemplateSpecification: &FleetLaunchTemplateSpecificationRequestProperty{
//   		Version: jsii.String("version"),
//
//   		// the properties below are optional
//   		LaunchTemplateId: jsii.String("launchTemplateId"),
//   		LaunchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	Overrides: []interface{}{
//   		&FleetLaunchTemplateOverridesRequestProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   			BlockDeviceMappings: []interface{}{
//   				&BlockDeviceMappingProperty{
//   					DeviceName: jsii.String("deviceName"),
//   					Ebs: &EbsBlockDeviceProperty{
//   						DeleteOnTermination: jsii.Boolean(false),
//   						Encrypted: jsii.Boolean(false),
//   						Iops: jsii.Number(123),
//   						KmsKeyId: jsii.String("kmsKeyId"),
//   						SnapshotId: jsii.String("snapshotId"),
//   						VolumeSize: jsii.Number(123),
//   						VolumeType: jsii.String("volumeType"),
//   					},
//   					NoDevice: jsii.String("noDevice"),
//   					VirtualName: jsii.String("virtualName"),
//   				},
//   			},
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
//   				BaselinePerformanceFactors: &BaselinePerformanceFactorsRequestProperty{
//   					Cpu: &CpuPerformanceFactorRequestProperty{
//   						References: []interface{}{
//   							&PerformanceFactorReferenceRequestProperty{
//   								InstanceFamily: jsii.String("instanceFamily"),
//   							},
//   						},
//   					},
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
//   				RequireEncryptionInTransit: jsii.Boolean(false),
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
//   			MaxPrice: jsii.String("maxPrice"),
//   			Placement: &PlacementProperty{
//   				Affinity: jsii.String("affinity"),
//   				AvailabilityZone: jsii.String("availabilityZone"),
//   				GroupName: jsii.String("groupName"),
//   				HostId: jsii.String("hostId"),
//   				HostResourceGroupArn: jsii.String("hostResourceGroupArn"),
//   				PartitionNumber: jsii.Number(123),
//   				SpreadDomain: jsii.String("spreadDomain"),
//   				Tenancy: jsii.String("tenancy"),
//   			},
//   			Priority: jsii.Number(123),
//   			SubnetId: jsii.String("subnetId"),
//   			WeightedCapacity: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplateconfigrequest.html
//
type CfnEC2Fleet_FleetLaunchTemplateConfigRequestProperty struct {
	// The launch template to use.
	//
	// You must specify either the launch template ID or launch template name in the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplateconfigrequest.html#cfn-ec2-ec2fleet-fleetlaunchtemplateconfigrequest-launchtemplatespecification
	//
	LaunchTemplateSpecification interface{} `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Any parameters that you specify override the same parameters in the launch template.
	//
	// For fleets of type `request` and `maintain` , a maximum of 300 items is allowed across all launch templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplateconfigrequest.html#cfn-ec2-ec2fleet-fleetlaunchtemplateconfigrequest-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}


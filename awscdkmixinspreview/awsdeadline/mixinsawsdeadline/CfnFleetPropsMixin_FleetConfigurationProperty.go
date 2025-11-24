package mixinsawsdeadline


// Fleet configuration details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fleetConfigurationProperty := &FleetConfigurationProperty{
//   	CustomerManaged: &CustomerManagedFleetConfigurationProperty{
//   		Mode: jsii.String("mode"),
//   		StorageProfileId: jsii.String("storageProfileId"),
//   		TagPropagationMode: jsii.String("tagPropagationMode"),
//   		WorkerCapabilities: &CustomerManagedWorkerCapabilitiesProperty{
//   			AcceleratorCount: &AcceleratorCountRangeProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRangeProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			AcceleratorTypes: []*string{
//   				jsii.String("acceleratorTypes"),
//   			},
//   			CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   			CustomAmounts: []interface{}{
//   				&FleetAmountCapabilityProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			CustomAttributes: []interface{}{
//   				&FleetAttributeCapabilityProperty{
//   					Name: jsii.String("name"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			MemoryMiB: &MemoryMiBRangeProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			OsFamily: jsii.String("osFamily"),
//   			VCpuCount: &VCpuCountRangeProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ServiceManagedEc2: &ServiceManagedEc2FleetConfigurationProperty{
//   		InstanceCapabilities: &ServiceManagedEc2InstanceCapabilitiesProperty{
//   			AcceleratorCapabilities: &AcceleratorCapabilitiesProperty{
//   				Count: &AcceleratorCountRangeProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   				},
//   				Selections: []interface{}{
//   					&AcceleratorSelectionProperty{
//   						Name: jsii.String("name"),
//   						Runtime: jsii.String("runtime"),
//   					},
//   				},
//   			},
//   			AllowedInstanceTypes: []*string{
//   				jsii.String("allowedInstanceTypes"),
//   			},
//   			CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   			CustomAmounts: []interface{}{
//   				&FleetAmountCapabilityProperty{
//   					Max: jsii.Number(123),
//   					Min: jsii.Number(123),
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			CustomAttributes: []interface{}{
//   				&FleetAttributeCapabilityProperty{
//   					Name: jsii.String("name"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			ExcludedInstanceTypes: []*string{
//   				jsii.String("excludedInstanceTypes"),
//   			},
//   			MemoryMiB: &MemoryMiBRangeProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			OsFamily: jsii.String("osFamily"),
//   			RootEbsVolume: &Ec2EbsVolumeProperty{
//   				Iops: jsii.Number(123),
//   				SizeGiB: jsii.Number(123),
//   				ThroughputMiB: jsii.Number(123),
//   			},
//   			VCpuCount: &VCpuCountRangeProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   		},
//   		InstanceMarketOptions: &ServiceManagedEc2InstanceMarketOptionsProperty{
//   			Type: jsii.String("type"),
//   		},
//   		StorageProfileId: jsii.String("storageProfileId"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			ResourceConfigurationArns: []*string{
//   				jsii.String("resourceConfigurationArns"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetconfiguration.html
//
type CfnFleetPropsMixin_FleetConfigurationProperty struct {
	// The customer managed fleets within a fleet configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetconfiguration.html#cfn-deadline-fleet-fleetconfiguration-customermanaged
	//
	CustomerManaged interface{} `field:"optional" json:"customerManaged" yaml:"customerManaged"`
	// The service managed Amazon EC2 instances for a fleet configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetconfiguration.html#cfn-deadline-fleet-fleetconfiguration-servicemanagedec2
	//
	ServiceManagedEc2 interface{} `field:"optional" json:"serviceManagedEc2" yaml:"serviceManagedEc2"`
}


package awsdeadline


// Fleet configuration details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fleetConfigurationProperty := &FleetConfigurationProperty{
//   	CustomerManaged: &CustomerManagedFleetConfigurationProperty{
//   		Mode: jsii.String("mode"),
//   		WorkerCapabilities: &CustomerManagedWorkerCapabilitiesProperty{
//   			CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   			MemoryMiB: &MemoryMiBRangeProperty{
//   				Min: jsii.Number(123),
//
//   				// the properties below are optional
//   				Max: jsii.Number(123),
//   			},
//   			OsFamily: jsii.String("osFamily"),
//   			VCpuCount: &VCpuCountRangeProperty{
//   				Min: jsii.Number(123),
//
//   				// the properties below are optional
//   				Max: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			AcceleratorCount: &AcceleratorCountRangeProperty{
//   				Min: jsii.Number(123),
//
//   				// the properties below are optional
//   				Max: jsii.Number(123),
//   			},
//   			AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRangeProperty{
//   				Min: jsii.Number(123),
//
//   				// the properties below are optional
//   				Max: jsii.Number(123),
//   			},
//   			AcceleratorTypes: []*string{
//   				jsii.String("acceleratorTypes"),
//   			},
//   			CustomAmounts: []interface{}{
//   				&FleetAmountCapabilityProperty{
//   					Min: jsii.Number(123),
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
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
//   		},
//
//   		// the properties below are optional
//   		StorageProfileId: jsii.String("storageProfileId"),
//   		TagPropagationMode: jsii.String("tagPropagationMode"),
//   	},
//   	ServiceManagedEc2: &ServiceManagedEc2FleetConfigurationProperty{
//   		InstanceCapabilities: &ServiceManagedEc2InstanceCapabilitiesProperty{
//   			CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   			MemoryMiB: &MemoryMiBRangeProperty{
//   				Min: jsii.Number(123),
//
//   				// the properties below are optional
//   				Max: jsii.Number(123),
//   			},
//   			OsFamily: jsii.String("osFamily"),
//   			VCpuCount: &VCpuCountRangeProperty{
//   				Min: jsii.Number(123),
//
//   				// the properties below are optional
//   				Max: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			AcceleratorCapabilities: &AcceleratorCapabilitiesProperty{
//   				Selections: []interface{}{
//   					&AcceleratorSelectionProperty{
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Runtime: jsii.String("runtime"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Count: &AcceleratorCountRangeProperty{
//   					Min: jsii.Number(123),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
//   				},
//   			},
//   			AllowedInstanceTypes: []*string{
//   				jsii.String("allowedInstanceTypes"),
//   			},
//   			CustomAmounts: []interface{}{
//   				&FleetAmountCapabilityProperty{
//   					Min: jsii.Number(123),
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Max: jsii.Number(123),
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
//   			RootEbsVolume: &Ec2EbsVolumeProperty{
//   				Iops: jsii.Number(123),
//   				SizeGiB: jsii.Number(123),
//   				ThroughputMiB: jsii.Number(123),
//   			},
//   		},
//   		InstanceMarketOptions: &ServiceManagedEc2InstanceMarketOptionsProperty{
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetconfiguration.html
//
type CfnFleet_FleetConfigurationProperty struct {
	// The customer managed fleets within a fleet configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetconfiguration.html#cfn-deadline-fleet-fleetconfiguration-customermanaged
	//
	CustomerManaged interface{} `field:"optional" json:"customerManaged" yaml:"customerManaged"`
	// The service managed Amazon EC2 instances for a fleet configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-fleetconfiguration.html#cfn-deadline-fleet-fleetconfiguration-servicemanagedec2
	//
	ServiceManagedEc2 interface{} `field:"optional" json:"serviceManagedEc2" yaml:"serviceManagedEc2"`
}


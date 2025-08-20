package awsdeadline


// The configuration details for a service managed Amazon EC2 fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceManagedEc2FleetConfigurationProperty := &ServiceManagedEc2FleetConfigurationProperty{
//   	InstanceCapabilities: &ServiceManagedEc2InstanceCapabilitiesProperty{
//   		CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   		MemoryMiB: &MemoryMiBRangeProperty{
//   			Min: jsii.Number(123),
//
//   			// the properties below are optional
//   			Max: jsii.Number(123),
//   		},
//   		OsFamily: jsii.String("osFamily"),
//   		VCpuCount: &VCpuCountRangeProperty{
//   			Min: jsii.Number(123),
//
//   			// the properties below are optional
//   			Max: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		AcceleratorCapabilities: &AcceleratorCapabilitiesProperty{
//   			Selections: []interface{}{
//   				&AcceleratorSelectionProperty{
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Runtime: jsii.String("runtime"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Count: &AcceleratorCountRangeProperty{
//   				Min: jsii.Number(123),
//
//   				// the properties below are optional
//   				Max: jsii.Number(123),
//   			},
//   		},
//   		AllowedInstanceTypes: []*string{
//   			jsii.String("allowedInstanceTypes"),
//   		},
//   		CustomAmounts: []interface{}{
//   			&FleetAmountCapabilityProperty{
//   				Min: jsii.Number(123),
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Max: jsii.Number(123),
//   			},
//   		},
//   		CustomAttributes: []interface{}{
//   			&FleetAttributeCapabilityProperty{
//   				Name: jsii.String("name"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		ExcludedInstanceTypes: []*string{
//   			jsii.String("excludedInstanceTypes"),
//   		},
//   		RootEbsVolume: &Ec2EbsVolumeProperty{
//   			Iops: jsii.Number(123),
//   			SizeGiB: jsii.Number(123),
//   			ThroughputMiB: jsii.Number(123),
//   		},
//   	},
//   	InstanceMarketOptions: &ServiceManagedEc2InstanceMarketOptionsProperty{
//   		Type: jsii.String("type"),
//   	},
//
//   	// the properties below are optional
//   	StorageProfileId: jsii.String("storageProfileId"),
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		ResourceConfigurationArns: []*string{
//   			jsii.String("resourceConfigurationArns"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.html
//
type CfnFleet_ServiceManagedEc2FleetConfigurationProperty struct {
	// The Amazon EC2 instance capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.html#cfn-deadline-fleet-servicemanagedec2fleetconfiguration-instancecapabilities
	//
	InstanceCapabilities interface{} `field:"required" json:"instanceCapabilities" yaml:"instanceCapabilities"`
	// The Amazon EC2 market type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.html#cfn-deadline-fleet-servicemanagedec2fleetconfiguration-instancemarketoptions
	//
	InstanceMarketOptions interface{} `field:"required" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// The storage profile ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.html#cfn-deadline-fleet-servicemanagedec2fleetconfiguration-storageprofileid
	//
	StorageProfileId *string `field:"optional" json:"storageProfileId" yaml:"storageProfileId"`
	// The VPC configuration details for a service managed Amazon EC2 fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.html#cfn-deadline-fleet-servicemanagedec2fleetconfiguration-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}


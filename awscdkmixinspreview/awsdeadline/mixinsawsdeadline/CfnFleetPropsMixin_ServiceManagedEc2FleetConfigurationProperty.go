package mixinsawsdeadline


// The configuration details for a service managed Amazon EC2 fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceManagedEc2FleetConfigurationProperty := &ServiceManagedEc2FleetConfigurationProperty{
//   	InstanceCapabilities: &ServiceManagedEc2InstanceCapabilitiesProperty{
//   		AcceleratorCapabilities: &AcceleratorCapabilitiesProperty{
//   			Count: &AcceleratorCountRangeProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   			},
//   			Selections: []interface{}{
//   				&AcceleratorSelectionProperty{
//   					Name: jsii.String("name"),
//   					Runtime: jsii.String("runtime"),
//   				},
//   			},
//   		},
//   		AllowedInstanceTypes: []*string{
//   			jsii.String("allowedInstanceTypes"),
//   		},
//   		CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   		CustomAmounts: []interface{}{
//   			&FleetAmountCapabilityProperty{
//   				Max: jsii.Number(123),
//   				Min: jsii.Number(123),
//   				Name: jsii.String("name"),
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
//   		MemoryMiB: &MemoryMiBRangeProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		OsFamily: jsii.String("osFamily"),
//   		RootEbsVolume: &Ec2EbsVolumeProperty{
//   			Iops: jsii.Number(123),
//   			SizeGiB: jsii.Number(123),
//   			ThroughputMiB: jsii.Number(123),
//   		},
//   		VCpuCount: &VCpuCountRangeProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   	},
//   	InstanceMarketOptions: &ServiceManagedEc2InstanceMarketOptionsProperty{
//   		Type: jsii.String("type"),
//   	},
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
type CfnFleetPropsMixin_ServiceManagedEc2FleetConfigurationProperty struct {
	// The Amazon EC2 instance capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.html#cfn-deadline-fleet-servicemanagedec2fleetconfiguration-instancecapabilities
	//
	InstanceCapabilities interface{} `field:"optional" json:"instanceCapabilities" yaml:"instanceCapabilities"`
	// The Amazon EC2 market type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.html#cfn-deadline-fleet-servicemanagedec2fleetconfiguration-instancemarketoptions
	//
	InstanceMarketOptions interface{} `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// The storage profile ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.html#cfn-deadline-fleet-servicemanagedec2fleetconfiguration-storageprofileid
	//
	StorageProfileId *string `field:"optional" json:"storageProfileId" yaml:"storageProfileId"`
	// The VPC configuration details for a service managed Amazon EC2 fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2fleetconfiguration.html#cfn-deadline-fleet-servicemanagedec2fleetconfiguration-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}


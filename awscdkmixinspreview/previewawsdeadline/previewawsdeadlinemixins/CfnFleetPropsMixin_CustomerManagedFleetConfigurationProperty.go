package previewawsdeadlinemixins


// The configuration details for a customer managed fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customerManagedFleetConfigurationProperty := &CustomerManagedFleetConfigurationProperty{
//   	Mode: jsii.String("mode"),
//   	StorageProfileId: jsii.String("storageProfileId"),
//   	TagPropagationMode: jsii.String("tagPropagationMode"),
//   	WorkerCapabilities: &CustomerManagedWorkerCapabilitiesProperty{
//   		AcceleratorCount: &AcceleratorCountRangeProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRangeProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		AcceleratorTypes: []*string{
//   			jsii.String("acceleratorTypes"),
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
//   		MemoryMiB: &MemoryMiBRangeProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   		OsFamily: jsii.String("osFamily"),
//   		VCpuCount: &VCpuCountRangeProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html
//
type CfnFleetPropsMixin_CustomerManagedFleetConfigurationProperty struct {
	// The Auto Scaling mode for the customer managed fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The storage profile ID for the customer managed fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-storageprofileid
	//
	StorageProfileId *string `field:"optional" json:"storageProfileId" yaml:"storageProfileId"`
	// The tag propagation mode for the customer managed fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-tagpropagationmode
	//
	TagPropagationMode *string `field:"optional" json:"tagPropagationMode" yaml:"tagPropagationMode"`
	// The worker capabilities for the customer managed fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-workercapabilities
	//
	WorkerCapabilities interface{} `field:"optional" json:"workerCapabilities" yaml:"workerCapabilities"`
}


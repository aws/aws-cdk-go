package awsdeadline


// The details of a customer managed fleet configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedFleetConfigurationProperty := &CustomerManagedFleetConfigurationProperty{
//   	Mode: jsii.String("mode"),
//   	WorkerCapabilities: &CustomerManagedWorkerCapabilitiesProperty{
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
//   		AcceleratorCount: &AcceleratorCountRangeProperty{
//   			Min: jsii.Number(123),
//
//   			// the properties below are optional
//   			Max: jsii.Number(123),
//   		},
//   		AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRangeProperty{
//   			Min: jsii.Number(123),
//
//   			// the properties below are optional
//   			Max: jsii.Number(123),
//   		},
//   		AcceleratorTypes: []*string{
//   			jsii.String("acceleratorTypes"),
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
//   	},
//
//   	// the properties below are optional
//   	StorageProfileId: jsii.String("storageProfileId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html
//
type CfnFleet_CustomerManagedFleetConfigurationProperty struct {
	// The AWS Auto Scaling mode for the customer managed fleet configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-mode
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The worker capabilities for a customer managed fleet configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-workercapabilities
	//
	WorkerCapabilities interface{} `field:"required" json:"workerCapabilities" yaml:"workerCapabilities"`
	// The storage profile ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-storageprofileid
	//
	StorageProfileId *string `field:"optional" json:"storageProfileId" yaml:"storageProfileId"`
}


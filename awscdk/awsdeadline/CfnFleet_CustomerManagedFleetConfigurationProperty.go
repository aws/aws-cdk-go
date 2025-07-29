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
//   	TagPropagationMode: jsii.String("tagPropagationMode"),
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
	// Specifies whether tags associated with a fleet are attached to workers when the worker is launched.
	//
	// When the `tagPropagationMode` is set to `PROPAGATE_TAGS_TO_WORKERS_AT_LAUNCH` any tag associated with a fleet is attached to workers when they launch. If the tags for a fleet change, the tags associated with running workers *do not* change.
	//
	// If you don't specify `tagPropagationMode` , the default is `NO_PROPAGATION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedfleetconfiguration.html#cfn-deadline-fleet-customermanagedfleetconfiguration-tagpropagationmode
	//
	TagPropagationMode *string `field:"optional" json:"tagPropagationMode" yaml:"tagPropagationMode"`
}


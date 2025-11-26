package previewawsdeadlinemixins


// The worker capabilities for a customer managed workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customerManagedWorkerCapabilitiesProperty := &CustomerManagedWorkerCapabilitiesProperty{
//   	AcceleratorCount: &AcceleratorCountRangeProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRangeProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	AcceleratorTypes: []*string{
//   		jsii.String("acceleratorTypes"),
//   	},
//   	CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   	CustomAmounts: []interface{}{
//   		&FleetAmountCapabilityProperty{
//   			Max: jsii.Number(123),
//   			Min: jsii.Number(123),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	CustomAttributes: []interface{}{
//   		&FleetAttributeCapabilityProperty{
//   			Name: jsii.String("name"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	MemoryMiB: &MemoryMiBRangeProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	OsFamily: jsii.String("osFamily"),
//   	VCpuCount: &VCpuCountRangeProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html
//
type CfnFleetPropsMixin_CustomerManagedWorkerCapabilitiesProperty struct {
	// The range of the accelerator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-acceleratorcount
	//
	AcceleratorCount interface{} `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// The total memory (MiB) for the customer managed worker capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-acceleratortotalmemorymib
	//
	AcceleratorTotalMemoryMiB interface{} `field:"optional" json:"acceleratorTotalMemoryMiB" yaml:"acceleratorTotalMemoryMiB"`
	// The accelerator types for the customer managed worker capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-acceleratortypes
	//
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// The CPU architecture type for the customer managed worker capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-cpuarchitecturetype
	//
	CpuArchitectureType *string `field:"optional" json:"cpuArchitectureType" yaml:"cpuArchitectureType"`
	// Custom requirement ranges for customer managed worker capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-customamounts
	//
	CustomAmounts interface{} `field:"optional" json:"customAmounts" yaml:"customAmounts"`
	// Custom attributes for the customer manged worker capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-customattributes
	//
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The memory (MiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-memorymib
	//
	MemoryMiB interface{} `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// The operating system (OS) family.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-osfamily
	//
	OsFamily *string `field:"optional" json:"osFamily" yaml:"osFamily"`
	// The vCPU count for the customer manged worker capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-vcpucount
	//
	VCpuCount interface{} `field:"optional" json:"vCpuCount" yaml:"vCpuCount"`
}


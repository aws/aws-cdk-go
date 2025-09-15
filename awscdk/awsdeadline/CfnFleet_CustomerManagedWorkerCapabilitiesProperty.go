package awsdeadline


// The worker capabilities for a customer managed workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedWorkerCapabilitiesProperty := &CustomerManagedWorkerCapabilitiesProperty{
//   	CpuArchitectureType: jsii.String("cpuArchitectureType"),
//   	MemoryMiB: &MemoryMiBRangeProperty{
//   		Min: jsii.Number(123),
//
//   		// the properties below are optional
//   		Max: jsii.Number(123),
//   	},
//   	OsFamily: jsii.String("osFamily"),
//   	VCpuCount: &VCpuCountRangeProperty{
//   		Min: jsii.Number(123),
//
//   		// the properties below are optional
//   		Max: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	AcceleratorCount: &AcceleratorCountRangeProperty{
//   		Min: jsii.Number(123),
//
//   		// the properties below are optional
//   		Max: jsii.Number(123),
//   	},
//   	AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRangeProperty{
//   		Min: jsii.Number(123),
//
//   		// the properties below are optional
//   		Max: jsii.Number(123),
//   	},
//   	AcceleratorTypes: []*string{
//   		jsii.String("acceleratorTypes"),
//   	},
//   	CustomAmounts: []interface{}{
//   		&FleetAmountCapabilityProperty{
//   			Min: jsii.Number(123),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Max: jsii.Number(123),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html
//
type CfnFleet_CustomerManagedWorkerCapabilitiesProperty struct {
	// The CPU architecture type for the customer managed worker capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-cpuarchitecturetype
	//
	CpuArchitectureType *string `field:"required" json:"cpuArchitectureType" yaml:"cpuArchitectureType"`
	// The memory (MiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-memorymib
	//
	MemoryMiB interface{} `field:"required" json:"memoryMiB" yaml:"memoryMiB"`
	// The operating system (OS) family.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-osfamily
	//
	OsFamily *string `field:"required" json:"osFamily" yaml:"osFamily"`
	// The vCPU count for the customer manged worker capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-vcpucount
	//
	VCpuCount interface{} `field:"required" json:"vCpuCount" yaml:"vCpuCount"`
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
	// Custom requirement ranges for customer managed worker capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-customamounts
	//
	CustomAmounts interface{} `field:"optional" json:"customAmounts" yaml:"customAmounts"`
	// Custom attributes for the customer manged worker capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-customermanagedworkercapabilities.html#cfn-deadline-fleet-customermanagedworkercapabilities-customattributes
	//
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
}


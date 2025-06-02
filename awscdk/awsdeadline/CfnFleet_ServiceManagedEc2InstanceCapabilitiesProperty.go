package awsdeadline


// The Amazon EC2 instance capabilities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceManagedEc2InstanceCapabilitiesProperty := &ServiceManagedEc2InstanceCapabilitiesProperty{
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
//   	AcceleratorCapabilities: &AcceleratorCapabilitiesProperty{
//   		Selections: []interface{}{
//   			&AcceleratorSelectionProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Runtime: jsii.String("runtime"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Count: &AcceleratorCountRangeProperty{
//   			Min: jsii.Number(123),
//
//   			// the properties below are optional
//   			Max: jsii.Number(123),
//   		},
//   	},
//   	AllowedInstanceTypes: []*string{
//   		jsii.String("allowedInstanceTypes"),
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
//   	ExcludedInstanceTypes: []*string{
//   		jsii.String("excludedInstanceTypes"),
//   	},
//   	RootEbsVolume: &Ec2EbsVolumeProperty{
//   		Iops: jsii.Number(123),
//   		SizeGiB: jsii.Number(123),
//   		ThroughputMiB: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html
//
type CfnFleet_ServiceManagedEc2InstanceCapabilitiesProperty struct {
	// The CPU architecture type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html#cfn-deadline-fleet-servicemanagedec2instancecapabilities-cpuarchitecturetype
	//
	CpuArchitectureType *string `field:"required" json:"cpuArchitectureType" yaml:"cpuArchitectureType"`
	// The memory, as MiB, for the Amazon EC2 instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html#cfn-deadline-fleet-servicemanagedec2instancecapabilities-memorymib
	//
	MemoryMiB interface{} `field:"required" json:"memoryMiB" yaml:"memoryMiB"`
	// The operating system (OS) family.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html#cfn-deadline-fleet-servicemanagedec2instancecapabilities-osfamily
	//
	OsFamily *string `field:"required" json:"osFamily" yaml:"osFamily"`
	// The amount of vCPU to require for instances in this fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html#cfn-deadline-fleet-servicemanagedec2instancecapabilities-vcpucount
	//
	VCpuCount interface{} `field:"required" json:"vCpuCount" yaml:"vCpuCount"`
	// Describes the GPU accelerator capabilities required for worker host instances in this fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html#cfn-deadline-fleet-servicemanagedec2instancecapabilities-acceleratorcapabilities
	//
	AcceleratorCapabilities interface{} `field:"optional" json:"acceleratorCapabilities" yaml:"acceleratorCapabilities"`
	// The allowable Amazon EC2 instance types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html#cfn-deadline-fleet-servicemanagedec2instancecapabilities-allowedinstancetypes
	//
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// The custom capability amounts to require for instances in this fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html#cfn-deadline-fleet-servicemanagedec2instancecapabilities-customamounts
	//
	CustomAmounts interface{} `field:"optional" json:"customAmounts" yaml:"customAmounts"`
	// The custom capability attributes to require for instances in this fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html#cfn-deadline-fleet-servicemanagedec2instancecapabilities-customattributes
	//
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The instance types to exclude from the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html#cfn-deadline-fleet-servicemanagedec2instancecapabilities-excludedinstancetypes
	//
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// The root EBS volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.html#cfn-deadline-fleet-servicemanagedec2instancecapabilities-rootebsvolume
	//
	RootEbsVolume interface{} `field:"optional" json:"rootEbsVolume" yaml:"rootEbsVolume"`
}


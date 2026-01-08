package previewawsdeadlinemixins


// Provides information about the GPU accelerators used for jobs processed by a fleet.
//
// > Accelerator capabilities cannot be used with wait-and-save fleets. If you specify accelerator capabilities, you must use either spot or on-demand instance market options. > Each accelerator type maps to specific EC2 instance families:
// >
// > - `t4` : Uses G4dn instance family
// > - `a10g` : Uses G5 instance family
// > - `l4` : Uses G6 and Gr6 instance families
// > - `l40s` : Uses G6e instance family.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   acceleratorCapabilitiesProperty := &AcceleratorCapabilitiesProperty{
//   	Count: &AcceleratorCountRangeProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	Selections: []interface{}{
//   		&AcceleratorSelectionProperty{
//   			Name: jsii.String("name"),
//   			Runtime: jsii.String("runtime"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcapabilities.html
//
type CfnFleetPropsMixin_AcceleratorCapabilitiesProperty struct {
	// The number of GPU accelerators specified for worker hosts in this fleet.
	//
	// > You must specify either `acceleratorCapabilities.count.max` or `allowedInstanceTypes` when using accelerator capabilities. If you don't specify a maximum count, AWS Deadline Cloud uses the instance types you specify in `allowedInstanceTypes` to determine the maximum number of accelerators.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcapabilities.html#cfn-deadline-fleet-acceleratorcapabilities-count
	//
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// A list of accelerator capabilities requested for this fleet.
	//
	// Only Amazon Elastic Compute Cloud instances that provide these capabilities will be used. For example, if you specify both L4 and T4 chips, AWS Deadline Cloud will use Amazon EC2 instances that have either the L4 or the T4 chip installed.
	//
	// > - You must specify at least one accelerator selection.
	// > - You cannot specify the same accelerator name multiple times in the selections list.
	// > - All accelerators in the selections must use the same runtime version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcapabilities.html#cfn-deadline-fleet-acceleratorcapabilities-selections
	//
	Selections interface{} `field:"optional" json:"selections" yaml:"selections"`
}


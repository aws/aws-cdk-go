package previewawsdeadlinemixins


// Provides information about the GPU accelerators used for jobs processed by a fleet.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcapabilities.html#cfn-deadline-fleet-acceleratorcapabilities-count
	//
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// A list of accelerator capabilities requested for this fleet.
	//
	// Only Amazon Elastic Compute Cloud instances that provide these capabilities will be used. For example, if you specify both L4 and T4 chips, AWS Deadline Cloud will use Amazon EC2 instances that have either the L4 or the T4 chip installed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcapabilities.html#cfn-deadline-fleet-acceleratorcapabilities-selections
	//
	Selections interface{} `field:"optional" json:"selections" yaml:"selections"`
}


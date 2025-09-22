package awsdeadline


// Provides information about the GPU accelerators used for jobs processed by a fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorCapabilitiesProperty := &AcceleratorCapabilitiesProperty{
//   	Selections: []interface{}{
//   		&AcceleratorSelectionProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Runtime: jsii.String("runtime"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Count: &AcceleratorCountRangeProperty{
//   		Min: jsii.Number(123),
//
//   		// the properties below are optional
//   		Max: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcapabilities.html
//
type CfnFleet_AcceleratorCapabilitiesProperty struct {
	// A list of accelerator capabilities requested for this fleet.
	//
	// Only Amazon Elastic Compute Cloud instances that provide these capabilities will be used. For example, if you specify both L4 and T4 chips, Deadline Cloud will use Amazon EC2 instances that have either the L4 or the T4 chip installed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcapabilities.html#cfn-deadline-fleet-acceleratorcapabilities-selections
	//
	Selections interface{} `field:"required" json:"selections" yaml:"selections"`
	// The number of GPU accelerators specified for worker hosts in this fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcapabilities.html#cfn-deadline-fleet-acceleratorcapabilities-count
	//
	Count interface{} `field:"optional" json:"count" yaml:"count"`
}


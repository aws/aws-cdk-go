package previewawsdeadlinemixins


// Defines the maximum and minimum number of GPU accelerators required for a worker instance..
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   acceleratorCountRangeProperty := &AcceleratorCountRangeProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcountrange.html
//
type CfnFleetPropsMixin_AcceleratorCountRangeProperty struct {
	// The maximum number of GPU accelerators in the worker host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcountrange.html#cfn-deadline-fleet-acceleratorcountrange-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of GPU accelerators in the worker host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratorcountrange.html#cfn-deadline-fleet-acceleratorcountrange-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}


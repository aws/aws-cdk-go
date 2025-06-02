package awsquicksight


// The logarithmic axis scale setup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   axisLogarithmicScaleProperty := &AxisLogarithmicScaleProperty{
//   	Base: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislogarithmicscale.html
//
type CfnDashboard_AxisLogarithmicScaleProperty struct {
	// The base setup of a logarithmic axis scale.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislogarithmicscale.html#cfn-quicksight-dashboard-axislogarithmicscale-base
	//
	Base *float64 `field:"optional" json:"base" yaml:"base"`
}


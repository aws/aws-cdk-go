package previewawsquicksightmixins


// The scale setup options for a numeric axis display.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   axisScaleProperty := &AxisScaleProperty{
//   	Linear: &AxisLinearScaleProperty{
//   		StepCount: jsii.Number(123),
//   		StepSize: jsii.Number(123),
//   	},
//   	Logarithmic: &AxisLogarithmicScaleProperty{
//   		Base: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisscale.html
//
type CfnDashboardPropsMixin_AxisScaleProperty struct {
	// The linear axis scale setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisscale.html#cfn-quicksight-dashboard-axisscale-linear
	//
	Linear interface{} `field:"optional" json:"linear" yaml:"linear"`
	// The logarithmic axis scale setup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axisscale.html#cfn-quicksight-dashboard-axisscale-logarithmic
	//
	Logarithmic interface{} `field:"optional" json:"logarithmic" yaml:"logarithmic"`
}


package previewawsquicksightmixins


// Determines the gradient stop configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gradientStopProperty := &GradientStopProperty{
//   	Color: jsii.String("color"),
//   	DataValue: jsii.Number(123),
//   	GradientOffset: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gradientstop.html
//
type CfnDashboardPropsMixin_GradientStopProperty struct {
	// Determines the color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gradientstop.html#cfn-quicksight-dashboard-gradientstop-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Determines the data value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gradientstop.html#cfn-quicksight-dashboard-gradientstop-datavalue
	//
	DataValue *float64 `field:"optional" json:"dataValue" yaml:"dataValue"`
	// Determines gradient offset value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gradientstop.html#cfn-quicksight-dashboard-gradientstop-gradientoffset
	//
	// Default: - 0.
	//
	GradientOffset *float64 `field:"optional" json:"gradientOffset" yaml:"gradientOffset"`
}


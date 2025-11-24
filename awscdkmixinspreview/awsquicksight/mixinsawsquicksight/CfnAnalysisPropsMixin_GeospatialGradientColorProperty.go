package mixinsawsquicksight


// The definition for a gradient color.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialGradientColorProperty := &GeospatialGradientColorProperty{
//   	DefaultOpacity: jsii.Number(123),
//   	NullDataSettings: &GeospatialNullDataSettingsProperty{
//   		SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   			FillColor: jsii.String("fillColor"),
//   			StrokeColor: jsii.String("strokeColor"),
//   			StrokeWidth: jsii.Number(123),
//   		},
//   	},
//   	NullDataVisibility: jsii.String("nullDataVisibility"),
//   	StepColors: []interface{}{
//   		&GeospatialGradientStepColorProperty{
//   			Color: jsii.String("color"),
//   			DataValue: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialgradientcolor.html
//
type CfnAnalysisPropsMixin_GeospatialGradientColorProperty struct {
	// The default opacity for the gradient color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialgradientcolor.html#cfn-quicksight-analysis-geospatialgradientcolor-defaultopacity
	//
	DefaultOpacity *float64 `field:"optional" json:"defaultOpacity" yaml:"defaultOpacity"`
	// The null data visualization settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialgradientcolor.html#cfn-quicksight-analysis-geospatialgradientcolor-nulldatasettings
	//
	NullDataSettings interface{} `field:"optional" json:"nullDataSettings" yaml:"nullDataSettings"`
	// The state of visibility for null data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialgradientcolor.html#cfn-quicksight-analysis-geospatialgradientcolor-nulldatavisibility
	//
	NullDataVisibility *string `field:"optional" json:"nullDataVisibility" yaml:"nullDataVisibility"`
	// A list of gradient step colors for the gradient.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialgradientcolor.html#cfn-quicksight-analysis-geospatialgradientcolor-stepcolors
	//
	StepColors interface{} `field:"optional" json:"stepColors" yaml:"stepColors"`
}


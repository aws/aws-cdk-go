package awsquicksight


// The definition for a gradient color.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialGradientColorProperty := &GeospatialGradientColorProperty{
//   	StepColors: []interface{}{
//   		&GeospatialGradientStepColorProperty{
//   			Color: jsii.String("color"),
//   			DataValue: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	DefaultOpacity: jsii.Number(123),
//   	NullDataSettings: &GeospatialNullDataSettingsProperty{
//   		SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   			FillColor: jsii.String("fillColor"),
//   			StrokeColor: jsii.String("strokeColor"),
//   			StrokeWidth: jsii.Number(123),
//   		},
//   	},
//   	NullDataVisibility: jsii.String("nullDataVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialgradientcolor.html
//
type CfnAnalysis_GeospatialGradientColorProperty struct {
	// A list of gradient step colors for the gradient.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialgradientcolor.html#cfn-quicksight-analysis-geospatialgradientcolor-stepcolors
	//
	StepColors interface{} `field:"required" json:"stepColors" yaml:"stepColors"`
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
}


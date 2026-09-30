package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialgradientcolor.html
//
type CfnTemplatePropsMixin_GeospatialGradientColorProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialgradientcolor.html#cfn-quicksight-template-geospatialgradientcolor-defaultopacity
	//
	DefaultOpacity *float64 `field:"optional" json:"defaultOpacity" yaml:"defaultOpacity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialgradientcolor.html#cfn-quicksight-template-geospatialgradientcolor-nulldatasettings
	//
	NullDataSettings interface{} `field:"optional" json:"nullDataSettings" yaml:"nullDataSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialgradientcolor.html#cfn-quicksight-template-geospatialgradientcolor-nulldatavisibility
	//
	NullDataVisibility *string `field:"optional" json:"nullDataVisibility" yaml:"nullDataVisibility"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialgradientcolor.html#cfn-quicksight-template-geospatialgradientcolor-stepcolors
	//
	StepColors interface{} `field:"optional" json:"stepColors" yaml:"stepColors"`
}


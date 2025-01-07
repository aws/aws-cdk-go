package awsquicksight


// The visualization properties for solid, gradient, and categorical colors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialColorProperty := &GeospatialColorProperty{
//   	Categorical: &GeospatialCategoricalColorProperty{
//   		CategoryDataColors: []interface{}{
//   			&GeospatialCategoricalDataColorProperty{
//   				Color: jsii.String("color"),
//   				DataValue: jsii.String("dataValue"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		DefaultOpacity: jsii.Number(123),
//   		NullDataSettings: &GeospatialNullDataSettingsProperty{
//   			SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   				FillColor: jsii.String("fillColor"),
//   				StrokeColor: jsii.String("strokeColor"),
//   				StrokeWidth: jsii.Number(123),
//   			},
//   		},
//   		NullDataVisibility: jsii.String("nullDataVisibility"),
//   	},
//   	Gradient: &GeospatialGradientColorProperty{
//   		StepColors: []interface{}{
//   			&GeospatialGradientStepColorProperty{
//   				Color: jsii.String("color"),
//   				DataValue: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		DefaultOpacity: jsii.Number(123),
//   		NullDataSettings: &GeospatialNullDataSettingsProperty{
//   			SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   				FillColor: jsii.String("fillColor"),
//   				StrokeColor: jsii.String("strokeColor"),
//   				StrokeWidth: jsii.Number(123),
//   			},
//   		},
//   		NullDataVisibility: jsii.String("nullDataVisibility"),
//   	},
//   	Solid: &GeospatialSolidColorProperty{
//   		Color: jsii.String("color"),
//
//   		// the properties below are optional
//   		State: jsii.String("state"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcolor.html
//
type CfnAnalysis_GeospatialColorProperty struct {
	// The visualization properties for the categorical color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcolor.html#cfn-quicksight-analysis-geospatialcolor-categorical
	//
	Categorical interface{} `field:"optional" json:"categorical" yaml:"categorical"`
	// The visualization properties for the gradient color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcolor.html#cfn-quicksight-analysis-geospatialcolor-gradient
	//
	Gradient interface{} `field:"optional" json:"gradient" yaml:"gradient"`
	// The visualization properties for the solid color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcolor.html#cfn-quicksight-analysis-geospatialcolor-solid
	//
	Solid interface{} `field:"optional" json:"solid" yaml:"solid"`
}


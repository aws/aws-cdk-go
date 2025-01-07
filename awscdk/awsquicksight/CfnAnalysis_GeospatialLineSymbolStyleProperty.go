package awsquicksight


// The symbol style for a line layer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialLineSymbolStyleProperty := &GeospatialLineSymbolStyleProperty{
//   	FillColor: &GeospatialColorProperty{
//   		Categorical: &GeospatialCategoricalColorProperty{
//   			CategoryDataColors: []interface{}{
//   				&GeospatialCategoricalDataColorProperty{
//   					Color: jsii.String("color"),
//   					DataValue: jsii.String("dataValue"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			DefaultOpacity: jsii.Number(123),
//   			NullDataSettings: &GeospatialNullDataSettingsProperty{
//   				SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   					FillColor: jsii.String("fillColor"),
//   					StrokeColor: jsii.String("strokeColor"),
//   					StrokeWidth: jsii.Number(123),
//   				},
//   			},
//   			NullDataVisibility: jsii.String("nullDataVisibility"),
//   		},
//   		Gradient: &GeospatialGradientColorProperty{
//   			StepColors: []interface{}{
//   				&GeospatialGradientStepColorProperty{
//   					Color: jsii.String("color"),
//   					DataValue: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			DefaultOpacity: jsii.Number(123),
//   			NullDataSettings: &GeospatialNullDataSettingsProperty{
//   				SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   					FillColor: jsii.String("fillColor"),
//   					StrokeColor: jsii.String("strokeColor"),
//   					StrokeWidth: jsii.Number(123),
//   				},
//   			},
//   			NullDataVisibility: jsii.String("nullDataVisibility"),
//   		},
//   		Solid: &GeospatialSolidColorProperty{
//   			Color: jsii.String("color"),
//
//   			// the properties below are optional
//   			State: jsii.String("state"),
//   		},
//   	},
//   	LineWidth: &GeospatialLineWidthProperty{
//   		LineWidth: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallinesymbolstyle.html
//
type CfnAnalysis_GeospatialLineSymbolStyleProperty struct {
	// The color and opacity values for the fill color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallinesymbolstyle.html#cfn-quicksight-analysis-geospatiallinesymbolstyle-fillcolor
	//
	FillColor interface{} `field:"optional" json:"fillColor" yaml:"fillColor"`
	// The width value for a line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallinesymbolstyle.html#cfn-quicksight-analysis-geospatiallinesymbolstyle-linewidth
	//
	LineWidth interface{} `field:"optional" json:"lineWidth" yaml:"lineWidth"`
}


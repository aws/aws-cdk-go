package mixinsawsquicksight


// The point style for a point layer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialPointStyleProperty := &GeospatialPointStyleProperty{
//   	CircleSymbolStyle: &GeospatialCircleSymbolStyleProperty{
//   		CircleRadius: &GeospatialCircleRadiusProperty{
//   			Radius: jsii.Number(123),
//   		},
//   		FillColor: &GeospatialColorProperty{
//   			Categorical: &GeospatialCategoricalColorProperty{
//   				CategoryDataColors: []interface{}{
//   					&GeospatialCategoricalDataColorProperty{
//   						Color: jsii.String("color"),
//   						DataValue: jsii.String("dataValue"),
//   					},
//   				},
//   				DefaultOpacity: jsii.Number(123),
//   				NullDataSettings: &GeospatialNullDataSettingsProperty{
//   					SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   						FillColor: jsii.String("fillColor"),
//   						StrokeColor: jsii.String("strokeColor"),
//   						StrokeWidth: jsii.Number(123),
//   					},
//   				},
//   				NullDataVisibility: jsii.String("nullDataVisibility"),
//   			},
//   			Gradient: &GeospatialGradientColorProperty{
//   				DefaultOpacity: jsii.Number(123),
//   				NullDataSettings: &GeospatialNullDataSettingsProperty{
//   					SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   						FillColor: jsii.String("fillColor"),
//   						StrokeColor: jsii.String("strokeColor"),
//   						StrokeWidth: jsii.Number(123),
//   					},
//   				},
//   				NullDataVisibility: jsii.String("nullDataVisibility"),
//   				StepColors: []interface{}{
//   					&GeospatialGradientStepColorProperty{
//   						Color: jsii.String("color"),
//   						DataValue: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Solid: &GeospatialSolidColorProperty{
//   				Color: jsii.String("color"),
//   				State: jsii.String("state"),
//   			},
//   		},
//   		StrokeColor: &GeospatialColorProperty{
//   			Categorical: &GeospatialCategoricalColorProperty{
//   				CategoryDataColors: []interface{}{
//   					&GeospatialCategoricalDataColorProperty{
//   						Color: jsii.String("color"),
//   						DataValue: jsii.String("dataValue"),
//   					},
//   				},
//   				DefaultOpacity: jsii.Number(123),
//   				NullDataSettings: &GeospatialNullDataSettingsProperty{
//   					SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   						FillColor: jsii.String("fillColor"),
//   						StrokeColor: jsii.String("strokeColor"),
//   						StrokeWidth: jsii.Number(123),
//   					},
//   				},
//   				NullDataVisibility: jsii.String("nullDataVisibility"),
//   			},
//   			Gradient: &GeospatialGradientColorProperty{
//   				DefaultOpacity: jsii.Number(123),
//   				NullDataSettings: &GeospatialNullDataSettingsProperty{
//   					SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   						FillColor: jsii.String("fillColor"),
//   						StrokeColor: jsii.String("strokeColor"),
//   						StrokeWidth: jsii.Number(123),
//   					},
//   				},
//   				NullDataVisibility: jsii.String("nullDataVisibility"),
//   				StepColors: []interface{}{
//   					&GeospatialGradientStepColorProperty{
//   						Color: jsii.String("color"),
//   						DataValue: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Solid: &GeospatialSolidColorProperty{
//   				Color: jsii.String("color"),
//   				State: jsii.String("state"),
//   			},
//   		},
//   		StrokeWidth: &GeospatialLineWidthProperty{
//   			LineWidth: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialpointstyle.html
//
type CfnAnalysisPropsMixin_GeospatialPointStyleProperty struct {
	// The circle symbol style for a point layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialpointstyle.html#cfn-quicksight-analysis-geospatialpointstyle-circlesymbolstyle
	//
	CircleSymbolStyle interface{} `field:"optional" json:"circleSymbolStyle" yaml:"circleSymbolStyle"`
}


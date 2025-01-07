package awsquicksight


// The point style for a point layer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//
//   				// the properties below are optional
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
//   				StepColors: []interface{}{
//   					&GeospatialGradientStepColorProperty{
//   						Color: jsii.String("color"),
//   						DataValue: jsii.Number(123),
//   					},
//   				},
//
//   				// the properties below are optional
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
//   			Solid: &GeospatialSolidColorProperty{
//   				Color: jsii.String("color"),
//
//   				// the properties below are optional
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
//
//   				// the properties below are optional
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
//   				StepColors: []interface{}{
//   					&GeospatialGradientStepColorProperty{
//   						Color: jsii.String("color"),
//   						DataValue: jsii.Number(123),
//   					},
//   				},
//
//   				// the properties below are optional
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
//   			Solid: &GeospatialSolidColorProperty{
//   				Color: jsii.String("color"),
//
//   				// the properties below are optional
//   				State: jsii.String("state"),
//   			},
//   		},
//   		StrokeWidth: &GeospatialLineWidthProperty{
//   			LineWidth: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpointstyle.html
//
type CfnDashboard_GeospatialPointStyleProperty struct {
	// The circle symbol style for a point layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpointstyle.html#cfn-quicksight-dashboard-geospatialpointstyle-circlesymbolstyle
	//
	CircleSymbolStyle interface{} `field:"optional" json:"circleSymbolStyle" yaml:"circleSymbolStyle"`
}


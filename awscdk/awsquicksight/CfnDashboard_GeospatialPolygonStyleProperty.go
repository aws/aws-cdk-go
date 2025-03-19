package awsquicksight


// The polygon style for a polygon layer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialPolygonStyleProperty := &GeospatialPolygonStyleProperty{
//   	PolygonSymbolStyle: &GeospatialPolygonSymbolStyleProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpolygonstyle.html
//
type CfnDashboard_GeospatialPolygonStyleProperty struct {
	// The polygon symbol style for a polygon layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpolygonstyle.html#cfn-quicksight-dashboard-geospatialpolygonstyle-polygonsymbolstyle
	//
	PolygonSymbolStyle interface{} `field:"optional" json:"polygonSymbolStyle" yaml:"polygonSymbolStyle"`
}


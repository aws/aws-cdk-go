package previewawsquicksightmixins


// The polygon symbol style for a polygon layer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialPolygonSymbolStyleProperty := &GeospatialPolygonSymbolStyleProperty{
//   	FillColor: &GeospatialColorProperty{
//   		Categorical: &GeospatialCategoricalColorProperty{
//   			CategoryDataColors: []interface{}{
//   				&GeospatialCategoricalDataColorProperty{
//   					Color: jsii.String("color"),
//   					DataValue: jsii.String("dataValue"),
//   				},
//   			},
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
//   			DefaultOpacity: jsii.Number(123),
//   			NullDataSettings: &GeospatialNullDataSettingsProperty{
//   				SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   					FillColor: jsii.String("fillColor"),
//   					StrokeColor: jsii.String("strokeColor"),
//   					StrokeWidth: jsii.Number(123),
//   				},
//   			},
//   			NullDataVisibility: jsii.String("nullDataVisibility"),
//   			StepColors: []interface{}{
//   				&GeospatialGradientStepColorProperty{
//   					Color: jsii.String("color"),
//   					DataValue: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Solid: &GeospatialSolidColorProperty{
//   			Color: jsii.String("color"),
//   			State: jsii.String("state"),
//   		},
//   	},
//   	StrokeColor: &GeospatialColorProperty{
//   		Categorical: &GeospatialCategoricalColorProperty{
//   			CategoryDataColors: []interface{}{
//   				&GeospatialCategoricalDataColorProperty{
//   					Color: jsii.String("color"),
//   					DataValue: jsii.String("dataValue"),
//   				},
//   			},
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
//   			DefaultOpacity: jsii.Number(123),
//   			NullDataSettings: &GeospatialNullDataSettingsProperty{
//   				SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   					FillColor: jsii.String("fillColor"),
//   					StrokeColor: jsii.String("strokeColor"),
//   					StrokeWidth: jsii.Number(123),
//   				},
//   			},
//   			NullDataVisibility: jsii.String("nullDataVisibility"),
//   			StepColors: []interface{}{
//   				&GeospatialGradientStepColorProperty{
//   					Color: jsii.String("color"),
//   					DataValue: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Solid: &GeospatialSolidColorProperty{
//   			Color: jsii.String("color"),
//   			State: jsii.String("state"),
//   		},
//   	},
//   	StrokeWidth: &GeospatialLineWidthProperty{
//   		LineWidth: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpolygonsymbolstyle.html
//
type CfnDashboardPropsMixin_GeospatialPolygonSymbolStyleProperty struct {
	// The color and opacity values for the fill color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpolygonsymbolstyle.html#cfn-quicksight-dashboard-geospatialpolygonsymbolstyle-fillcolor
	//
	FillColor interface{} `field:"optional" json:"fillColor" yaml:"fillColor"`
	// The color and opacity values for the stroke color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpolygonsymbolstyle.html#cfn-quicksight-dashboard-geospatialpolygonsymbolstyle-strokecolor
	//
	StrokeColor interface{} `field:"optional" json:"strokeColor" yaml:"strokeColor"`
	// The width of the border stroke.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpolygonsymbolstyle.html#cfn-quicksight-dashboard-geospatialpolygonsymbolstyle-strokewidth
	//
	StrokeWidth interface{} `field:"optional" json:"strokeWidth" yaml:"strokeWidth"`
}


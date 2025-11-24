package mixinsawsquicksight


// The geospatial Point layer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialPointLayerProperty := &GeospatialPointLayerProperty{
//   	Style: &GeospatialPointStyleProperty{
//   		CircleSymbolStyle: &GeospatialCircleSymbolStyleProperty{
//   			CircleRadius: &GeospatialCircleRadiusProperty{
//   				Radius: jsii.Number(123),
//   			},
//   			FillColor: &GeospatialColorProperty{
//   				Categorical: &GeospatialCategoricalColorProperty{
//   					CategoryDataColors: []interface{}{
//   						&GeospatialCategoricalDataColorProperty{
//   							Color: jsii.String("color"),
//   							DataValue: jsii.String("dataValue"),
//   						},
//   					},
//   					DefaultOpacity: jsii.Number(123),
//   					NullDataSettings: &GeospatialNullDataSettingsProperty{
//   						SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   							FillColor: jsii.String("fillColor"),
//   							StrokeColor: jsii.String("strokeColor"),
//   							StrokeWidth: jsii.Number(123),
//   						},
//   					},
//   					NullDataVisibility: jsii.String("nullDataVisibility"),
//   				},
//   				Gradient: &GeospatialGradientColorProperty{
//   					DefaultOpacity: jsii.Number(123),
//   					NullDataSettings: &GeospatialNullDataSettingsProperty{
//   						SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   							FillColor: jsii.String("fillColor"),
//   							StrokeColor: jsii.String("strokeColor"),
//   							StrokeWidth: jsii.Number(123),
//   						},
//   					},
//   					NullDataVisibility: jsii.String("nullDataVisibility"),
//   					StepColors: []interface{}{
//   						&GeospatialGradientStepColorProperty{
//   							Color: jsii.String("color"),
//   							DataValue: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Solid: &GeospatialSolidColorProperty{
//   					Color: jsii.String("color"),
//   					State: jsii.String("state"),
//   				},
//   			},
//   			StrokeColor: &GeospatialColorProperty{
//   				Categorical: &GeospatialCategoricalColorProperty{
//   					CategoryDataColors: []interface{}{
//   						&GeospatialCategoricalDataColorProperty{
//   							Color: jsii.String("color"),
//   							DataValue: jsii.String("dataValue"),
//   						},
//   					},
//   					DefaultOpacity: jsii.Number(123),
//   					NullDataSettings: &GeospatialNullDataSettingsProperty{
//   						SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   							FillColor: jsii.String("fillColor"),
//   							StrokeColor: jsii.String("strokeColor"),
//   							StrokeWidth: jsii.Number(123),
//   						},
//   					},
//   					NullDataVisibility: jsii.String("nullDataVisibility"),
//   				},
//   				Gradient: &GeospatialGradientColorProperty{
//   					DefaultOpacity: jsii.Number(123),
//   					NullDataSettings: &GeospatialNullDataSettingsProperty{
//   						SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   							FillColor: jsii.String("fillColor"),
//   							StrokeColor: jsii.String("strokeColor"),
//   							StrokeWidth: jsii.Number(123),
//   						},
//   					},
//   					NullDataVisibility: jsii.String("nullDataVisibility"),
//   					StepColors: []interface{}{
//   						&GeospatialGradientStepColorProperty{
//   							Color: jsii.String("color"),
//   							DataValue: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Solid: &GeospatialSolidColorProperty{
//   					Color: jsii.String("color"),
//   					State: jsii.String("state"),
//   				},
//   			},
//   			StrokeWidth: &GeospatialLineWidthProperty{
//   				LineWidth: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpointlayer.html
//
type CfnDashboardPropsMixin_GeospatialPointLayerProperty struct {
	// The visualization style for a point layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialpointlayer.html#cfn-quicksight-dashboard-geospatialpointlayer-style
	//
	Style interface{} `field:"optional" json:"style" yaml:"style"`
}


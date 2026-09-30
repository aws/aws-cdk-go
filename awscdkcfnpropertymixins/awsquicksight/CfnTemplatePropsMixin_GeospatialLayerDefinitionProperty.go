package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   geospatialLayerDefinitionProperty := &GeospatialLayerDefinitionProperty{
//   	LineLayer: &GeospatialLineLayerProperty{
//   		Style: &GeospatialLineStyleProperty{
//   			LineSymbolStyle: &GeospatialLineSymbolStyleProperty{
//   				FillColor: &GeospatialColorProperty{
//   					Categorical: &GeospatialCategoricalColorProperty{
//   						CategoryDataColors: []interface{}{
//   							&GeospatialCategoricalDataColorProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.String("dataValue"),
//   							},
//   						},
//   						DefaultOpacity: jsii.Number(123),
//   						NullDataSettings: &GeospatialNullDataSettingsProperty{
//   							SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   								FillColor: jsii.String("fillColor"),
//   								StrokeColor: jsii.String("strokeColor"),
//   								StrokeWidth: jsii.Number(123),
//   							},
//   						},
//   						NullDataVisibility: jsii.String("nullDataVisibility"),
//   					},
//   					Gradient: &GeospatialGradientColorProperty{
//   						DefaultOpacity: jsii.Number(123),
//   						NullDataSettings: &GeospatialNullDataSettingsProperty{
//   							SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   								FillColor: jsii.String("fillColor"),
//   								StrokeColor: jsii.String("strokeColor"),
//   								StrokeWidth: jsii.Number(123),
//   							},
//   						},
//   						NullDataVisibility: jsii.String("nullDataVisibility"),
//   						StepColors: []interface{}{
//   							&GeospatialGradientStepColorProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Solid: &GeospatialSolidColorProperty{
//   						Color: jsii.String("color"),
//   						State: jsii.String("state"),
//   					},
//   				},
//   				LineWidth: &GeospatialLineWidthProperty{
//   					LineWidth: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	PointLayer: &GeospatialPointLayerProperty{
//   		Style: &GeospatialPointStyleProperty{
//   			CircleSymbolStyle: &GeospatialCircleSymbolStyleProperty{
//   				CircleRadius: &GeospatialCircleRadiusProperty{
//   					Radius: jsii.Number(123),
//   				},
//   				FillColor: &GeospatialColorProperty{
//   					Categorical: &GeospatialCategoricalColorProperty{
//   						CategoryDataColors: []interface{}{
//   							&GeospatialCategoricalDataColorProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.String("dataValue"),
//   							},
//   						},
//   						DefaultOpacity: jsii.Number(123),
//   						NullDataSettings: &GeospatialNullDataSettingsProperty{
//   							SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   								FillColor: jsii.String("fillColor"),
//   								StrokeColor: jsii.String("strokeColor"),
//   								StrokeWidth: jsii.Number(123),
//   							},
//   						},
//   						NullDataVisibility: jsii.String("nullDataVisibility"),
//   					},
//   					Gradient: &GeospatialGradientColorProperty{
//   						DefaultOpacity: jsii.Number(123),
//   						NullDataSettings: &GeospatialNullDataSettingsProperty{
//   							SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   								FillColor: jsii.String("fillColor"),
//   								StrokeColor: jsii.String("strokeColor"),
//   								StrokeWidth: jsii.Number(123),
//   							},
//   						},
//   						NullDataVisibility: jsii.String("nullDataVisibility"),
//   						StepColors: []interface{}{
//   							&GeospatialGradientStepColorProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Solid: &GeospatialSolidColorProperty{
//   						Color: jsii.String("color"),
//   						State: jsii.String("state"),
//   					},
//   				},
//   				StrokeColor: &GeospatialColorProperty{
//   					Categorical: &GeospatialCategoricalColorProperty{
//   						CategoryDataColors: []interface{}{
//   							&GeospatialCategoricalDataColorProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.String("dataValue"),
//   							},
//   						},
//   						DefaultOpacity: jsii.Number(123),
//   						NullDataSettings: &GeospatialNullDataSettingsProperty{
//   							SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   								FillColor: jsii.String("fillColor"),
//   								StrokeColor: jsii.String("strokeColor"),
//   								StrokeWidth: jsii.Number(123),
//   							},
//   						},
//   						NullDataVisibility: jsii.String("nullDataVisibility"),
//   					},
//   					Gradient: &GeospatialGradientColorProperty{
//   						DefaultOpacity: jsii.Number(123),
//   						NullDataSettings: &GeospatialNullDataSettingsProperty{
//   							SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   								FillColor: jsii.String("fillColor"),
//   								StrokeColor: jsii.String("strokeColor"),
//   								StrokeWidth: jsii.Number(123),
//   							},
//   						},
//   						NullDataVisibility: jsii.String("nullDataVisibility"),
//   						StepColors: []interface{}{
//   							&GeospatialGradientStepColorProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Solid: &GeospatialSolidColorProperty{
//   						Color: jsii.String("color"),
//   						State: jsii.String("state"),
//   					},
//   				},
//   				StrokeWidth: &GeospatialLineWidthProperty{
//   					LineWidth: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	PolygonLayer: &GeospatialPolygonLayerProperty{
//   		Style: &GeospatialPolygonStyleProperty{
//   			PolygonSymbolStyle: &GeospatialPolygonSymbolStyleProperty{
//   				FillColor: &GeospatialColorProperty{
//   					Categorical: &GeospatialCategoricalColorProperty{
//   						CategoryDataColors: []interface{}{
//   							&GeospatialCategoricalDataColorProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.String("dataValue"),
//   							},
//   						},
//   						DefaultOpacity: jsii.Number(123),
//   						NullDataSettings: &GeospatialNullDataSettingsProperty{
//   							SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   								FillColor: jsii.String("fillColor"),
//   								StrokeColor: jsii.String("strokeColor"),
//   								StrokeWidth: jsii.Number(123),
//   							},
//   						},
//   						NullDataVisibility: jsii.String("nullDataVisibility"),
//   					},
//   					Gradient: &GeospatialGradientColorProperty{
//   						DefaultOpacity: jsii.Number(123),
//   						NullDataSettings: &GeospatialNullDataSettingsProperty{
//   							SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   								FillColor: jsii.String("fillColor"),
//   								StrokeColor: jsii.String("strokeColor"),
//   								StrokeWidth: jsii.Number(123),
//   							},
//   						},
//   						NullDataVisibility: jsii.String("nullDataVisibility"),
//   						StepColors: []interface{}{
//   							&GeospatialGradientStepColorProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Solid: &GeospatialSolidColorProperty{
//   						Color: jsii.String("color"),
//   						State: jsii.String("state"),
//   					},
//   				},
//   				StrokeColor: &GeospatialColorProperty{
//   					Categorical: &GeospatialCategoricalColorProperty{
//   						CategoryDataColors: []interface{}{
//   							&GeospatialCategoricalDataColorProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.String("dataValue"),
//   							},
//   						},
//   						DefaultOpacity: jsii.Number(123),
//   						NullDataSettings: &GeospatialNullDataSettingsProperty{
//   							SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   								FillColor: jsii.String("fillColor"),
//   								StrokeColor: jsii.String("strokeColor"),
//   								StrokeWidth: jsii.Number(123),
//   							},
//   						},
//   						NullDataVisibility: jsii.String("nullDataVisibility"),
//   					},
//   					Gradient: &GeospatialGradientColorProperty{
//   						DefaultOpacity: jsii.Number(123),
//   						NullDataSettings: &GeospatialNullDataSettingsProperty{
//   							SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   								FillColor: jsii.String("fillColor"),
//   								StrokeColor: jsii.String("strokeColor"),
//   								StrokeWidth: jsii.Number(123),
//   							},
//   						},
//   						NullDataVisibility: jsii.String("nullDataVisibility"),
//   						StepColors: []interface{}{
//   							&GeospatialGradientStepColorProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Solid: &GeospatialSolidColorProperty{
//   						Color: jsii.String("color"),
//   						State: jsii.String("state"),
//   					},
//   				},
//   				StrokeWidth: &GeospatialLineWidthProperty{
//   					LineWidth: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallayerdefinition.html
//
type CfnTemplatePropsMixin_GeospatialLayerDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallayerdefinition.html#cfn-quicksight-template-geospatiallayerdefinition-linelayer
	//
	LineLayer interface{} `field:"optional" json:"lineLayer" yaml:"lineLayer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallayerdefinition.html#cfn-quicksight-template-geospatiallayerdefinition-pointlayer
	//
	PointLayer interface{} `field:"optional" json:"pointLayer" yaml:"pointLayer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallayerdefinition.html#cfn-quicksight-template-geospatiallayerdefinition-polygonlayer
	//
	PolygonLayer interface{} `field:"optional" json:"polygonLayer" yaml:"polygonLayer"`
}


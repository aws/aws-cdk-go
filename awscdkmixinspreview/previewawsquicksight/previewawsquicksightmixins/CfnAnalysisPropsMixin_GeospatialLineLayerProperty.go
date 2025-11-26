package previewawsquicksightmixins


// The geospatial Line layer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialLineLayerProperty := &GeospatialLineLayerProperty{
//   	Style: &GeospatialLineStyleProperty{
//   		LineSymbolStyle: &GeospatialLineSymbolStyleProperty{
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
//   			LineWidth: &GeospatialLineWidthProperty{
//   				LineWidth: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallinelayer.html
//
type CfnAnalysisPropsMixin_GeospatialLineLayerProperty struct {
	// The visualization style for a line layer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallinelayer.html#cfn-quicksight-analysis-geospatiallinelayer-style
	//
	Style interface{} `field:"optional" json:"style" yaml:"style"`
}


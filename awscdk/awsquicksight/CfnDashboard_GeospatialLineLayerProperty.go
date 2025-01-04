package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//
//   					// the properties below are optional
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
//   					StepColors: []interface{}{
//   						&GeospatialGradientStepColorProperty{
//   							Color: jsii.String("color"),
//   							DataValue: jsii.Number(123),
//   						},
//   					},
//
//   					// the properties below are optional
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
//   				Solid: &GeospatialSolidColorProperty{
//   					Color: jsii.String("color"),
//
//   					// the properties below are optional
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallinelayer.html
//
type CfnDashboard_GeospatialLineLayerProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallinelayer.html#cfn-quicksight-dashboard-geospatiallinelayer-style
	//
	Style interface{} `field:"required" json:"style" yaml:"style"`
}


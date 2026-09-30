package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
//   	LineWidth: &GeospatialLineWidthProperty{
//   		LineWidth: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallinesymbolstyle.html
//
type CfnTemplatePropsMixin_GeospatialLineSymbolStyleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallinesymbolstyle.html#cfn-quicksight-template-geospatiallinesymbolstyle-fillcolor
	//
	FillColor interface{} `field:"optional" json:"fillColor" yaml:"fillColor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatiallinesymbolstyle.html#cfn-quicksight-template-geospatiallinesymbolstyle-linewidth
	//
	LineWidth interface{} `field:"optional" json:"lineWidth" yaml:"lineWidth"`
}


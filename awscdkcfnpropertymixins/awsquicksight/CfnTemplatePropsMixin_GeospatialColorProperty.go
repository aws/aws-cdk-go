package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   geospatialColorProperty := &GeospatialColorProperty{
//   	Categorical: &GeospatialCategoricalColorProperty{
//   		CategoryDataColors: []interface{}{
//   			&GeospatialCategoricalDataColorProperty{
//   				Color: jsii.String("color"),
//   				DataValue: jsii.String("dataValue"),
//   			},
//   		},
//   		DefaultOpacity: jsii.Number(123),
//   		NullDataSettings: &GeospatialNullDataSettingsProperty{
//   			SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   				FillColor: jsii.String("fillColor"),
//   				StrokeColor: jsii.String("strokeColor"),
//   				StrokeWidth: jsii.Number(123),
//   			},
//   		},
//   		NullDataVisibility: jsii.String("nullDataVisibility"),
//   	},
//   	Gradient: &GeospatialGradientColorProperty{
//   		DefaultOpacity: jsii.Number(123),
//   		NullDataSettings: &GeospatialNullDataSettingsProperty{
//   			SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   				FillColor: jsii.String("fillColor"),
//   				StrokeColor: jsii.String("strokeColor"),
//   				StrokeWidth: jsii.Number(123),
//   			},
//   		},
//   		NullDataVisibility: jsii.String("nullDataVisibility"),
//   		StepColors: []interface{}{
//   			&GeospatialGradientStepColorProperty{
//   				Color: jsii.String("color"),
//   				DataValue: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Solid: &GeospatialSolidColorProperty{
//   		Color: jsii.String("color"),
//   		State: jsii.String("state"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcolor.html
//
type CfnTemplatePropsMixin_GeospatialColorProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcolor.html#cfn-quicksight-template-geospatialcolor-categorical
	//
	Categorical interface{} `field:"optional" json:"categorical" yaml:"categorical"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcolor.html#cfn-quicksight-template-geospatialcolor-gradient
	//
	Gradient interface{} `field:"optional" json:"gradient" yaml:"gradient"`
	// Describes the properties for a solid color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcolor.html#cfn-quicksight-template-geospatialcolor-solid
	//
	Solid interface{} `field:"optional" json:"solid" yaml:"solid"`
}


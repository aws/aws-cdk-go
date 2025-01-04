package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialCategoricalColorProperty := &GeospatialCategoricalColorProperty{
//   	CategoryDataColors: []interface{}{
//   		&GeospatialCategoricalDataColorProperty{
//   			Color: jsii.String("color"),
//   			DataValue: jsii.String("dataValue"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	DefaultOpacity: jsii.Number(123),
//   	NullDataSettings: &GeospatialNullDataSettingsProperty{
//   		SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   			FillColor: jsii.String("fillColor"),
//   			StrokeColor: jsii.String("strokeColor"),
//   			StrokeWidth: jsii.Number(123),
//   		},
//   	},
//   	NullDataVisibility: jsii.String("nullDataVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcategoricalcolor.html
//
type CfnAnalysis_GeospatialCategoricalColorProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcategoricalcolor.html#cfn-quicksight-analysis-geospatialcategoricalcolor-categorydatacolors
	//
	CategoryDataColors interface{} `field:"required" json:"categoryDataColors" yaml:"categoryDataColors"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcategoricalcolor.html#cfn-quicksight-analysis-geospatialcategoricalcolor-defaultopacity
	//
	DefaultOpacity *float64 `field:"optional" json:"defaultOpacity" yaml:"defaultOpacity"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcategoricalcolor.html#cfn-quicksight-analysis-geospatialcategoricalcolor-nulldatasettings
	//
	NullDataSettings interface{} `field:"optional" json:"nullDataSettings" yaml:"nullDataSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcategoricalcolor.html#cfn-quicksight-analysis-geospatialcategoricalcolor-nulldatavisibility
	//
	NullDataVisibility *string `field:"optional" json:"nullDataVisibility" yaml:"nullDataVisibility"`
}


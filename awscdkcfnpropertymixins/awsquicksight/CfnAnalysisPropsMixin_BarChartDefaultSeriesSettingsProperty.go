package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   barChartDefaultSeriesSettingsProperty := &BarChartDefaultSeriesSettingsProperty{
//   	BorderSettings: &BorderSettingsProperty{
//   		BorderColor: jsii.String("borderColor"),
//   		BorderVisibility: jsii.String("borderVisibility"),
//   		BorderWidth: jsii.String("borderWidth"),
//   	},
//   	DecalSettings: &DecalSettingsProperty{
//   		DecalColor: jsii.String("decalColor"),
//   		DecalPatternType: jsii.String("decalPatternType"),
//   		DecalStyleType: jsii.String("decalStyleType"),
//   		DecalVisibility: jsii.String("decalVisibility"),
//   		ElementValue: jsii.String("elementValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-barchartdefaultseriessettings.html
//
type CfnAnalysisPropsMixin_BarChartDefaultSeriesSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-barchartdefaultseriessettings.html#cfn-quicksight-analysis-barchartdefaultseriessettings-bordersettings
	//
	BorderSettings interface{} `field:"optional" json:"borderSettings" yaml:"borderSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-barchartdefaultseriessettings.html#cfn-quicksight-analysis-barchartdefaultseriessettings-decalsettings
	//
	DecalSettings interface{} `field:"optional" json:"decalSettings" yaml:"decalSettings"`
}


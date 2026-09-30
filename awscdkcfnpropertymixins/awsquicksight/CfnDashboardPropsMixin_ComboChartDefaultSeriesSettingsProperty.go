package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   comboChartDefaultSeriesSettingsProperty := &ComboChartDefaultSeriesSettingsProperty{
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
//   	LineStyleSettings: &LineChartLineStyleSettingsProperty{
//   		LineInterpolation: jsii.String("lineInterpolation"),
//   		LineStyle: jsii.String("lineStyle"),
//   		LineVisibility: jsii.String("lineVisibility"),
//   		LineWidth: jsii.String("lineWidth"),
//   	},
//   	MarkerStyleSettings: &LineChartMarkerStyleSettingsProperty{
//   		MarkerColor: jsii.String("markerColor"),
//   		MarkerShape: jsii.String("markerShape"),
//   		MarkerSize: jsii.String("markerSize"),
//   		MarkerVisibility: jsii.String("markerVisibility"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartdefaultseriessettings.html
//
type CfnDashboardPropsMixin_ComboChartDefaultSeriesSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartdefaultseriessettings.html#cfn-quicksight-dashboard-combochartdefaultseriessettings-bordersettings
	//
	BorderSettings interface{} `field:"optional" json:"borderSettings" yaml:"borderSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartdefaultseriessettings.html#cfn-quicksight-dashboard-combochartdefaultseriessettings-decalsettings
	//
	DecalSettings interface{} `field:"optional" json:"decalSettings" yaml:"decalSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartdefaultseriessettings.html#cfn-quicksight-dashboard-combochartdefaultseriessettings-linestylesettings
	//
	LineStyleSettings interface{} `field:"optional" json:"lineStyleSettings" yaml:"lineStyleSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-combochartdefaultseriessettings.html#cfn-quicksight-dashboard-combochartdefaultseriessettings-markerstylesettings
	//
	MarkerStyleSettings interface{} `field:"optional" json:"markerStyleSettings" yaml:"markerStyleSettings"`
}


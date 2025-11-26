package previewawsquicksightmixins


// The series settings of a radar chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   radarChartSeriesSettingsProperty := &RadarChartSeriesSettingsProperty{
//   	AreaStyleSettings: &RadarChartAreaStyleSettingsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartseriessettings.html
//
type CfnDashboardPropsMixin_RadarChartSeriesSettingsProperty struct {
	// The area style settings of a radar chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-radarchartseriessettings.html#cfn-quicksight-dashboard-radarchartseriessettings-areastylesettings
	//
	AreaStyleSettings interface{} `field:"optional" json:"areaStyleSettings" yaml:"areaStyleSettings"`
}


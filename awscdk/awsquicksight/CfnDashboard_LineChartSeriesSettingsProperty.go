package awsquicksight


// The options that determine the presentation of a line series in the visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lineChartSeriesSettingsProperty := &LineChartSeriesSettingsProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartseriessettings.html
//
type CfnDashboard_LineChartSeriesSettingsProperty struct {
	// Line styles options for a line series in `LineChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartseriessettings.html#cfn-quicksight-dashboard-linechartseriessettings-linestylesettings
	//
	LineStyleSettings interface{} `field:"optional" json:"lineStyleSettings" yaml:"lineStyleSettings"`
	// Marker styles options for a line series in `LineChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartseriessettings.html#cfn-quicksight-dashboard-linechartseriessettings-markerstylesettings
	//
	MarkerStyleSettings interface{} `field:"optional" json:"markerStyleSettings" yaml:"markerStyleSettings"`
}


package awsquicksight


// The options that determine the default presentation of all line series in `LineChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lineChartDefaultSeriesSettingsProperty := &LineChartDefaultSeriesSettingsProperty{
//   	AxisBinding: jsii.String("axisBinding"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartdefaultseriessettings.html
//
type CfnAnalysis_LineChartDefaultSeriesSettingsProperty struct {
	// The axis to which you are binding all line series to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartdefaultseriessettings.html#cfn-quicksight-analysis-linechartdefaultseriessettings-axisbinding
	//
	AxisBinding *string `field:"optional" json:"axisBinding" yaml:"axisBinding"`
	// Line styles options for all line series in the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartdefaultseriessettings.html#cfn-quicksight-analysis-linechartdefaultseriessettings-linestylesettings
	//
	LineStyleSettings interface{} `field:"optional" json:"lineStyleSettings" yaml:"lineStyleSettings"`
	// Marker styles options for all line series in the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartdefaultseriessettings.html#cfn-quicksight-analysis-linechartdefaultseriessettings-markerstylesettings
	//
	MarkerStyleSettings interface{} `field:"optional" json:"markerStyleSettings" yaml:"markerStyleSettings"`
}


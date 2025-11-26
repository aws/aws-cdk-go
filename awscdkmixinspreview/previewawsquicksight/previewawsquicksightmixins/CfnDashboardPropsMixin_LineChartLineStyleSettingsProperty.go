package previewawsquicksightmixins


// Line styles options for a line series in `LineChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lineChartLineStyleSettingsProperty := &LineChartLineStyleSettingsProperty{
//   	LineInterpolation: jsii.String("lineInterpolation"),
//   	LineStyle: jsii.String("lineStyle"),
//   	LineVisibility: jsii.String("lineVisibility"),
//   	LineWidth: jsii.String("lineWidth"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartlinestylesettings.html
//
type CfnDashboardPropsMixin_LineChartLineStyleSettingsProperty struct {
	// Interpolation style for line series.
	//
	// - `LINEAR` : Show as default, linear style.
	// - `SMOOTH` : Show as a smooth curve.
	// - `STEPPED` : Show steps in line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartlinestylesettings.html#cfn-quicksight-dashboard-linechartlinestylesettings-lineinterpolation
	//
	LineInterpolation *string `field:"optional" json:"lineInterpolation" yaml:"lineInterpolation"`
	// Line style for line series.
	//
	// - `SOLID` : Show as a solid line.
	// - `DOTTED` : Show as a dotted line.
	// - `DASHED` : Show as a dashed line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartlinestylesettings.html#cfn-quicksight-dashboard-linechartlinestylesettings-linestyle
	//
	LineStyle *string `field:"optional" json:"lineStyle" yaml:"lineStyle"`
	// Configuration option that determines whether to show the line for the series.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartlinestylesettings.html#cfn-quicksight-dashboard-linechartlinestylesettings-linevisibility
	//
	LineVisibility *string `field:"optional" json:"lineVisibility" yaml:"lineVisibility"`
	// Width that determines the line thickness.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartlinestylesettings.html#cfn-quicksight-dashboard-linechartlinestylesettings-linewidth
	//
	LineWidth *string `field:"optional" json:"lineWidth" yaml:"lineWidth"`
}


package awsquicksight


// Line styles options for a line series in `LineChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lineChartLineStyleSettingsProperty := &LineChartLineStyleSettingsProperty{
//   	LineInterpolation: jsii.String("lineInterpolation"),
//   	LineStyle: jsii.String("lineStyle"),
//   	LineVisibility: jsii.String("lineVisibility"),
//   	LineWidth: jsii.String("lineWidth"),
//   }
//
type CfnDashboard_LineChartLineStyleSettingsProperty struct {
	// Interpolation style for line series.
	//
	// - `LINEAR` : Show as default, linear style.
	// - `SMOOTH` : Show as a smooth curve.
	// - `STEPPED` : Show steps in line.
	LineInterpolation *string `field:"optional" json:"lineInterpolation" yaml:"lineInterpolation"`
	// Line style for line series.
	//
	// - `SOLID` : Show as a solid line.
	// - `DOTTED` : Show as a dotted line.
	// - `DASHED` : Show as a dashed line.
	LineStyle *string `field:"optional" json:"lineStyle" yaml:"lineStyle"`
	// Configuration option that determines whether to show the line for the series.
	LineVisibility *string `field:"optional" json:"lineVisibility" yaml:"lineVisibility"`
	// Width that determines the line thickness.
	LineWidth *string `field:"optional" json:"lineWidth" yaml:"lineWidth"`
}


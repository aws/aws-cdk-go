package awsquicksight


// Marker styles options for a line series in `LineChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lineChartMarkerStyleSettingsProperty := &LineChartMarkerStyleSettingsProperty{
//   	MarkerColor: jsii.String("markerColor"),
//   	MarkerShape: jsii.String("markerShape"),
//   	MarkerSize: jsii.String("markerSize"),
//   	MarkerVisibility: jsii.String("markerVisibility"),
//   }
//
type CfnAnalysis_LineChartMarkerStyleSettingsProperty struct {
	// Color of marker in the series.
	MarkerColor *string `field:"optional" json:"markerColor" yaml:"markerColor"`
	// Shape option for markers in the series.
	//
	// - `CIRCLE` : Show marker as a circle.
	// - `TRIANGLE` : Show marker as a triangle.
	// - `SQUARE` : Show marker as a square.
	// - `DIAMOND` : Show marker as a diamond.
	// - `ROUNDED_SQUARE` : Show marker as a rounded square.
	MarkerShape *string `field:"optional" json:"markerShape" yaml:"markerShape"`
	// Size of marker in the series.
	MarkerSize *string `field:"optional" json:"markerSize" yaml:"markerSize"`
	// Configuration option that determines whether to show the markers in the series.
	MarkerVisibility *string `field:"optional" json:"markerVisibility" yaml:"markerVisibility"`
}


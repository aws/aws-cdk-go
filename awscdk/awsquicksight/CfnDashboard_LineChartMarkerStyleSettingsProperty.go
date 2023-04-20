package awsquicksight


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
type CfnDashboard_LineChartMarkerStyleSettingsProperty struct {
	// `CfnDashboard.LineChartMarkerStyleSettingsProperty.MarkerColor`.
	MarkerColor *string `field:"optional" json:"markerColor" yaml:"markerColor"`
	// `CfnDashboard.LineChartMarkerStyleSettingsProperty.MarkerShape`.
	MarkerShape *string `field:"optional" json:"markerShape" yaml:"markerShape"`
	// `CfnDashboard.LineChartMarkerStyleSettingsProperty.MarkerSize`.
	MarkerSize *string `field:"optional" json:"markerSize" yaml:"markerSize"`
	// `CfnDashboard.LineChartMarkerStyleSettingsProperty.MarkerVisibility`.
	MarkerVisibility *string `field:"optional" json:"markerVisibility" yaml:"markerVisibility"`
}


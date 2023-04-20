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
type CfnAnalysis_LineChartMarkerStyleSettingsProperty struct {
	// `CfnAnalysis.LineChartMarkerStyleSettingsProperty.MarkerColor`.
	MarkerColor *string `field:"optional" json:"markerColor" yaml:"markerColor"`
	// `CfnAnalysis.LineChartMarkerStyleSettingsProperty.MarkerShape`.
	MarkerShape *string `field:"optional" json:"markerShape" yaml:"markerShape"`
	// `CfnAnalysis.LineChartMarkerStyleSettingsProperty.MarkerSize`.
	MarkerSize *string `field:"optional" json:"markerSize" yaml:"markerSize"`
	// `CfnAnalysis.LineChartMarkerStyleSettingsProperty.MarkerVisibility`.
	MarkerVisibility *string `field:"optional" json:"markerVisibility" yaml:"markerVisibility"`
}


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
type CfnTemplate_LineChartMarkerStyleSettingsProperty struct {
	// `CfnTemplate.LineChartMarkerStyleSettingsProperty.MarkerColor`.
	MarkerColor *string `field:"optional" json:"markerColor" yaml:"markerColor"`
	// `CfnTemplate.LineChartMarkerStyleSettingsProperty.MarkerShape`.
	MarkerShape *string `field:"optional" json:"markerShape" yaml:"markerShape"`
	// `CfnTemplate.LineChartMarkerStyleSettingsProperty.MarkerSize`.
	MarkerSize *string `field:"optional" json:"markerSize" yaml:"markerSize"`
	// `CfnTemplate.LineChartMarkerStyleSettingsProperty.MarkerVisibility`.
	MarkerVisibility *string `field:"optional" json:"markerVisibility" yaml:"markerVisibility"`
}


package awsquicksight


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
type CfnAnalysis_LineChartSeriesSettingsProperty struct {
	// `CfnAnalysis.LineChartSeriesSettingsProperty.LineStyleSettings`.
	LineStyleSettings interface{} `field:"optional" json:"lineStyleSettings" yaml:"lineStyleSettings"`
	// `CfnAnalysis.LineChartSeriesSettingsProperty.MarkerStyleSettings`.
	MarkerStyleSettings interface{} `field:"optional" json:"markerStyleSettings" yaml:"markerStyleSettings"`
}


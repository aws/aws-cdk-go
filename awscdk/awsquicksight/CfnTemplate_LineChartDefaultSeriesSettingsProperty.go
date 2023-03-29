package awsquicksight


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
type CfnTemplate_LineChartDefaultSeriesSettingsProperty struct {
	// `CfnTemplate.LineChartDefaultSeriesSettingsProperty.AxisBinding`.
	AxisBinding *string `field:"optional" json:"axisBinding" yaml:"axisBinding"`
	// `CfnTemplate.LineChartDefaultSeriesSettingsProperty.LineStyleSettings`.
	LineStyleSettings interface{} `field:"optional" json:"lineStyleSettings" yaml:"lineStyleSettings"`
	// `CfnTemplate.LineChartDefaultSeriesSettingsProperty.MarkerStyleSettings`.
	MarkerStyleSettings interface{} `field:"optional" json:"markerStyleSettings" yaml:"markerStyleSettings"`
}


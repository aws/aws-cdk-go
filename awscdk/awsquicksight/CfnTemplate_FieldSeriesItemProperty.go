package awsquicksight


// The field series item configuration of a line chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldSeriesItemProperty := &FieldSeriesItemProperty{
//   	AxisBinding: jsii.String("axisBinding"),
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	Settings: &LineChartSeriesSettingsProperty{
//   		LineStyleSettings: &LineChartLineStyleSettingsProperty{
//   			LineInterpolation: jsii.String("lineInterpolation"),
//   			LineStyle: jsii.String("lineStyle"),
//   			LineVisibility: jsii.String("lineVisibility"),
//   			LineWidth: jsii.String("lineWidth"),
//   		},
//   		MarkerStyleSettings: &LineChartMarkerStyleSettingsProperty{
//   			MarkerColor: jsii.String("markerColor"),
//   			MarkerShape: jsii.String("markerShape"),
//   			MarkerSize: jsii.String("markerSize"),
//   			MarkerVisibility: jsii.String("markerVisibility"),
//   		},
//   	},
//   }
//
type CfnTemplate_FieldSeriesItemProperty struct {
	// The axis that you are binding the field to.
	AxisBinding *string `field:"required" json:"axisBinding" yaml:"axisBinding"`
	// The field ID of the field for which you are setting the axis binding.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The options that determine the presentation of line series associated to the field.
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}


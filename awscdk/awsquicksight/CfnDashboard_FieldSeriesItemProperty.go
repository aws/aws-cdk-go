package awsquicksight


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
type CfnDashboard_FieldSeriesItemProperty struct {
	// `CfnDashboard.FieldSeriesItemProperty.AxisBinding`.
	AxisBinding *string `field:"required" json:"axisBinding" yaml:"axisBinding"`
	// `CfnDashboard.FieldSeriesItemProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// `CfnDashboard.FieldSeriesItemProperty.Settings`.
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}


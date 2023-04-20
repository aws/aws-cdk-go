package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataFieldSeriesItemProperty := &DataFieldSeriesItemProperty{
//   	AxisBinding: jsii.String("axisBinding"),
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	FieldValue: jsii.String("fieldValue"),
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
type CfnTemplate_DataFieldSeriesItemProperty struct {
	// `CfnTemplate.DataFieldSeriesItemProperty.AxisBinding`.
	AxisBinding *string `field:"required" json:"axisBinding" yaml:"axisBinding"`
	// `CfnTemplate.DataFieldSeriesItemProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// `CfnTemplate.DataFieldSeriesItemProperty.FieldValue`.
	FieldValue *string `field:"optional" json:"fieldValue" yaml:"fieldValue"`
	// `CfnTemplate.DataFieldSeriesItemProperty.Settings`.
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}


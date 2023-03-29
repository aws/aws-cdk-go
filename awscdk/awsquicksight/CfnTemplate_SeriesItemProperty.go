package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   seriesItemProperty := &SeriesItemProperty{
//   	DataFieldSeriesItem: &DataFieldSeriesItemProperty{
//   		AxisBinding: jsii.String("axisBinding"),
//   		FieldId: jsii.String("fieldId"),
//
//   		// the properties below are optional
//   		FieldValue: jsii.String("fieldValue"),
//   		Settings: &LineChartSeriesSettingsProperty{
//   			LineStyleSettings: &LineChartLineStyleSettingsProperty{
//   				LineInterpolation: jsii.String("lineInterpolation"),
//   				LineStyle: jsii.String("lineStyle"),
//   				LineVisibility: jsii.String("lineVisibility"),
//   				LineWidth: jsii.String("lineWidth"),
//   			},
//   			MarkerStyleSettings: &LineChartMarkerStyleSettingsProperty{
//   				MarkerColor: jsii.String("markerColor"),
//   				MarkerShape: jsii.String("markerShape"),
//   				MarkerSize: jsii.String("markerSize"),
//   				MarkerVisibility: jsii.String("markerVisibility"),
//   			},
//   		},
//   	},
//   	FieldSeriesItem: &FieldSeriesItemProperty{
//   		AxisBinding: jsii.String("axisBinding"),
//   		FieldId: jsii.String("fieldId"),
//
//   		// the properties below are optional
//   		Settings: &LineChartSeriesSettingsProperty{
//   			LineStyleSettings: &LineChartLineStyleSettingsProperty{
//   				LineInterpolation: jsii.String("lineInterpolation"),
//   				LineStyle: jsii.String("lineStyle"),
//   				LineVisibility: jsii.String("lineVisibility"),
//   				LineWidth: jsii.String("lineWidth"),
//   			},
//   			MarkerStyleSettings: &LineChartMarkerStyleSettingsProperty{
//   				MarkerColor: jsii.String("markerColor"),
//   				MarkerShape: jsii.String("markerShape"),
//   				MarkerSize: jsii.String("markerSize"),
//   				MarkerVisibility: jsii.String("markerVisibility"),
//   			},
//   		},
//   	},
//   }
//
type CfnTemplate_SeriesItemProperty struct {
	// `CfnTemplate.SeriesItemProperty.DataFieldSeriesItem`.
	DataFieldSeriesItem interface{} `field:"optional" json:"dataFieldSeriesItem" yaml:"dataFieldSeriesItem"`
	// `CfnTemplate.SeriesItemProperty.FieldSeriesItem`.
	FieldSeriesItem interface{} `field:"optional" json:"fieldSeriesItem" yaml:"fieldSeriesItem"`
}


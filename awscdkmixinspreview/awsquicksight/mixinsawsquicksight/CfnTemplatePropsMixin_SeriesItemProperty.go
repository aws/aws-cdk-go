package mixinsawsquicksight


// The series item configuration of a line chart.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   seriesItemProperty := &SeriesItemProperty{
//   	DataFieldSeriesItem: &DataFieldSeriesItemProperty{
//   		AxisBinding: jsii.String("axisBinding"),
//   		FieldId: jsii.String("fieldId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-seriesitem.html
//
type CfnTemplatePropsMixin_SeriesItemProperty struct {
	// The data field series item configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-seriesitem.html#cfn-quicksight-template-seriesitem-datafieldseriesitem
	//
	DataFieldSeriesItem interface{} `field:"optional" json:"dataFieldSeriesItem" yaml:"dataFieldSeriesItem"`
	// The field series item configuration of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-seriesitem.html#cfn-quicksight-template-seriesitem-fieldseriesitem
	//
	FieldSeriesItem interface{} `field:"optional" json:"fieldSeriesItem" yaml:"fieldSeriesItem"`
}


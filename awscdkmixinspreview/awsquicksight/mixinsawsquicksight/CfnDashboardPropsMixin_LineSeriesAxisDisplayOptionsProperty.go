package mixinsawsquicksight


// The series axis configuration of a line chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var dataDriven interface{}
//
//   lineSeriesAxisDisplayOptionsProperty := &LineSeriesAxisDisplayOptionsProperty{
//   	AxisOptions: &AxisDisplayOptionsProperty{
//   		AxisLineVisibility: jsii.String("axisLineVisibility"),
//   		AxisOffset: jsii.String("axisOffset"),
//   		DataOptions: &AxisDataOptionsProperty{
//   			DateAxisOptions: &DateAxisOptionsProperty{
//   				MissingDateVisibility: jsii.String("missingDateVisibility"),
//   			},
//   			NumericAxisOptions: &NumericAxisOptionsProperty{
//   				Range: &AxisDisplayRangeProperty{
//   					DataDriven: dataDriven,
//   					MinMax: &AxisDisplayMinMaxRangeProperty{
//   						Maximum: jsii.Number(123),
//   						Minimum: jsii.Number(123),
//   					},
//   				},
//   				Scale: &AxisScaleProperty{
//   					Linear: &AxisLinearScaleProperty{
//   						StepCount: jsii.Number(123),
//   						StepSize: jsii.Number(123),
//   					},
//   					Logarithmic: &AxisLogarithmicScaleProperty{
//   						Base: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		GridLineVisibility: jsii.String("gridLineVisibility"),
//   		ScrollbarOptions: &ScrollBarOptionsProperty{
//   			Visibility: jsii.String("visibility"),
//   			VisibleRange: &VisibleRangeOptionsProperty{
//   				PercentRange: &PercentVisibleRangeProperty{
//   					From: jsii.Number(123),
//   					To: jsii.Number(123),
//   				},
//   			},
//   		},
//   		TickLabelOptions: &AxisTickLabelOptionsProperty{
//   			LabelOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontFamily: jsii.String("fontFamily"),
//   					FontSize: &FontSizeProperty{
//   						Absolute: jsii.String("absolute"),
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   			RotationAngle: jsii.Number(123),
//   		},
//   	},
//   	MissingDataConfigurations: []interface{}{
//   		&MissingDataConfigurationProperty{
//   			TreatmentOption: jsii.String("treatmentOption"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-lineseriesaxisdisplayoptions.html
//
type CfnDashboardPropsMixin_LineSeriesAxisDisplayOptionsProperty struct {
	// The options that determine the presentation of the line series axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-lineseriesaxisdisplayoptions.html#cfn-quicksight-dashboard-lineseriesaxisdisplayoptions-axisoptions
	//
	AxisOptions interface{} `field:"optional" json:"axisOptions" yaml:"axisOptions"`
	// The configuration options that determine how missing data is treated during the rendering of a line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-lineseriesaxisdisplayoptions.html#cfn-quicksight-dashboard-lineseriesaxisdisplayoptions-missingdataconfigurations
	//
	MissingDataConfigurations interface{} `field:"optional" json:"missingDataConfigurations" yaml:"missingDataConfigurations"`
}


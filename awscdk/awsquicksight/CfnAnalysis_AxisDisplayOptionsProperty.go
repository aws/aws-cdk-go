package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataDriven interface{}
//
//   axisDisplayOptionsProperty := &AxisDisplayOptionsProperty{
//   	AxisLineVisibility: jsii.String("axisLineVisibility"),
//   	AxisOffset: jsii.String("axisOffset"),
//   	DataOptions: &AxisDataOptionsProperty{
//   		DateAxisOptions: &DateAxisOptionsProperty{
//   			MissingDateVisibility: jsii.String("missingDateVisibility"),
//   		},
//   		NumericAxisOptions: &NumericAxisOptionsProperty{
//   			Range: &AxisDisplayRangeProperty{
//   				DataDriven: dataDriven,
//   				MinMax: &AxisDisplayMinMaxRangeProperty{
//   					Maximum: jsii.Number(123),
//   					Minimum: jsii.Number(123),
//   				},
//   			},
//   			Scale: &AxisScaleProperty{
//   				Linear: &AxisLinearScaleProperty{
//   					StepCount: jsii.Number(123),
//   					StepSize: jsii.Number(123),
//   				},
//   				Logarithmic: &AxisLogarithmicScaleProperty{
//   					Base: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	GridLineVisibility: jsii.String("gridLineVisibility"),
//   	ScrollbarOptions: &ScrollBarOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   		VisibleRange: &VisibleRangeOptionsProperty{
//   			PercentRange: &PercentVisibleRangeProperty{
//   				From: jsii.Number(123),
//   				To: jsii.Number(123),
//   			},
//   		},
//   	},
//   	TickLabelOptions: &AxisTickLabelOptionsProperty{
//   		LabelOptions: &LabelOptionsProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontSize: &FontSizeProperty{
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Visibility: jsii.String("visibility"),
//   		},
//   		RotationAngle: jsii.Number(123),
//   	},
//   }
//
type CfnAnalysis_AxisDisplayOptionsProperty struct {
	// `CfnAnalysis.AxisDisplayOptionsProperty.AxisLineVisibility`.
	AxisLineVisibility *string `field:"optional" json:"axisLineVisibility" yaml:"axisLineVisibility"`
	// `CfnAnalysis.AxisDisplayOptionsProperty.AxisOffset`.
	AxisOffset *string `field:"optional" json:"axisOffset" yaml:"axisOffset"`
	// `CfnAnalysis.AxisDisplayOptionsProperty.DataOptions`.
	DataOptions interface{} `field:"optional" json:"dataOptions" yaml:"dataOptions"`
	// `CfnAnalysis.AxisDisplayOptionsProperty.GridLineVisibility`.
	GridLineVisibility *string `field:"optional" json:"gridLineVisibility" yaml:"gridLineVisibility"`
	// `CfnAnalysis.AxisDisplayOptionsProperty.ScrollbarOptions`.
	ScrollbarOptions interface{} `field:"optional" json:"scrollbarOptions" yaml:"scrollbarOptions"`
	// `CfnAnalysis.AxisDisplayOptionsProperty.TickLabelOptions`.
	TickLabelOptions interface{} `field:"optional" json:"tickLabelOptions" yaml:"tickLabelOptions"`
}


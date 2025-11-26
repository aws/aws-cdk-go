package previewawsquicksightmixins


// The options that determine the presentation of the arc of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gaugeChartArcConditionalFormattingProperty := &GaugeChartArcConditionalFormattingProperty{
//   	ForegroundColor: &ConditionalFormattingColorProperty{
//   		Gradient: &ConditionalFormattingGradientColorProperty{
//   			Color: &GradientColorProperty{
//   				Stops: []interface{}{
//   					&GradientStopProperty{
//   						Color: jsii.String("color"),
//   						DataValue: jsii.Number(123),
//   						GradientOffset: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Expression: jsii.String("expression"),
//   		},
//   		Solid: &ConditionalFormattingSolidColorProperty{
//   			Color: jsii.String("color"),
//   			Expression: jsii.String("expression"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartarcconditionalformatting.html
//
type CfnDashboardPropsMixin_GaugeChartArcConditionalFormattingProperty struct {
	// The conditional formatting of the arc foreground color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartarcconditionalformatting.html#cfn-quicksight-dashboard-gaugechartarcconditionalformatting-foregroundcolor
	//
	ForegroundColor interface{} `field:"optional" json:"foregroundColor" yaml:"foregroundColor"`
}


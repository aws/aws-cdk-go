package previewawsquicksightmixins


// Conditional formatting options of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gaugeChartConditionalFormattingOptionProperty := &GaugeChartConditionalFormattingOptionProperty{
//   	Arc: &GaugeChartArcConditionalFormattingProperty{
//   		ForegroundColor: &ConditionalFormattingColorProperty{
//   			Gradient: &ConditionalFormattingGradientColorProperty{
//   				Color: &GradientColorProperty{
//   					Stops: []interface{}{
//   						&GradientStopProperty{
//   							Color: jsii.String("color"),
//   							DataValue: jsii.Number(123),
//   							GradientOffset: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Expression: jsii.String("expression"),
//   			},
//   			Solid: &ConditionalFormattingSolidColorProperty{
//   				Color: jsii.String("color"),
//   				Expression: jsii.String("expression"),
//   			},
//   		},
//   	},
//   	PrimaryValue: &GaugeChartPrimaryValueConditionalFormattingProperty{
//   		Icon: &ConditionalFormattingIconProperty{
//   			CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   				Color: jsii.String("color"),
//   				DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   					IconDisplayOption: jsii.String("iconDisplayOption"),
//   				},
//   				Expression: jsii.String("expression"),
//   				IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   					Icon: jsii.String("icon"),
//   					UnicodeIcon: jsii.String("unicodeIcon"),
//   				},
//   			},
//   			IconSet: &ConditionalFormattingIconSetProperty{
//   				Expression: jsii.String("expression"),
//   				IconSetType: jsii.String("iconSetType"),
//   			},
//   		},
//   		TextColor: &ConditionalFormattingColorProperty{
//   			Gradient: &ConditionalFormattingGradientColorProperty{
//   				Color: &GradientColorProperty{
//   					Stops: []interface{}{
//   						&GradientStopProperty{
//   							Color: jsii.String("color"),
//   							DataValue: jsii.Number(123),
//   							GradientOffset: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Expression: jsii.String("expression"),
//   			},
//   			Solid: &ConditionalFormattingSolidColorProperty{
//   				Color: jsii.String("color"),
//   				Expression: jsii.String("expression"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconditionalformattingoption.html
//
type CfnAnalysisPropsMixin_GaugeChartConditionalFormattingOptionProperty struct {
	// The options that determine the presentation of the arc of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconditionalformattingoption.html#cfn-quicksight-analysis-gaugechartconditionalformattingoption-arc
	//
	Arc interface{} `field:"optional" json:"arc" yaml:"arc"`
	// The conditional formatting for the primary value of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gaugechartconditionalformattingoption.html#cfn-quicksight-analysis-gaugechartconditionalformattingoption-primaryvalue
	//
	PrimaryValue interface{} `field:"optional" json:"primaryValue" yaml:"primaryValue"`
}


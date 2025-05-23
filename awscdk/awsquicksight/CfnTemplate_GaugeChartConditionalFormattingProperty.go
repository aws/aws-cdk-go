package awsquicksight


// The conditional formatting of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gaugeChartConditionalFormattingProperty := &GaugeChartConditionalFormattingProperty{
//   	ConditionalFormattingOptions: []interface{}{
//   		&GaugeChartConditionalFormattingOptionProperty{
//   			Arc: &GaugeChartArcConditionalFormattingProperty{
//   				ForegroundColor: &ConditionalFormattingColorProperty{
//   					Gradient: &ConditionalFormattingGradientColorProperty{
//   						Color: &GradientColorProperty{
//   							Stops: []interface{}{
//   								&GradientStopProperty{
//   									GradientOffset: jsii.Number(123),
//
//   									// the properties below are optional
//   									Color: jsii.String("color"),
//   									DataValue: jsii.Number(123),
//   								},
//   							},
//   						},
//   						Expression: jsii.String("expression"),
//   					},
//   					Solid: &ConditionalFormattingSolidColorProperty{
//   						Expression: jsii.String("expression"),
//
//   						// the properties below are optional
//   						Color: jsii.String("color"),
//   					},
//   				},
//   			},
//   			PrimaryValue: &GaugeChartPrimaryValueConditionalFormattingProperty{
//   				Icon: &ConditionalFormattingIconProperty{
//   					CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   						Expression: jsii.String("expression"),
//   						IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   							Icon: jsii.String("icon"),
//   							UnicodeIcon: jsii.String("unicodeIcon"),
//   						},
//
//   						// the properties below are optional
//   						Color: jsii.String("color"),
//   						DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   							IconDisplayOption: jsii.String("iconDisplayOption"),
//   						},
//   					},
//   					IconSet: &ConditionalFormattingIconSetProperty{
//   						Expression: jsii.String("expression"),
//
//   						// the properties below are optional
//   						IconSetType: jsii.String("iconSetType"),
//   					},
//   				},
//   				TextColor: &ConditionalFormattingColorProperty{
//   					Gradient: &ConditionalFormattingGradientColorProperty{
//   						Color: &GradientColorProperty{
//   							Stops: []interface{}{
//   								&GradientStopProperty{
//   									GradientOffset: jsii.Number(123),
//
//   									// the properties below are optional
//   									Color: jsii.String("color"),
//   									DataValue: jsii.Number(123),
//   								},
//   							},
//   						},
//   						Expression: jsii.String("expression"),
//   					},
//   					Solid: &ConditionalFormattingSolidColorProperty{
//   						Expression: jsii.String("expression"),
//
//   						// the properties below are optional
//   						Color: jsii.String("color"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconditionalformatting.html
//
type CfnTemplate_GaugeChartConditionalFormattingProperty struct {
	// Conditional formatting options of a `GaugeChartVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartconditionalformatting.html#cfn-quicksight-template-gaugechartconditionalformatting-conditionalformattingoptions
	//
	ConditionalFormattingOptions interface{} `field:"optional" json:"conditionalFormattingOptions" yaml:"conditionalFormattingOptions"`
}


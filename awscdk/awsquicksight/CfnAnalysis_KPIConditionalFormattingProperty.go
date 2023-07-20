package awsquicksight


// The conditional formatting of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kPIConditionalFormattingProperty := &KPIConditionalFormattingProperty{
//   	ConditionalFormattingOptions: []interface{}{
//   		&KPIConditionalFormattingOptionProperty{
//   			PrimaryValue: &KPIPrimaryValueConditionalFormattingProperty{
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
//   			ProgressBar: &KPIProgressBarConditionalFormattingProperty{
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
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpiconditionalformatting.html
//
type CfnAnalysis_KPIConditionalFormattingProperty struct {
	// The conditional formatting options of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpiconditionalformatting.html#cfn-quicksight-analysis-kpiconditionalformatting-conditionalformattingoptions
	//
	ConditionalFormattingOptions interface{} `field:"optional" json:"conditionalFormattingOptions" yaml:"conditionalFormattingOptions"`
}


package mixinsawsquicksight


// The conditional formatting of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kPIConditionalFormattingProperty := &KPIConditionalFormattingProperty{
//   	ConditionalFormattingOptions: []interface{}{
//   		&KPIConditionalFormattingOptionProperty{
//   			ActualValue: &KPIActualValueConditionalFormattingProperty{
//   				Icon: &ConditionalFormattingIconProperty{
//   					CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   						Color: jsii.String("color"),
//   						DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   							IconDisplayOption: jsii.String("iconDisplayOption"),
//   						},
//   						Expression: jsii.String("expression"),
//   						IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   							Icon: jsii.String("icon"),
//   							UnicodeIcon: jsii.String("unicodeIcon"),
//   						},
//   					},
//   					IconSet: &ConditionalFormattingIconSetProperty{
//   						Expression: jsii.String("expression"),
//   						IconSetType: jsii.String("iconSetType"),
//   					},
//   				},
//   				TextColor: &ConditionalFormattingColorProperty{
//   					Gradient: &ConditionalFormattingGradientColorProperty{
//   						Color: &GradientColorProperty{
//   							Stops: []interface{}{
//   								&GradientStopProperty{
//   									Color: jsii.String("color"),
//   									DataValue: jsii.Number(123),
//   									GradientOffset: jsii.Number(123),
//   								},
//   							},
//   						},
//   						Expression: jsii.String("expression"),
//   					},
//   					Solid: &ConditionalFormattingSolidColorProperty{
//   						Color: jsii.String("color"),
//   						Expression: jsii.String("expression"),
//   					},
//   				},
//   			},
//   			ComparisonValue: &KPIComparisonValueConditionalFormattingProperty{
//   				Icon: &ConditionalFormattingIconProperty{
//   					CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   						Color: jsii.String("color"),
//   						DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   							IconDisplayOption: jsii.String("iconDisplayOption"),
//   						},
//   						Expression: jsii.String("expression"),
//   						IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   							Icon: jsii.String("icon"),
//   							UnicodeIcon: jsii.String("unicodeIcon"),
//   						},
//   					},
//   					IconSet: &ConditionalFormattingIconSetProperty{
//   						Expression: jsii.String("expression"),
//   						IconSetType: jsii.String("iconSetType"),
//   					},
//   				},
//   				TextColor: &ConditionalFormattingColorProperty{
//   					Gradient: &ConditionalFormattingGradientColorProperty{
//   						Color: &GradientColorProperty{
//   							Stops: []interface{}{
//   								&GradientStopProperty{
//   									Color: jsii.String("color"),
//   									DataValue: jsii.Number(123),
//   									GradientOffset: jsii.Number(123),
//   								},
//   							},
//   						},
//   						Expression: jsii.String("expression"),
//   					},
//   					Solid: &ConditionalFormattingSolidColorProperty{
//   						Color: jsii.String("color"),
//   						Expression: jsii.String("expression"),
//   					},
//   				},
//   			},
//   			PrimaryValue: &KPIPrimaryValueConditionalFormattingProperty{
//   				Icon: &ConditionalFormattingIconProperty{
//   					CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   						Color: jsii.String("color"),
//   						DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   							IconDisplayOption: jsii.String("iconDisplayOption"),
//   						},
//   						Expression: jsii.String("expression"),
//   						IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   							Icon: jsii.String("icon"),
//   							UnicodeIcon: jsii.String("unicodeIcon"),
//   						},
//   					},
//   					IconSet: &ConditionalFormattingIconSetProperty{
//   						Expression: jsii.String("expression"),
//   						IconSetType: jsii.String("iconSetType"),
//   					},
//   				},
//   				TextColor: &ConditionalFormattingColorProperty{
//   					Gradient: &ConditionalFormattingGradientColorProperty{
//   						Color: &GradientColorProperty{
//   							Stops: []interface{}{
//   								&GradientStopProperty{
//   									Color: jsii.String("color"),
//   									DataValue: jsii.Number(123),
//   									GradientOffset: jsii.Number(123),
//   								},
//   							},
//   						},
//   						Expression: jsii.String("expression"),
//   					},
//   					Solid: &ConditionalFormattingSolidColorProperty{
//   						Color: jsii.String("color"),
//   						Expression: jsii.String("expression"),
//   					},
//   				},
//   			},
//   			ProgressBar: &KPIProgressBarConditionalFormattingProperty{
//   				ForegroundColor: &ConditionalFormattingColorProperty{
//   					Gradient: &ConditionalFormattingGradientColorProperty{
//   						Color: &GradientColorProperty{
//   							Stops: []interface{}{
//   								&GradientStopProperty{
//   									Color: jsii.String("color"),
//   									DataValue: jsii.Number(123),
//   									GradientOffset: jsii.Number(123),
//   								},
//   							},
//   						},
//   						Expression: jsii.String("expression"),
//   					},
//   					Solid: &ConditionalFormattingSolidColorProperty{
//   						Color: jsii.String("color"),
//   						Expression: jsii.String("expression"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpiconditionalformatting.html
//
type CfnTemplatePropsMixin_KPIConditionalFormattingProperty struct {
	// The conditional formatting options of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpiconditionalformatting.html#cfn-quicksight-template-kpiconditionalformatting-conditionalformattingoptions
	//
	ConditionalFormattingOptions interface{} `field:"optional" json:"conditionalFormattingOptions" yaml:"conditionalFormattingOptions"`
}


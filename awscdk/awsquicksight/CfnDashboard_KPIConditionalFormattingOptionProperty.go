package awsquicksight


// The conditional formatting options of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kPIConditionalFormattingOptionProperty := &KPIConditionalFormattingOptionProperty{
//   	ActualValue: &KPIActualValueConditionalFormattingProperty{
//   		Icon: &ConditionalFormattingIconProperty{
//   			CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   				Expression: jsii.String("expression"),
//   				IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   					Icon: jsii.String("icon"),
//   					UnicodeIcon: jsii.String("unicodeIcon"),
//   				},
//
//   				// the properties below are optional
//   				Color: jsii.String("color"),
//   				DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   					IconDisplayOption: jsii.String("iconDisplayOption"),
//   				},
//   			},
//   			IconSet: &ConditionalFormattingIconSetProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				IconSetType: jsii.String("iconSetType"),
//   			},
//   		},
//   		TextColor: &ConditionalFormattingColorProperty{
//   			Gradient: &ConditionalFormattingGradientColorProperty{
//   				Color: &GradientColorProperty{
//   					Stops: []interface{}{
//   						&GradientStopProperty{
//   							GradientOffset: jsii.Number(123),
//
//   							// the properties below are optional
//   							Color: jsii.String("color"),
//   							DataValue: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Expression: jsii.String("expression"),
//   			},
//   			Solid: &ConditionalFormattingSolidColorProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				Color: jsii.String("color"),
//   			},
//   		},
//   	},
//   	ComparisonValue: &KPIComparisonValueConditionalFormattingProperty{
//   		Icon: &ConditionalFormattingIconProperty{
//   			CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   				Expression: jsii.String("expression"),
//   				IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   					Icon: jsii.String("icon"),
//   					UnicodeIcon: jsii.String("unicodeIcon"),
//   				},
//
//   				// the properties below are optional
//   				Color: jsii.String("color"),
//   				DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   					IconDisplayOption: jsii.String("iconDisplayOption"),
//   				},
//   			},
//   			IconSet: &ConditionalFormattingIconSetProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				IconSetType: jsii.String("iconSetType"),
//   			},
//   		},
//   		TextColor: &ConditionalFormattingColorProperty{
//   			Gradient: &ConditionalFormattingGradientColorProperty{
//   				Color: &GradientColorProperty{
//   					Stops: []interface{}{
//   						&GradientStopProperty{
//   							GradientOffset: jsii.Number(123),
//
//   							// the properties below are optional
//   							Color: jsii.String("color"),
//   							DataValue: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Expression: jsii.String("expression"),
//   			},
//   			Solid: &ConditionalFormattingSolidColorProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				Color: jsii.String("color"),
//   			},
//   		},
//   	},
//   	PrimaryValue: &KPIPrimaryValueConditionalFormattingProperty{
//   		Icon: &ConditionalFormattingIconProperty{
//   			CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   				Expression: jsii.String("expression"),
//   				IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   					Icon: jsii.String("icon"),
//   					UnicodeIcon: jsii.String("unicodeIcon"),
//   				},
//
//   				// the properties below are optional
//   				Color: jsii.String("color"),
//   				DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   					IconDisplayOption: jsii.String("iconDisplayOption"),
//   				},
//   			},
//   			IconSet: &ConditionalFormattingIconSetProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				IconSetType: jsii.String("iconSetType"),
//   			},
//   		},
//   		TextColor: &ConditionalFormattingColorProperty{
//   			Gradient: &ConditionalFormattingGradientColorProperty{
//   				Color: &GradientColorProperty{
//   					Stops: []interface{}{
//   						&GradientStopProperty{
//   							GradientOffset: jsii.Number(123),
//
//   							// the properties below are optional
//   							Color: jsii.String("color"),
//   							DataValue: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Expression: jsii.String("expression"),
//   			},
//   			Solid: &ConditionalFormattingSolidColorProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				Color: jsii.String("color"),
//   			},
//   		},
//   	},
//   	ProgressBar: &KPIProgressBarConditionalFormattingProperty{
//   		ForegroundColor: &ConditionalFormattingColorProperty{
//   			Gradient: &ConditionalFormattingGradientColorProperty{
//   				Color: &GradientColorProperty{
//   					Stops: []interface{}{
//   						&GradientStopProperty{
//   							GradientOffset: jsii.Number(123),
//
//   							// the properties below are optional
//   							Color: jsii.String("color"),
//   							DataValue: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Expression: jsii.String("expression"),
//   			},
//   			Solid: &ConditionalFormattingSolidColorProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				Color: jsii.String("color"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiconditionalformattingoption.html
//
type CfnDashboard_KPIConditionalFormattingOptionProperty struct {
	// The conditional formatting for the actual value of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiconditionalformattingoption.html#cfn-quicksight-dashboard-kpiconditionalformattingoption-actualvalue
	//
	ActualValue interface{} `field:"optional" json:"actualValue" yaml:"actualValue"`
	// The conditional formatting for the comparison value of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiconditionalformattingoption.html#cfn-quicksight-dashboard-kpiconditionalformattingoption-comparisonvalue
	//
	ComparisonValue interface{} `field:"optional" json:"comparisonValue" yaml:"comparisonValue"`
	// The conditional formatting for the primary value of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiconditionalformattingoption.html#cfn-quicksight-dashboard-kpiconditionalformattingoption-primaryvalue
	//
	PrimaryValue interface{} `field:"optional" json:"primaryValue" yaml:"primaryValue"`
	// The conditional formatting for the progress bar of a KPI visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiconditionalformattingoption.html#cfn-quicksight-dashboard-kpiconditionalformattingoption-progressbar
	//
	ProgressBar interface{} `field:"optional" json:"progressBar" yaml:"progressBar"`
}


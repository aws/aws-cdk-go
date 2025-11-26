package previewawsquicksightmixins


// The conditional formatting for the comparison value of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kPIComparisonValueConditionalFormattingProperty := &KPIComparisonValueConditionalFormattingProperty{
//   	Icon: &ConditionalFormattingIconProperty{
//   		CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   			Color: jsii.String("color"),
//   			DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   				IconDisplayOption: jsii.String("iconDisplayOption"),
//   			},
//   			Expression: jsii.String("expression"),
//   			IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   				Icon: jsii.String("icon"),
//   				UnicodeIcon: jsii.String("unicodeIcon"),
//   			},
//   		},
//   		IconSet: &ConditionalFormattingIconSetProperty{
//   			Expression: jsii.String("expression"),
//   			IconSetType: jsii.String("iconSetType"),
//   		},
//   	},
//   	TextColor: &ConditionalFormattingColorProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpicomparisonvalueconditionalformatting.html
//
type CfnAnalysisPropsMixin_KPIComparisonValueConditionalFormattingProperty struct {
	// The conditional formatting of the comparison value's icon.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpicomparisonvalueconditionalformatting.html#cfn-quicksight-analysis-kpicomparisonvalueconditionalformatting-icon
	//
	Icon interface{} `field:"optional" json:"icon" yaml:"icon"`
	// The conditional formatting of the comparison value's text color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-kpicomparisonvalueconditionalformatting.html#cfn-quicksight-analysis-kpicomparisonvalueconditionalformatting-textcolor
	//
	TextColor interface{} `field:"optional" json:"textColor" yaml:"textColor"`
}


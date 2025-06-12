package awsquicksight


// The conditional formatting for a `PivotTableVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableConditionalFormattingProperty := &TableConditionalFormattingProperty{
//   	ConditionalFormattingOptions: []interface{}{
//   		&TableConditionalFormattingOptionProperty{
//   			Cell: &TableCellConditionalFormattingProperty{
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				TextFormat: &TextConditionalFormatProperty{
//   					BackgroundColor: &ConditionalFormattingColorProperty{
//   						Gradient: &ConditionalFormattingGradientColorProperty{
//   							Color: &GradientColorProperty{
//   								Stops: []interface{}{
//   									&GradientStopProperty{
//   										GradientOffset: jsii.Number(123),
//
//   										// the properties below are optional
//   										Color: jsii.String("color"),
//   										DataValue: jsii.Number(123),
//   									},
//   								},
//   							},
//   							Expression: jsii.String("expression"),
//   						},
//   						Solid: &ConditionalFormattingSolidColorProperty{
//   							Expression: jsii.String("expression"),
//
//   							// the properties below are optional
//   							Color: jsii.String("color"),
//   						},
//   					},
//   					Icon: &ConditionalFormattingIconProperty{
//   						CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   							Expression: jsii.String("expression"),
//   							IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   								Icon: jsii.String("icon"),
//   								UnicodeIcon: jsii.String("unicodeIcon"),
//   							},
//
//   							// the properties below are optional
//   							Color: jsii.String("color"),
//   							DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   								IconDisplayOption: jsii.String("iconDisplayOption"),
//   							},
//   						},
//   						IconSet: &ConditionalFormattingIconSetProperty{
//   							Expression: jsii.String("expression"),
//
//   							// the properties below are optional
//   							IconSetType: jsii.String("iconSetType"),
//   						},
//   					},
//   					TextColor: &ConditionalFormattingColorProperty{
//   						Gradient: &ConditionalFormattingGradientColorProperty{
//   							Color: &GradientColorProperty{
//   								Stops: []interface{}{
//   									&GradientStopProperty{
//   										GradientOffset: jsii.Number(123),
//
//   										// the properties below are optional
//   										Color: jsii.String("color"),
//   										DataValue: jsii.Number(123),
//   									},
//   								},
//   							},
//   							Expression: jsii.String("expression"),
//   						},
//   						Solid: &ConditionalFormattingSolidColorProperty{
//   							Expression: jsii.String("expression"),
//
//   							// the properties below are optional
//   							Color: jsii.String("color"),
//   						},
//   					},
//   				},
//   			},
//   			Row: &TableRowConditionalFormattingProperty{
//   				BackgroundColor: &ConditionalFormattingColorProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tableconditionalformatting.html
//
type CfnAnalysis_TableConditionalFormattingProperty struct {
	// Conditional formatting options for a `PivotTableVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tableconditionalformatting.html#cfn-quicksight-analysis-tableconditionalformatting-conditionalformattingoptions
	//
	ConditionalFormattingOptions interface{} `field:"optional" json:"conditionalFormattingOptions" yaml:"conditionalFormattingOptions"`
}


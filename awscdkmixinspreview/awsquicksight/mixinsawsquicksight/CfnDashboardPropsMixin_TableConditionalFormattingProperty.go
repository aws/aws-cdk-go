package mixinsawsquicksight


// The conditional formatting for a `PivotTableVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableConditionalFormattingProperty := &TableConditionalFormattingProperty{
//   	ConditionalFormattingOptions: []interface{}{
//   		&TableConditionalFormattingOptionProperty{
//   			Cell: &TableCellConditionalFormattingProperty{
//   				FieldId: jsii.String("fieldId"),
//   				TextFormat: &TextConditionalFormatProperty{
//   					BackgroundColor: &ConditionalFormattingColorProperty{
//   						Gradient: &ConditionalFormattingGradientColorProperty{
//   							Color: &GradientColorProperty{
//   								Stops: []interface{}{
//   									&GradientStopProperty{
//   										Color: jsii.String("color"),
//   										DataValue: jsii.Number(123),
//   										GradientOffset: jsii.Number(123),
//   									},
//   								},
//   							},
//   							Expression: jsii.String("expression"),
//   						},
//   						Solid: &ConditionalFormattingSolidColorProperty{
//   							Color: jsii.String("color"),
//   							Expression: jsii.String("expression"),
//   						},
//   					},
//   					Icon: &ConditionalFormattingIconProperty{
//   						CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   							Color: jsii.String("color"),
//   							DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   								IconDisplayOption: jsii.String("iconDisplayOption"),
//   							},
//   							Expression: jsii.String("expression"),
//   							IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   								Icon: jsii.String("icon"),
//   								UnicodeIcon: jsii.String("unicodeIcon"),
//   							},
//   						},
//   						IconSet: &ConditionalFormattingIconSetProperty{
//   							Expression: jsii.String("expression"),
//   							IconSetType: jsii.String("iconSetType"),
//   						},
//   					},
//   					TextColor: &ConditionalFormattingColorProperty{
//   						Gradient: &ConditionalFormattingGradientColorProperty{
//   							Color: &GradientColorProperty{
//   								Stops: []interface{}{
//   									&GradientStopProperty{
//   										Color: jsii.String("color"),
//   										DataValue: jsii.Number(123),
//   										GradientOffset: jsii.Number(123),
//   									},
//   								},
//   							},
//   							Expression: jsii.String("expression"),
//   						},
//   						Solid: &ConditionalFormattingSolidColorProperty{
//   							Color: jsii.String("color"),
//   							Expression: jsii.String("expression"),
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
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableconditionalformatting.html
//
type CfnDashboardPropsMixin_TableConditionalFormattingProperty struct {
	// Conditional formatting options for a `PivotTableVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tableconditionalformatting.html#cfn-quicksight-dashboard-tableconditionalformatting-conditionalformattingoptions
	//
	ConditionalFormattingOptions interface{} `field:"optional" json:"conditionalFormattingOptions" yaml:"conditionalFormattingOptions"`
}


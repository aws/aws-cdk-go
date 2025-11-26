package previewawsquicksightmixins


// The conditional formatting for a `PivotTableVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotTableConditionalFormattingProperty := &PivotTableConditionalFormattingProperty{
//   	ConditionalFormattingOptions: []interface{}{
//   		&PivotTableConditionalFormattingOptionProperty{
//   			Cell: &PivotTableCellConditionalFormattingProperty{
//   				FieldId: jsii.String("fieldId"),
//   				Scope: &PivotTableConditionalFormattingScopeProperty{
//   					Role: jsii.String("role"),
//   				},
//   				Scopes: []interface{}{
//   					&PivotTableConditionalFormattingScopeProperty{
//   						Role: jsii.String("role"),
//   					},
//   				},
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
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconditionalformatting.html
//
type CfnDashboardPropsMixin_PivotTableConditionalFormattingProperty struct {
	// Conditional formatting options for a `PivotTableVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconditionalformatting.html#cfn-quicksight-dashboard-pivottableconditionalformatting-conditionalformattingoptions
	//
	ConditionalFormattingOptions interface{} `field:"optional" json:"conditionalFormattingOptions" yaml:"conditionalFormattingOptions"`
}


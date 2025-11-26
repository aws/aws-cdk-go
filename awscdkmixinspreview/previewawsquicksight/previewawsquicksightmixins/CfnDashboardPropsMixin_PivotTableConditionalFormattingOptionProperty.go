package previewawsquicksightmixins


// Conditional formatting options for a `PivotTableVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotTableConditionalFormattingOptionProperty := &PivotTableConditionalFormattingOptionProperty{
//   	Cell: &PivotTableCellConditionalFormattingProperty{
//   		FieldId: jsii.String("fieldId"),
//   		Scope: &PivotTableConditionalFormattingScopeProperty{
//   			Role: jsii.String("role"),
//   		},
//   		Scopes: []interface{}{
//   			&PivotTableConditionalFormattingScopeProperty{
//   				Role: jsii.String("role"),
//   			},
//   		},
//   		TextFormat: &TextConditionalFormatProperty{
//   			BackgroundColor: &ConditionalFormattingColorProperty{
//   				Gradient: &ConditionalFormattingGradientColorProperty{
//   					Color: &GradientColorProperty{
//   						Stops: []interface{}{
//   							&GradientStopProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.Number(123),
//   								GradientOffset: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Expression: jsii.String("expression"),
//   				},
//   				Solid: &ConditionalFormattingSolidColorProperty{
//   					Color: jsii.String("color"),
//   					Expression: jsii.String("expression"),
//   				},
//   			},
//   			Icon: &ConditionalFormattingIconProperty{
//   				CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   					Color: jsii.String("color"),
//   					DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   						IconDisplayOption: jsii.String("iconDisplayOption"),
//   					},
//   					Expression: jsii.String("expression"),
//   					IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   						Icon: jsii.String("icon"),
//   						UnicodeIcon: jsii.String("unicodeIcon"),
//   					},
//   				},
//   				IconSet: &ConditionalFormattingIconSetProperty{
//   					Expression: jsii.String("expression"),
//   					IconSetType: jsii.String("iconSetType"),
//   				},
//   			},
//   			TextColor: &ConditionalFormattingColorProperty{
//   				Gradient: &ConditionalFormattingGradientColorProperty{
//   					Color: &GradientColorProperty{
//   						Stops: []interface{}{
//   							&GradientStopProperty{
//   								Color: jsii.String("color"),
//   								DataValue: jsii.Number(123),
//   								GradientOffset: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Expression: jsii.String("expression"),
//   				},
//   				Solid: &ConditionalFormattingSolidColorProperty{
//   					Color: jsii.String("color"),
//   					Expression: jsii.String("expression"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconditionalformattingoption.html
//
type CfnDashboardPropsMixin_PivotTableConditionalFormattingOptionProperty struct {
	// The cell conditional formatting option for a pivot table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconditionalformattingoption.html#cfn-quicksight-dashboard-pivottableconditionalformattingoption-cell
	//
	Cell interface{} `field:"optional" json:"cell" yaml:"cell"`
}


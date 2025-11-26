package previewawsquicksightmixins


// The cell conditional formatting option for a table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableCellConditionalFormattingProperty := &TableCellConditionalFormattingProperty{
//   	FieldId: jsii.String("fieldId"),
//   	TextFormat: &TextConditionalFormatProperty{
//   		BackgroundColor: &ConditionalFormattingColorProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablecellconditionalformatting.html
//
type CfnTemplatePropsMixin_TableCellConditionalFormattingProperty struct {
	// The field ID of the cell for conditional formatting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablecellconditionalformatting.html#cfn-quicksight-template-tablecellconditionalformatting-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The text format of the cell for conditional formatting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablecellconditionalformatting.html#cfn-quicksight-template-tablecellconditionalformatting-textformat
	//
	TextFormat interface{} `field:"optional" json:"textFormat" yaml:"textFormat"`
}


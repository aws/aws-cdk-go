package previewawsquicksightmixins


// The cell conditional formatting option for a pivot table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotTableCellConditionalFormattingProperty := &PivotTableCellConditionalFormattingProperty{
//   	FieldId: jsii.String("fieldId"),
//   	Scope: &PivotTableConditionalFormattingScopeProperty{
//   		Role: jsii.String("role"),
//   	},
//   	Scopes: []interface{}{
//   		&PivotTableConditionalFormattingScopeProperty{
//   			Role: jsii.String("role"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablecellconditionalformatting.html
//
type CfnDashboardPropsMixin_PivotTableCellConditionalFormattingProperty struct {
	// The field ID of the cell for conditional formatting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablecellconditionalformatting.html#cfn-quicksight-dashboard-pivottablecellconditionalformatting-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The scope of the cell for conditional formatting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablecellconditionalformatting.html#cfn-quicksight-dashboard-pivottablecellconditionalformatting-scope
	//
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// A list of cell scopes for conditional formatting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablecellconditionalformatting.html#cfn-quicksight-dashboard-pivottablecellconditionalformatting-scopes
	//
	Scopes interface{} `field:"optional" json:"scopes" yaml:"scopes"`
	// The text format of the cell for conditional formatting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablecellconditionalformatting.html#cfn-quicksight-dashboard-pivottablecellconditionalformatting-textformat
	//
	TextFormat interface{} `field:"optional" json:"textFormat" yaml:"textFormat"`
}


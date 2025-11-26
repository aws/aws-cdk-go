package previewawsquicksightmixins


// The conditional formatting of a table row.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableRowConditionalFormattingProperty := &TableRowConditionalFormattingProperty{
//   	BackgroundColor: &ConditionalFormattingColorProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablerowconditionalformatting.html
//
type CfnAnalysisPropsMixin_TableRowConditionalFormattingProperty struct {
	// The conditional formatting color (solid, gradient) of the background for a table row.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablerowconditionalformatting.html#cfn-quicksight-analysis-tablerowconditionalformatting-backgroundcolor
	//
	BackgroundColor interface{} `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The conditional formatting color (solid, gradient) of the text for a table row.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablerowconditionalformatting.html#cfn-quicksight-analysis-tablerowconditionalformatting-textcolor
	//
	TextColor interface{} `field:"optional" json:"textColor" yaml:"textColor"`
}


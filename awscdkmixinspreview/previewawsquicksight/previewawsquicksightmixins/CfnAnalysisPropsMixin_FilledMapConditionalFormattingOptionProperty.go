package previewawsquicksightmixins


// Conditional formatting options of a `FilledMapVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filledMapConditionalFormattingOptionProperty := &FilledMapConditionalFormattingOptionProperty{
//   	Shape: &FilledMapShapeConditionalFormattingProperty{
//   		FieldId: jsii.String("fieldId"),
//   		Format: &ShapeConditionalFormatProperty{
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
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconditionalformattingoption.html
//
type CfnAnalysisPropsMixin_FilledMapConditionalFormattingOptionProperty struct {
	// The conditional formatting that determines the shape of the filled map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapconditionalformattingoption.html#cfn-quicksight-analysis-filledmapconditionalformattingoption-shape
	//
	Shape interface{} `field:"optional" json:"shape" yaml:"shape"`
}


package previewawsquicksightmixins


// The shape conditional formatting of a filled map visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   shapeConditionalFormatProperty := &ShapeConditionalFormatProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-shapeconditionalformat.html
//
type CfnAnalysisPropsMixin_ShapeConditionalFormatProperty struct {
	// The conditional formatting for the shape background color of a filled map visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-shapeconditionalformat.html#cfn-quicksight-analysis-shapeconditionalformat-backgroundcolor
	//
	BackgroundColor interface{} `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
}


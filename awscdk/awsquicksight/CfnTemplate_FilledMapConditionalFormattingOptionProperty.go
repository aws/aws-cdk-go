package awsquicksight


// Conditional formatting options of a `FilledMapVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filledMapConditionalFormattingOptionProperty := &FilledMapConditionalFormattingOptionProperty{
//   	Shape: &FilledMapShapeConditionalFormattingProperty{
//   		FieldId: jsii.String("fieldId"),
//
//   		// the properties below are optional
//   		Format: &ShapeConditionalFormatProperty{
//   			BackgroundColor: &ConditionalFormattingColorProperty{
//   				Gradient: &ConditionalFormattingGradientColorProperty{
//   					Color: &GradientColorProperty{
//   						Stops: []interface{}{
//   							&GradientStopProperty{
//   								GradientOffset: jsii.Number(123),
//
//   								// the properties below are optional
//   								Color: jsii.String("color"),
//   								DataValue: jsii.Number(123),
//   							},
//   						},
//   					},
//   					Expression: jsii.String("expression"),
//   				},
//   				Solid: &ConditionalFormattingSolidColorProperty{
//   					Expression: jsii.String("expression"),
//
//   					// the properties below are optional
//   					Color: jsii.String("color"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filledmapconditionalformattingoption.html
//
type CfnTemplate_FilledMapConditionalFormattingOptionProperty struct {
	// The conditional formatting that determines the shape of the filled map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filledmapconditionalformattingoption.html#cfn-quicksight-template-filledmapconditionalformattingoption-shape
	//
	Shape interface{} `field:"required" json:"shape" yaml:"shape"`
}


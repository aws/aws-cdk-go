package awsquicksight


// The conditional formatting that determines the shape of the filled map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filledMapShapeConditionalFormattingProperty := &FilledMapShapeConditionalFormattingProperty{
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	Format: &ShapeConditionalFormatProperty{
//   		BackgroundColor: &ConditionalFormattingColorProperty{
//   			Gradient: &ConditionalFormattingGradientColorProperty{
//   				Color: &GradientColorProperty{
//   					Stops: []interface{}{
//   						&GradientStopProperty{
//   							GradientOffset: jsii.Number(123),
//
//   							// the properties below are optional
//   							Color: jsii.String("color"),
//   							DataValue: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Expression: jsii.String("expression"),
//   			},
//   			Solid: &ConditionalFormattingSolidColorProperty{
//   				Expression: jsii.String("expression"),
//
//   				// the properties below are optional
//   				Color: jsii.String("color"),
//   			},
//   		},
//   	},
//   }
//
type CfnAnalysis_FilledMapShapeConditionalFormattingProperty struct {
	// The field ID of the filled map shape.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The conditional formatting that determines the background color of a filled map's shape.
	Format interface{} `field:"optional" json:"format" yaml:"format"`
}


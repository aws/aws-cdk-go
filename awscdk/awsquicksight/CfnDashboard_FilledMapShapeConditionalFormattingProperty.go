package awsquicksight


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
type CfnDashboard_FilledMapShapeConditionalFormattingProperty struct {
	// `CfnDashboard.FilledMapShapeConditionalFormattingProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// `CfnDashboard.FilledMapShapeConditionalFormattingProperty.Format`.
	Format interface{} `field:"optional" json:"format" yaml:"format"`
}


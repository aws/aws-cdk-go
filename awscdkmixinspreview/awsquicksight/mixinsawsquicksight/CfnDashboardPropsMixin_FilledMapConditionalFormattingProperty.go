package mixinsawsquicksight


// The conditional formatting of a `FilledMapVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filledMapConditionalFormattingProperty := &FilledMapConditionalFormattingProperty{
//   	ConditionalFormattingOptions: []interface{}{
//   		&FilledMapConditionalFormattingOptionProperty{
//   			Shape: &FilledMapShapeConditionalFormattingProperty{
//   				FieldId: jsii.String("fieldId"),
//   				Format: &ShapeConditionalFormatProperty{
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
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapconditionalformatting.html
//
type CfnDashboardPropsMixin_FilledMapConditionalFormattingProperty struct {
	// Conditional formatting options of a `FilledMapVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapconditionalformatting.html#cfn-quicksight-dashboard-filledmapconditionalformatting-conditionalformattingoptions
	//
	ConditionalFormattingOptions interface{} `field:"optional" json:"conditionalFormattingOptions" yaml:"conditionalFormattingOptions"`
}


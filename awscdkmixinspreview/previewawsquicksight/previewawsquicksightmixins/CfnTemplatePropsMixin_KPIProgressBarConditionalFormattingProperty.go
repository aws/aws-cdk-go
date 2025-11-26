package previewawsquicksightmixins


// The conditional formatting for the progress bar of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kPIProgressBarConditionalFormattingProperty := &KPIProgressBarConditionalFormattingProperty{
//   	ForegroundColor: &ConditionalFormattingColorProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpiprogressbarconditionalformatting.html
//
type CfnTemplatePropsMixin_KPIProgressBarConditionalFormattingProperty struct {
	// The conditional formatting of the progress bar's foreground color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpiprogressbarconditionalformatting.html#cfn-quicksight-template-kpiprogressbarconditionalformatting-foregroundcolor
	//
	ForegroundColor interface{} `field:"optional" json:"foregroundColor" yaml:"foregroundColor"`
}


package awsquicksight


// The conditional formatting for the progress bar of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kPIProgressBarConditionalFormattingProperty := &KPIProgressBarConditionalFormattingProperty{
//   	ForegroundColor: &ConditionalFormattingColorProperty{
//   		Gradient: &ConditionalFormattingGradientColorProperty{
//   			Color: &GradientColorProperty{
//   				Stops: []interface{}{
//   					&GradientStopProperty{
//   						GradientOffset: jsii.Number(123),
//
//   						// the properties below are optional
//   						Color: jsii.String("color"),
//   						DataValue: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Expression: jsii.String("expression"),
//   		},
//   		Solid: &ConditionalFormattingSolidColorProperty{
//   			Expression: jsii.String("expression"),
//
//   			// the properties below are optional
//   			Color: jsii.String("color"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiprogressbarconditionalformatting.html
//
type CfnDashboard_KPIProgressBarConditionalFormattingProperty struct {
	// The conditional formatting of the progress bar's foreground color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-kpiprogressbarconditionalformatting.html#cfn-quicksight-dashboard-kpiprogressbarconditionalformatting-foregroundcolor
	//
	ForegroundColor interface{} `field:"optional" json:"foregroundColor" yaml:"foregroundColor"`
}


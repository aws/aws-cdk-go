package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingGradientColorProperty := &ConditionalFormattingGradientColorProperty{
//   	Color: &GradientColorProperty{
//   		Stops: []interface{}{
//   			&GradientStopProperty{
//   				GradientOffset: jsii.Number(123),
//
//   				// the properties below are optional
//   				Color: jsii.String("color"),
//   				DataValue: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Expression: jsii.String("expression"),
//   }
//
type CfnAnalysis_ConditionalFormattingGradientColorProperty struct {
	// `CfnAnalysis.ConditionalFormattingGradientColorProperty.Color`.
	Color interface{} `field:"required" json:"color" yaml:"color"`
	// `CfnAnalysis.ConditionalFormattingGradientColorProperty.Expression`.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}


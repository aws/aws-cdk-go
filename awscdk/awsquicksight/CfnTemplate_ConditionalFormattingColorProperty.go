package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingColorProperty := &ConditionalFormattingColorProperty{
//   	Gradient: &ConditionalFormattingGradientColorProperty{
//   		Color: &GradientColorProperty{
//   			Stops: []interface{}{
//   				&GradientStopProperty{
//   					GradientOffset: jsii.Number(123),
//
//   					// the properties below are optional
//   					Color: jsii.String("color"),
//   					DataValue: jsii.Number(123),
//   				},
//   			},
//   		},
//   		Expression: jsii.String("expression"),
//   	},
//   	Solid: &ConditionalFormattingSolidColorProperty{
//   		Expression: jsii.String("expression"),
//
//   		// the properties below are optional
//   		Color: jsii.String("color"),
//   	},
//   }
//
type CfnTemplate_ConditionalFormattingColorProperty struct {
	// `CfnTemplate.ConditionalFormattingColorProperty.Gradient`.
	Gradient interface{} `field:"optional" json:"gradient" yaml:"gradient"`
	// `CfnTemplate.ConditionalFormattingColorProperty.Solid`.
	Solid interface{} `field:"optional" json:"solid" yaml:"solid"`
}


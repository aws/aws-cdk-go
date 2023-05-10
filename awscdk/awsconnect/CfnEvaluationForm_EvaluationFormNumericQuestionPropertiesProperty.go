package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormNumericQuestionPropertiesProperty := &EvaluationFormNumericQuestionPropertiesProperty{
//   	MaxValue: jsii.Number(123),
//   	MinValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   		PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   			Label: jsii.String("label"),
//   		},
//   	},
//   	Options: []interface{}{
//   		&EvaluationFormNumericQuestionOptionProperty{
//   			MaxValue: jsii.Number(123),
//   			MinValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			AutomaticFail: jsii.Boolean(false),
//   			Score: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnEvaluationForm_EvaluationFormNumericQuestionPropertiesProperty struct {
	// `CfnEvaluationForm.EvaluationFormNumericQuestionPropertiesProperty.MaxValue`.
	MaxValue *float64 `field:"required" json:"maxValue" yaml:"maxValue"`
	// `CfnEvaluationForm.EvaluationFormNumericQuestionPropertiesProperty.MinValue`.
	MinValue *float64 `field:"required" json:"minValue" yaml:"minValue"`
	// `CfnEvaluationForm.EvaluationFormNumericQuestionPropertiesProperty.Automation`.
	Automation interface{} `field:"optional" json:"automation" yaml:"automation"`
	// `CfnEvaluationForm.EvaluationFormNumericQuestionPropertiesProperty.Options`.
	Options interface{} `field:"optional" json:"options" yaml:"options"`
}


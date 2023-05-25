package awsconnect


// Information about the automation configuration in numeric questions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormNumericQuestionAutomationProperty := &EvaluationFormNumericQuestionAutomationProperty{
//   	PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   		Label: jsii.String("label"),
//   	},
//   }
//
type CfnEvaluationForm_EvaluationFormNumericQuestionAutomationProperty struct {
	// The property value of the automation.
	PropertyValue interface{} `field:"required" json:"propertyValue" yaml:"propertyValue"`
}


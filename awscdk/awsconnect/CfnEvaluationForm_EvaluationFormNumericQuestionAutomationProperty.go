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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformnumericquestionautomation.html
//
type CfnEvaluationForm_EvaluationFormNumericQuestionAutomationProperty struct {
	// The property value of the automation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformnumericquestionautomation.html#cfn-connect-evaluationform-evaluationformnumericquestionautomation-propertyvalue
	//
	PropertyValue interface{} `field:"optional" json:"propertyValue" yaml:"propertyValue"`
}


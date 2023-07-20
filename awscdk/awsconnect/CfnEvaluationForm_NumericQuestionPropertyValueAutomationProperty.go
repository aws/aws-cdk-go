package awsconnect


// Information about the property value used in automation of a numeric questions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericQuestionPropertyValueAutomationProperty := &NumericQuestionPropertyValueAutomationProperty{
//   	Label: jsii.String("label"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-numericquestionpropertyvalueautomation.html
//
type CfnEvaluationForm_NumericQuestionPropertyValueAutomationProperty struct {
	// The property label of the automation.
	//
	// *Allowed values* : `OVERALL_CUSTOMER_SENTIMENT_SCORE` , `OVERALL_AGENT_SENTIMENT_SCORE` | `NON_TALK_TIME` | `NON_TALK_TIME_PERCENTAGE` | `NUMBER_OF_INTERRUPTIONS` | `CONTACT_DURATION` | `AGENT_INTERACTION_DURATION` | `CUSTOMER_HOLD_TIME`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-numericquestionpropertyvalueautomation.html#cfn-connect-evaluationform-numericquestionpropertyvalueautomation-label
	//
	Label *string `field:"required" json:"label" yaml:"label"`
}


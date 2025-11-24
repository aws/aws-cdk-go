package mixinsawsconnect


// Information about the property value used in automation of a numeric questions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   numericQuestionPropertyValueAutomationProperty := &NumericQuestionPropertyValueAutomationProperty{
//   	Label: jsii.String("label"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-numericquestionpropertyvalueautomation.html
//
type CfnEvaluationFormPropsMixin_NumericQuestionPropertyValueAutomationProperty struct {
	// The property label of the automation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-numericquestionpropertyvalueautomation.html#cfn-connect-evaluationform-numericquestionpropertyvalueautomation-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
}


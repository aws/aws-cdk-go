package awsconnect


// An option for a multi-select question in an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormMultiSelectQuestionOptionProperty := &EvaluationFormMultiSelectQuestionOptionProperty{
//   	RefId: jsii.String("refId"),
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html
//
type CfnEvaluationForm_EvaluationFormMultiSelectQuestionOptionProperty struct {
	// Reference identifier for this option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-refid
	//
	RefId *string `field:"required" json:"refId" yaml:"refId"`
	// Display text for this option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformmultiselectquestionoption.html#cfn-connect-evaluationform-evaluationformmultiselectquestionoption-text
	//
	Text *string `field:"required" json:"text" yaml:"text"`
}


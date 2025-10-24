package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormQuestionAutomationAnswerSourceProperty := &EvaluationFormQuestionAutomationAnswerSourceProperty{
//   	SourceType: jsii.String("sourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestionautomationanswersource.html
//
type CfnEvaluationForm_EvaluationFormQuestionAutomationAnswerSourceProperty struct {
	// The type of the answer sourcr.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformquestionautomationanswersource.html#cfn-connect-evaluationform-evaluationformquestionautomationanswersource-sourcetype
	//
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
}


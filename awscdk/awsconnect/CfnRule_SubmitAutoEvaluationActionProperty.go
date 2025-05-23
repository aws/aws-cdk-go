package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   submitAutoEvaluationActionProperty := &SubmitAutoEvaluationActionProperty{
//   	EvaluationFormArn: jsii.String("evaluationFormArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-submitautoevaluationaction.html
//
type CfnRule_SubmitAutoEvaluationActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-submitautoevaluationaction.html#cfn-connect-rule-submitautoevaluationaction-evaluationformarn
	//
	EvaluationFormArn *string `field:"required" json:"evaluationFormArn" yaml:"evaluationFormArn"`
}


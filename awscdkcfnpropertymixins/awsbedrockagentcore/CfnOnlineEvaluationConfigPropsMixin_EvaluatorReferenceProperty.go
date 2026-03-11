package awsbedrockagentcore


// The reference to an evaluator used in online evaluation configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   evaluatorReferenceProperty := &EvaluatorReferenceProperty{
//   	EvaluatorId: jsii.String("evaluatorId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-evaluatorreference.html
//
type CfnOnlineEvaluationConfigPropsMixin_EvaluatorReferenceProperty struct {
	// The unique identifier of the evaluator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-evaluatorreference.html#cfn-bedrockagentcore-onlineevaluationconfig-evaluatorreference-evaluatorid
	//
	EvaluatorId *string `field:"optional" json:"evaluatorId" yaml:"evaluatorId"`
}


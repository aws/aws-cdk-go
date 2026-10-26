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
	// The ID of the evaluator: a built-in evaluator ID (Builtin.*), a managed third-party evaluator ID (ThirdParty.<provider>.<name>), or a custom evaluator ID. For a custom evaluator in the same template, use Fn::GetAtt with EvaluatorId. Ref returns the evaluator ARN, which this property does not accept.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-evaluatorreference.html#cfn-bedrockagentcore-onlineevaluationconfig-evaluatorreference-evaluatorid
	//
	EvaluatorId *string `field:"optional" json:"evaluatorId" yaml:"evaluatorId"`
}


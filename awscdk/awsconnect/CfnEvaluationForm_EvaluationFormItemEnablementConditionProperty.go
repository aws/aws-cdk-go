package awsconnect


// A condition for item enablement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormItemEnablementConditionProperty := &EvaluationFormItemEnablementConditionProperty{
//   	Operands: []interface{}{
//   		&EvaluationFormItemEnablementConditionOperandProperty{
//   			Expression: &EvaluationFormItemEnablementExpressionProperty{
//   				Comparator: jsii.String("comparator"),
//   				Source: &EvaluationFormItemEnablementSourceProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					RefId: jsii.String("refId"),
//   				},
//   				Values: []interface{}{
//   					&EvaluationFormItemEnablementSourceValueProperty{
//   						RefId: jsii.String("refId"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Operator: jsii.String("operator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementcondition.html
//
type CfnEvaluationForm_EvaluationFormItemEnablementConditionProperty struct {
	// Operands of the enablement condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementcondition.html#cfn-connect-evaluationform-evaluationformitemenablementcondition-operands
	//
	Operands interface{} `field:"required" json:"operands" yaml:"operands"`
	// The operator to be used to be applied to operands if more than one provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementcondition.html#cfn-connect-evaluationform-evaluationformitemenablementcondition-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}


package awsconnect


// An operand of the enablement condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormItemEnablementConditionOperandProperty := &EvaluationFormItemEnablementConditionOperandProperty{
//   	Expression: &EvaluationFormItemEnablementExpressionProperty{
//   		Comparator: jsii.String("comparator"),
//   		Source: &EvaluationFormItemEnablementSourceProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			RefId: jsii.String("refId"),
//   		},
//   		Values: []interface{}{
//   			&EvaluationFormItemEnablementSourceValueProperty{
//   				RefId: jsii.String("refId"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementconditionoperand.html
//
type CfnEvaluationForm_EvaluationFormItemEnablementConditionOperandProperty struct {
	// An expression of the enablement condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementconditionoperand.html#cfn-connect-evaluationform-evaluationformitemenablementconditionoperand-expression
	//
	Expression interface{} `field:"optional" json:"expression" yaml:"expression"`
}


package awsconnect


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
	// The list of operands that compose the condition.
	//
	// Each operand represents a specific criteria to be evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementcondition.html#cfn-connect-evaluationform-evaluationformitemenablementcondition-operands
	//
	Operands interface{} `field:"required" json:"operands" yaml:"operands"`
	// The logical operator used to combine multiple operands, determining how the condition is evaluated as a whole.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementcondition.html#cfn-connect-evaluationform-evaluationformitemenablementcondition-operator
	//
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}


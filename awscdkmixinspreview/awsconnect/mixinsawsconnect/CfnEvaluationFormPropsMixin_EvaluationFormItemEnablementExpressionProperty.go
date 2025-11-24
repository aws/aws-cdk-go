package mixinsawsconnect


// An expression that defines a basic building block of conditional enablement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormItemEnablementExpressionProperty := &EvaluationFormItemEnablementExpressionProperty{
//   	Comparator: jsii.String("comparator"),
//   	Source: &EvaluationFormItemEnablementSourceProperty{
//   		RefId: jsii.String("refId"),
//   		Type: jsii.String("type"),
//   	},
//   	Values: []interface{}{
//   		&EvaluationFormItemEnablementSourceValueProperty{
//   			RefId: jsii.String("refId"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementexpression.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormItemEnablementExpressionProperty struct {
	// A comparator to be used against list of values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementexpression.html#cfn-connect-evaluationform-evaluationformitemenablementexpression-comparator
	//
	Comparator *string `field:"optional" json:"comparator" yaml:"comparator"`
	// A source item of enablement expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementexpression.html#cfn-connect-evaluationform-evaluationformitemenablementexpression-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// A list of values from source item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementexpression.html#cfn-connect-evaluationform-evaluationformitemenablementexpression-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}


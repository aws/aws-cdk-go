package awsconnect


// An item enablement configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormItemEnablementConfigurationProperty := &EvaluationFormItemEnablementConfigurationProperty{
//   	Action: jsii.String("action"),
//   	Condition: &EvaluationFormItemEnablementConditionProperty{
//   		Operands: []interface{}{
//   			&EvaluationFormItemEnablementConditionOperandProperty{
//   				Expression: &EvaluationFormItemEnablementExpressionProperty{
//   					Comparator: jsii.String("comparator"),
//   					Source: &EvaluationFormItemEnablementSourceProperty{
//   						Type: jsii.String("type"),
//
//   						// the properties below are optional
//   						RefId: jsii.String("refId"),
//   					},
//   					Values: []interface{}{
//   						&EvaluationFormItemEnablementSourceValueProperty{
//   							RefId: jsii.String("refId"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Operator: jsii.String("operator"),
//   	},
//
//   	// the properties below are optional
//   	DefaultAction: jsii.String("defaultAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementconfiguration.html
//
type CfnEvaluationForm_EvaluationFormItemEnablementConfigurationProperty struct {
	// An enablement action that if condition is satisfied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementconfiguration.html#cfn-connect-evaluationform-evaluationformitemenablementconfiguration-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// A condition for item enablement configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementconfiguration.html#cfn-connect-evaluationform-evaluationformitemenablementconfiguration-condition
	//
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
	// An enablement action that if condition is not satisfied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformitemenablementconfiguration.html#cfn-connect-evaluationform-evaluationformitemenablementconfiguration-defaultaction
	//
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
}


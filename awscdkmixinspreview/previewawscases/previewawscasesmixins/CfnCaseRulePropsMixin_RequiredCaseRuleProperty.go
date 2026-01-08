package previewawscasesmixins


// Required rule type, used to indicate whether a field is required.
//
// In the Amazon Connect admin website, case rules are known as *case field conditions* . For more information about case field conditions, see [Add case field conditions to a case template](https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var emptyValue interface{}
//
//   requiredCaseRuleProperty := &RequiredCaseRuleProperty{
//   	Conditions: []interface{}{
//   		&BooleanConditionProperty{
//   			EqualTo: &BooleanOperandsProperty{
//   				OperandOne: &OperandOneProperty{
//   					FieldId: jsii.String("fieldId"),
//   				},
//   				OperandTwo: &OperandTwoProperty{
//   					BooleanValue: jsii.Boolean(false),
//   					DoubleValue: jsii.Number(123),
//   					EmptyValue: emptyValue,
//   					StringValue: jsii.String("stringValue"),
//   				},
//   				Result: jsii.Boolean(false),
//   			},
//   			NotEqualTo: &BooleanOperandsProperty{
//   				OperandOne: &OperandOneProperty{
//   					FieldId: jsii.String("fieldId"),
//   				},
//   				OperandTwo: &OperandTwoProperty{
//   					BooleanValue: jsii.Boolean(false),
//   					DoubleValue: jsii.Number(123),
//   					EmptyValue: emptyValue,
//   					StringValue: jsii.String("stringValue"),
//   				},
//   				Result: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	DefaultValue: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-requiredcaserule.html
//
type CfnCaseRulePropsMixin_RequiredCaseRuleProperty struct {
	// List of conditions for the required rule;
	//
	// the first condition to evaluate to true dictates the value of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-requiredcaserule.html#cfn-cases-caserule-requiredcaserule-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The value of the rule (that is, whether the field is required) should none of the conditions evaluate to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-requiredcaserule.html#cfn-cases-caserule-requiredcaserule-defaultvalue
	//
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
}


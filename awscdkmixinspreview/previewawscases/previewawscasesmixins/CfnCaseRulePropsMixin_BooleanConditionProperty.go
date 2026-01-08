package previewawscasesmixins


// Boolean condition for a rule.
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
//   booleanConditionProperty := &BooleanConditionProperty{
//   	EqualTo: &BooleanOperandsProperty{
//   		OperandOne: &OperandOneProperty{
//   			FieldId: jsii.String("fieldId"),
//   		},
//   		OperandTwo: &OperandTwoProperty{
//   			BooleanValue: jsii.Boolean(false),
//   			DoubleValue: jsii.Number(123),
//   			EmptyValue: emptyValue,
//   			StringValue: jsii.String("stringValue"),
//   		},
//   		Result: jsii.Boolean(false),
//   	},
//   	NotEqualTo: &BooleanOperandsProperty{
//   		OperandOne: &OperandOneProperty{
//   			FieldId: jsii.String("fieldId"),
//   		},
//   		OperandTwo: &OperandTwoProperty{
//   			BooleanValue: jsii.Boolean(false),
//   			DoubleValue: jsii.Number(123),
//   			EmptyValue: emptyValue,
//   			StringValue: jsii.String("stringValue"),
//   		},
//   		Result: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-booleancondition.html
//
type CfnCaseRulePropsMixin_BooleanConditionProperty struct {
	// Tests that operandOne is equal to operandTwo.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-booleancondition.html#cfn-cases-caserule-booleancondition-equalto
	//
	EqualTo interface{} `field:"optional" json:"equalTo" yaml:"equalTo"`
	// Tests that operandOne is not equal to operandTwo.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-booleancondition.html#cfn-cases-caserule-booleancondition-notequalto
	//
	NotEqualTo interface{} `field:"optional" json:"notEqualTo" yaml:"notEqualTo"`
}


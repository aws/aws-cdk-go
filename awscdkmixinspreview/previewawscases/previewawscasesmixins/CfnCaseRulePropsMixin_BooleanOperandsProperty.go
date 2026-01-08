package previewawscasesmixins


// Boolean operands for a condition.
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
//   booleanOperandsProperty := &BooleanOperandsProperty{
//   	OperandOne: &OperandOneProperty{
//   		FieldId: jsii.String("fieldId"),
//   	},
//   	OperandTwo: &OperandTwoProperty{
//   		BooleanValue: jsii.Boolean(false),
//   		DoubleValue: jsii.Number(123),
//   		EmptyValue: emptyValue,
//   		StringValue: jsii.String("stringValue"),
//   	},
//   	Result: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-booleanoperands.html
//
type CfnCaseRulePropsMixin_BooleanOperandsProperty struct {
	// Represents the left hand operand in the condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-booleanoperands.html#cfn-cases-caserule-booleanoperands-operandone
	//
	OperandOne interface{} `field:"optional" json:"operandOne" yaml:"operandOne"`
	// Represents the right hand operand in the condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-booleanoperands.html#cfn-cases-caserule-booleanoperands-operandtwo
	//
	OperandTwo interface{} `field:"optional" json:"operandTwo" yaml:"operandTwo"`
	// The value of the outer rule if the condition evaluates to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-booleanoperands.html#cfn-cases-caserule-booleanoperands-result
	//
	Result interface{} `field:"optional" json:"result" yaml:"result"`
}


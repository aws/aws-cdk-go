package previewawscasesmixins


// Represents the right hand operand in the condition.
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
//   operandTwoProperty := &OperandTwoProperty{
//   	BooleanValue: jsii.Boolean(false),
//   	DoubleValue: jsii.Number(123),
//   	EmptyValue: emptyValue,
//   	StringValue: jsii.String("stringValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-operandtwo.html
//
type CfnCaseRulePropsMixin_OperandTwoProperty struct {
	// Boolean value type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-operandtwo.html#cfn-cases-caserule-operandtwo-booleanvalue
	//
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// Double value type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-operandtwo.html#cfn-cases-caserule-operandtwo-doublevalue
	//
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// Represents an empty operand value.
	//
	// In the Amazon Connect admin website, case rules are known as *case field conditions* . For more information about case field conditions, see [Add case field conditions to a case template](https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-operandtwo.html#cfn-cases-caserule-operandtwo-emptyvalue
	//
	EmptyValue interface{} `field:"optional" json:"emptyValue" yaml:"emptyValue"`
	// String value type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-operandtwo.html#cfn-cases-caserule-operandtwo-stringvalue
	//
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}


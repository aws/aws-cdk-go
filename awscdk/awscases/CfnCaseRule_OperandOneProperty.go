package awscases


// Represents the left hand operand in the condition.
//
// In the Amazon Connect admin website, case rules are known as *case field conditions* . For more information about case field conditions, see [Add case field conditions to a case template](https://docs.aws.amazon.com/connect/latest/adminguide/case-field-conditions.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   operandOneProperty := &OperandOneProperty{
//   	FieldId: jsii.String("fieldId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-operandone.html
//
type CfnCaseRule_OperandOneProperty struct {
	// The field ID that this operand should take the value of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-operandone.html#cfn-cases-caserule-operandone-fieldid
	//
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
}


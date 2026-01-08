package previewawscasesmixins


// Represents what rule type should take place, under what conditions.
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
//   caseRuleDetailsProperty := &CaseRuleDetailsProperty{
//   	Hidden: &HiddenCaseRuleProperty{
//   		Conditions: []interface{}{
//   			&BooleanConditionProperty{
//   				EqualTo: &BooleanOperandsProperty{
//   					OperandOne: &OperandOneProperty{
//   						FieldId: jsii.String("fieldId"),
//   					},
//   					OperandTwo: &OperandTwoProperty{
//   						BooleanValue: jsii.Boolean(false),
//   						DoubleValue: jsii.Number(123),
//   						EmptyValue: emptyValue,
//   						StringValue: jsii.String("stringValue"),
//   					},
//   					Result: jsii.Boolean(false),
//   				},
//   				NotEqualTo: &BooleanOperandsProperty{
//   					OperandOne: &OperandOneProperty{
//   						FieldId: jsii.String("fieldId"),
//   					},
//   					OperandTwo: &OperandTwoProperty{
//   						BooleanValue: jsii.Boolean(false),
//   						DoubleValue: jsii.Number(123),
//   						EmptyValue: emptyValue,
//   						StringValue: jsii.String("stringValue"),
//   					},
//   					Result: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		DefaultValue: jsii.Boolean(false),
//   	},
//   	Required: &RequiredCaseRuleProperty{
//   		Conditions: []interface{}{
//   			&BooleanConditionProperty{
//   				EqualTo: &BooleanOperandsProperty{
//   					OperandOne: &OperandOneProperty{
//   						FieldId: jsii.String("fieldId"),
//   					},
//   					OperandTwo: &OperandTwoProperty{
//   						BooleanValue: jsii.Boolean(false),
//   						DoubleValue: jsii.Number(123),
//   						EmptyValue: emptyValue,
//   						StringValue: jsii.String("stringValue"),
//   					},
//   					Result: jsii.Boolean(false),
//   				},
//   				NotEqualTo: &BooleanOperandsProperty{
//   					OperandOne: &OperandOneProperty{
//   						FieldId: jsii.String("fieldId"),
//   					},
//   					OperandTwo: &OperandTwoProperty{
//   						BooleanValue: jsii.Boolean(false),
//   						DoubleValue: jsii.Number(123),
//   						EmptyValue: emptyValue,
//   						StringValue: jsii.String("stringValue"),
//   					},
//   					Result: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		DefaultValue: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-caseruledetails.html
//
type CfnCaseRulePropsMixin_CaseRuleDetailsProperty struct {
	// Whether a field is visible, based on values in other fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-caseruledetails.html#cfn-cases-caserule-caseruledetails-hidden
	//
	Hidden interface{} `field:"optional" json:"hidden" yaml:"hidden"`
	// Required rule type, used to indicate whether a field is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-caseruledetails.html#cfn-cases-caserule-caseruledetails-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}


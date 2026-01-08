package previewawscasesmixins


// A rule that controls field visibility based on conditions.
//
// Fields can be shown or hidden dynamically based on values in other fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var emptyValue interface{}
//
//   hiddenCaseRuleProperty := &HiddenCaseRuleProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-hiddencaserule.html
//
type CfnCaseRulePropsMixin_HiddenCaseRuleProperty struct {
	// A list of conditions that determine field visibility.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-hiddencaserule.html#cfn-cases-caserule-hiddencaserule-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Whether the field is hidden when no conditions match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-caserule-hiddencaserule.html#cfn-cases-caserule-hiddencaserule-defaultvalue
	//
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
}


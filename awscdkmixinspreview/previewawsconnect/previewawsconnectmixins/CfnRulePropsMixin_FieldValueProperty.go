package previewawsconnectmixins


// Object for case field values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var emptyValue interface{}
//
//   fieldValueProperty := &FieldValueProperty{
//   	BooleanValue: jsii.Boolean(false),
//   	DoubleValue: jsii.Number(123),
//   	EmptyValue: emptyValue,
//   	StringValue: jsii.String("stringValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-fieldvalue.html
//
type CfnRulePropsMixin_FieldValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-fieldvalue.html#cfn-connect-rule-fieldvalue-booleanvalue
	//
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-fieldvalue.html#cfn-connect-rule-fieldvalue-doublevalue
	//
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-fieldvalue.html#cfn-connect-rule-fieldvalue-emptyvalue
	//
	EmptyValue interface{} `field:"optional" json:"emptyValue" yaml:"emptyValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-fieldvalue.html#cfn-connect-rule-fieldvalue-stringvalue
	//
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}


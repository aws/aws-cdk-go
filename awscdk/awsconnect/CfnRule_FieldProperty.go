package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var emptyValue interface{}
//
//   fieldProperty := &FieldProperty{
//   	Id: jsii.String("id"),
//   	Value: &FieldValueProperty{
//   		BooleanValue: jsii.Boolean(false),
//   		DoubleValue: jsii.Number(123),
//   		EmptyValue: emptyValue,
//   		StringValue: jsii.String("stringValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-field.html
//
type CfnRule_FieldProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-field.html#cfn-connect-rule-field-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Object for case field values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-field.html#cfn-connect-rule-field-value
	//
	Value interface{} `field:"required" json:"value" yaml:"value"`
}


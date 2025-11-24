package mixinsawsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var emptyValue interface{}
//
//   updateCaseActionProperty := &UpdateCaseActionProperty{
//   	Fields: []interface{}{
//   		&FieldProperty{
//   			Id: jsii.String("id"),
//   			Value: &FieldValueProperty{
//   				BooleanValue: jsii.Boolean(false),
//   				DoubleValue: jsii.Number(123),
//   				EmptyValue: emptyValue,
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-updatecaseaction.html
//
type CfnRulePropsMixin_UpdateCaseActionProperty struct {
	// An array of case fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-updatecaseaction.html#cfn-connect-rule-updatecaseaction-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
}


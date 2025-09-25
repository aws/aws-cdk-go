package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var emptyValue interface{}
//
//   createCaseActionProperty := &CreateCaseActionProperty{
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
//   	TemplateId: jsii.String("templateId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-createcaseaction.html
//
type CfnRule_CreateCaseActionProperty struct {
	// An array of case fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-createcaseaction.html#cfn-connect-rule-createcaseaction-fields
	//
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-rule-createcaseaction.html#cfn-connect-rule-createcaseaction-templateid
	//
	TemplateId *string `field:"required" json:"templateId" yaml:"templateId"`
}


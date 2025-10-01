package awskendra


// The condition used for the target document attribute or metadata field when ingesting documents into Amazon Kendra.
//
// You use this with [DocumentAttributeTarget to apply the condition](https://docs.aws.amazon.com/kendra/latest/dg/API_DocumentAttributeTarget.html) .
//
// For example, you can create the 'Department' target field and have it prefill department names associated with the documents based on information in the 'Source_URI' field. Set the condition that if the 'Source_URI' field contains 'financial' in its URI value, then prefill the target field 'Department' with the target value 'Finance' for the document.
//
// Amazon Kendra cannot create a target field if it has not already been created as an index field. After you create your index field, you can create a document metadata field using `DocumentAttributeTarget` . Amazon Kendra then will map your newly created metadata field to your index field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentAttributeConditionProperty := &DocumentAttributeConditionProperty{
//   	ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   	Operator: jsii.String("operator"),
//
//   	// the properties below are optional
//   	ConditionOnValue: &DocumentAttributeValueProperty{
//   		DateValue: jsii.String("dateValue"),
//   		LongValue: jsii.Number(123),
//   		StringListValue: []*string{
//   			jsii.String("stringListValue"),
//   		},
//   		StringValue: jsii.String("stringValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-documentattributecondition.html
//
type CfnDataSource_DocumentAttributeConditionProperty struct {
	// The identifier of the document attribute used for the condition.
	//
	// For example, 'Source_URI' could be an identifier for the attribute or metadata field that contains source URIs associated with the documents.
	//
	// Amazon Kendra currently does not support `_document_body` as an attribute key used for the condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-documentattributecondition.html#cfn-kendra-datasource-documentattributecondition-conditiondocumentattributekey
	//
	ConditionDocumentAttributeKey *string `field:"required" json:"conditionDocumentAttributeKey" yaml:"conditionDocumentAttributeKey"`
	// The condition operator.
	//
	// For example, you can use 'Contains' to partially match a string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-documentattributecondition.html#cfn-kendra-datasource-documentattributecondition-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The value used by the operator.
	//
	// For example, you can specify the value 'financial' for strings in the 'Source_URI' field that partially match or contain this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-documentattributecondition.html#cfn-kendra-datasource-documentattributecondition-conditiononvalue
	//
	ConditionOnValue interface{} `field:"optional" json:"conditionOnValue" yaml:"conditionOnValue"`
}


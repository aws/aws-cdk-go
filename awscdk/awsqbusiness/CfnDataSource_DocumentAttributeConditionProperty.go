package awsqbusiness


// The condition used for the target document attribute or metadata field when ingesting documents into Amazon Q Business.
//
// You use this with [`DocumentAttributeTarget`](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_DocumentAttributeTarget.html) to apply the condition.
//
// For example, you can create the 'Department' target field and have it prefill department names associated with the documents based on information in the 'Source_URI' field. Set the condition that if the 'Source_URI' field contains 'financial' in its URI value, then prefill the target field 'Department' with the target value 'Finance' for the document.
//
// Amazon Q Business can't create a target field if it has not already been created as an index field. After you create your index field, you can create a document metadata field using `DocumentAttributeTarget` . Amazon Q Business then will map your newly created metadata field to your index field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentAttributeConditionProperty := &DocumentAttributeConditionProperty{
//   	Key: jsii.String("key"),
//   	Operator: jsii.String("operator"),
//
//   	// the properties below are optional
//   	Value: &DocumentAttributeValueProperty{
//   		DateValue: jsii.String("dateValue"),
//   		LongValue: jsii.Number(123),
//   		StringListValue: []*string{
//   			jsii.String("stringListValue"),
//   		},
//   		StringValue: jsii.String("stringValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributecondition.html
//
type CfnDataSource_DocumentAttributeConditionProperty struct {
	// The identifier of the document attribute used for the condition.
	//
	// For example, 'Source_URI' could be an identifier for the attribute or metadata field that contains source URIs associated with the documents.
	//
	// Amazon Q Business currently doesn't support `_document_body` as an attribute key used for the condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributecondition.html#cfn-qbusiness-datasource-documentattributecondition-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The identifier of the document attribute used for the condition.
	//
	// For example, 'Source_URI' could be an identifier for the attribute or metadata field that contains source URIs associated with the documents.
	//
	// Amazon Q Business currently does not support `_document_body` as an attribute key used for the condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributecondition.html#cfn-qbusiness-datasource-documentattributecondition-operator
	//
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The value of a document attribute.
	//
	// You can only provide one value for a document attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributecondition.html#cfn-qbusiness-datasource-documentattributecondition-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}


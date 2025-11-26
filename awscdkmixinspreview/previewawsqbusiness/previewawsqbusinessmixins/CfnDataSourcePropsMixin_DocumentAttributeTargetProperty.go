package previewawsqbusinessmixins


// The target document attribute or metadata field you want to alter when ingesting documents into Amazon Q Business.
//
// For example, you can delete all customer identification numbers associated with the documents, stored in the document metadata field called 'Customer_ID' by setting the target key as 'Customer_ID' and the deletion flag to `TRUE` . This removes all customer ID values in the field 'Customer_ID'. This would scrub personally identifiable information from each document's metadata.
//
// Amazon Q Business can't create a target field if it has not already been created as an index field. After you create your index field, you can create a document metadata field using [`DocumentAttributeTarget`](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_DocumentAttributeTarget.html) . Amazon Q Business will then map your newly created document attribute to your index field.
//
// You can also use this with [`DocumentAttributeCondition`](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_DocumentAttributeCondition.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   documentAttributeTargetProperty := &DocumentAttributeTargetProperty{
//   	AttributeValueOperator: jsii.String("attributeValueOperator"),
//   	Key: jsii.String("key"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributetarget.html
//
type CfnDataSourcePropsMixin_DocumentAttributeTargetProperty struct {
	// `TRUE` to delete the existing target value for your specified target attribute key.
	//
	// You cannot create a target value and set this to `TRUE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributetarget.html#cfn-qbusiness-datasource-documentattributetarget-attributevalueoperator
	//
	AttributeValueOperator *string `field:"optional" json:"attributeValueOperator" yaml:"attributeValueOperator"`
	// The identifier of the target document attribute or metadata field.
	//
	// For example, 'Department' could be an identifier for the target attribute or metadata field that includes the department names associated with the documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributetarget.html#cfn-qbusiness-datasource-documentattributetarget-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of a document attribute.
	//
	// You can only provide one value for a document attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributetarget.html#cfn-qbusiness-datasource-documentattributetarget-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}


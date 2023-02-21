package awskendra


// The target document attribute or metadata field you want to alter when ingesting documents into Amazon Kendra.
//
// For example, you can delete customer identification numbers associated with the documents, stored in the document metadata field called 'Customer_ID'. You set the target key as 'Customer_ID' and the deletion flag to `TRUE` . This removes all customer ID values in the field 'Customer_ID'. This would scrub personally identifiable information from each document's metadata.
//
// Amazon Kendra cannot create a target field if it has not already been created as an index field. After you create your index field, you can create a document metadata field using `DocumentAttributeTarget` . Amazon Kendra then will map your newly created metadata field to your index field.
//
// You can also use this with [DocumentAttributeCondition](https://docs.aws.amazon.com/kendra/latest/dg/API_DocumentAttributeCondition.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentAttributeTargetProperty := &documentAttributeTargetProperty{
//   	targetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//
//   	// the properties below are optional
//   	targetDocumentAttributeValue: &documentAttributeValueProperty{
//   		dateValue: jsii.String("dateValue"),
//   		longValue: jsii.Number(123),
//   		stringListValue: []*string{
//   			jsii.String("stringListValue"),
//   		},
//   		stringValue: jsii.String("stringValue"),
//   	},
//   	targetDocumentAttributeValueDeletion: jsii.Boolean(false),
//   }
//
type CfnDataSource_DocumentAttributeTargetProperty struct {
	// The identifier of the target document attribute or metadata field.
	//
	// For example, 'Department' could be an identifier for the target attribute or metadata field that includes the department names associated with the documents.
	TargetDocumentAttributeKey *string `field:"required" json:"targetDocumentAttributeKey" yaml:"targetDocumentAttributeKey"`
	// The target value you want to create for the target attribute.
	//
	// For example, 'Finance' could be the target value for the target attribute key 'Department'.
	TargetDocumentAttributeValue interface{} `field:"optional" json:"targetDocumentAttributeValue" yaml:"targetDocumentAttributeValue"`
	// `TRUE` to delete the existing target value for your specified target attribute key.
	//
	// You cannot create a target value and set this to `TRUE` . To create a target value ( `TargetDocumentAttributeValue` ), set this to `FALSE` .
	TargetDocumentAttributeValueDeletion interface{} `field:"optional" json:"targetDocumentAttributeValueDeletion" yaml:"targetDocumentAttributeValueDeletion"`
}


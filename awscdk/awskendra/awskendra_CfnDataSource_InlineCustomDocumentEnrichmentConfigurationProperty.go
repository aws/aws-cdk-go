package awskendra


// Provides the configuration information for applying basic logic to alter document metadata and content when ingesting documents into Amazon Kendra.
//
// To apply advanced logic, to go beyond what you can do with basic logic, see [HookConfiguration](https://docs.aws.amazon.com/kendra/latest/dg/API_HookConfiguration.html) .
//
// For more information, see [Customizing document metadata during the ingestion process](https://docs.aws.amazon.com/kendra/latest/dg/custom-document-enrichment.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inlineCustomDocumentEnrichmentConfigurationProperty := &InlineCustomDocumentEnrichmentConfigurationProperty{
//   	Condition: &DocumentAttributeConditionProperty{
//   		ConditionDocumentAttributeKey: jsii.String("conditionDocumentAttributeKey"),
//   		Operator: jsii.String("operator"),
//
//   		// the properties below are optional
//   		ConditionOnValue: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	DocumentContentDeletion: jsii.Boolean(false),
//   	Target: &DocumentAttributeTargetProperty{
//   		TargetDocumentAttributeKey: jsii.String("targetDocumentAttributeKey"),
//
//   		// the properties below are optional
//   		TargetDocumentAttributeValue: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   		TargetDocumentAttributeValueDeletion: jsii.Boolean(false),
//   	},
//   }
//
type CfnDataSource_InlineCustomDocumentEnrichmentConfigurationProperty struct {
	// Configuration of the condition used for the target document attribute or metadata field when ingesting documents into Amazon Kendra.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// `TRUE` to delete content if the condition used for the target attribute is met.
	DocumentContentDeletion interface{} `field:"optional" json:"documentContentDeletion" yaml:"documentContentDeletion"`
	// Configuration of the target document attribute or metadata field when ingesting documents into Amazon Kendra.
	//
	// You can also include a value.
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}


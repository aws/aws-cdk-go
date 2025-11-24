package mixinsawsqbusiness


// Provides the configuration information for applying basic logic to alter document metadata and content when ingesting documents into Amazon Q Business.
//
// To apply advanced logic, to go beyond what you can do with basic logic, see [`HookConfiguration`](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_HookConfiguration.html) .
//
// For more information, see [Custom document enrichment](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/custom-document-enrichment.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inlineDocumentEnrichmentConfigurationProperty := &InlineDocumentEnrichmentConfigurationProperty{
//   	Condition: &DocumentAttributeConditionProperty{
//   		Key: jsii.String("key"),
//   		Operator: jsii.String("operator"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	DocumentContentOperator: jsii.String("documentContentOperator"),
//   	Target: &DocumentAttributeTargetProperty{
//   		AttributeValueOperator: jsii.String("attributeValueOperator"),
//   		Key: jsii.String("key"),
//   		Value: &DocumentAttributeValueProperty{
//   			DateValue: jsii.String("dateValue"),
//   			LongValue: jsii.Number(123),
//   			StringListValue: []*string{
//   				jsii.String("stringListValue"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html
//
type CfnDataSourcePropsMixin_InlineDocumentEnrichmentConfigurationProperty struct {
	// Configuration of the condition used for the target document attribute or metadata field when ingesting documents into Amazon Q Business .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html#cfn-qbusiness-datasource-inlinedocumentenrichmentconfiguration-condition
	//
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// `TRUE` to delete content if the condition used for the target attribute is met.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html#cfn-qbusiness-datasource-inlinedocumentenrichmentconfiguration-documentcontentoperator
	//
	DocumentContentOperator *string `field:"optional" json:"documentContentOperator" yaml:"documentContentOperator"`
	// Configuration of the target document attribute or metadata field when ingesting documents into Amazon Q Business .
	//
	// You can also include a value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html#cfn-qbusiness-datasource-inlinedocumentenrichmentconfiguration-target
	//
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}


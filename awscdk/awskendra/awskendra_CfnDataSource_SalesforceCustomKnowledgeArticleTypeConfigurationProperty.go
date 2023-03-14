package awskendra


// Provides the configuration information for indexing Salesforce custom articles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceCustomKnowledgeArticleTypeConfigurationProperty := &salesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   	documentDataFieldName: jsii.String("documentDataFieldName"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_SalesforceCustomKnowledgeArticleTypeConfigurationProperty struct {
	// The name of the field in the custom knowledge article that contains the document data to index.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the configuration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the field in the custom knowledge article that contains the document title.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Maps attributes or field names of the custom knowledge article to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Salesforce fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Salesforce data source field names must exist in your Salesforce custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}


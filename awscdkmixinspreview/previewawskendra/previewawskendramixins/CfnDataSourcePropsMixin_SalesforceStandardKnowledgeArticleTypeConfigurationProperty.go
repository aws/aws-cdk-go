package previewawskendramixins


// Provides the configuration information for standard Salesforce knowledge articles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   salesforceStandardKnowledgeArticleTypeConfigurationProperty := &SalesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   	DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   	DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	FieldMappings: []interface{}{
//   		&DataSourceToIndexFieldMappingProperty{
//   			DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			DateFieldFormat: jsii.String("dateFieldFormat"),
//   			IndexFieldName: jsii.String("indexFieldName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration.html
//
type CfnDataSourcePropsMixin_SalesforceStandardKnowledgeArticleTypeConfigurationProperty struct {
	// The name of the field that contains the document data to index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration.html#cfn-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration-documentdatafieldname
	//
	DocumentDataFieldName *string `field:"optional" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the field that contains the document title.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration.html#cfn-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration-documenttitlefieldname
	//
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Maps attributes or field names of the knowledge article to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Salesforce fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Salesforce data source field names must exist in your Salesforce custom metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration.html#cfn-kendra-datasource-salesforcestandardknowledgearticletypeconfiguration-fieldmappings
	//
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

